package events

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"kwil/x/common/config"
	cc "kwil/x/deposits/chainclient"
	ct "kwil/x/deposits/chainclient/types"
)

type EventFeed interface {
	Listen(context.Context) (<-chan int64, <-chan error, error)
	Contract() ct.Contract
}

type eventFeed struct {
	log       *zerolog.Logger
	conf      config.Config
	client    ct.Client
	reqConfs  uint16
	timeout   time.Duration
	lastBlock int64
	sc        ct.Contract
}

func New(conf config.Config, start int64) (EventFeed, error) {
	chnid := conf.String("chain")
	logger := log.With().Str("module", "deposits").Str("chainID", chnid).Logger()

	// build client
	cb := cc.Builder()

	client, err := cb.Chain(chnid).Endpoint(conf.String("provider-endpoint")).Build()
	if err != nil {
		return nil, err
	}

	// get contract
	sc, err := client.GetContract(conf.String("contract-address"))
	if err != nil {
		return nil, err
	}

	return &eventFeed{
		log:       &logger,
		conf:      conf,
		client:    client,
		reqConfs:  uint16(conf.Int("required-confirmations")),
		timeout:   time.Duration(conf.Int("block-timeout")) * time.Second,
		lastBlock: start,
		sc:        sc,
	}, nil
}

func (ef *eventFeed) Listen(
	ctx context.Context,
) (<-chan int64, <-chan error, error) {
	ef.log.Debug().Msg("starting event feed")

	return ef.listenForBlocks(ctx)
}

// ListenForBlocks listens for new block headers and sends them to the headers channel
// it should automatically reconnect if the connection is lost, and only send blocks that are finalized
func (ef *eventFeed) listenForBlocks(ctx context.Context) (<-chan int64, <-chan error, error) {
	ef.log.Debug().Msg("starting block listener")

	headers := make(chan int64, ef.conf.Int("block-buffer")) // this channel is for blocks before finalization
	errs := make(chan error)
	sub, err := ef.client.SubscribeBlocks(ctx, headers)
	if err != nil {
		return headers, errs, err
	}

	retChan := make(chan int64) // getting returned, this is for blocks after finalization

	q := NewQueue()

	go func() {
		for {
			select {
			case err := <-sub.Err():
				ef.log.Error().Err(err).Msg("error in block subscription, resubscribing")
				sub.Unsubscribe()
				sub, err = ef.resubscribe(ctx, headers)
				if err != nil {
					ef.log.Error().Err(err).Msg("error resubscribing to eth client, shutting down listener...")
					errs <- err
					return
				}
			case <-ctx.Done():
				ef.log.Debug().Msg("shutting down block listener")
				sub.Unsubscribe()
				return
			case <-time.After(ef.timeout):
				ef.log.Debug().Msg("block listener timed out, resubscribing")
				sub.Unsubscribe()
				sub, err = ef.resubscribe(ctx, headers)
				if err != nil {
					ef.log.Error().Err(err).Msg("error resubscribing to eth client, shutting down listener...")
					errs <- err
					return
				}
			case header := <-headers:
				ef.log.Debug().Int64("block", header).Msgf("received new unfinalized block")

				// now we need to ensure it is one greater than last
				if header == ef.lastBlock+1 { // expected is received
					ef.log.Debug().Int64("block", header).Msgf("received new finalized block")
					ef.lastBlock = header
					q.Append(header)
				} else if header > ef.lastBlock+1 { // received is greater than expected
					ef.log.Debug().Int64("received", header).Int64("expected", ef.lastBlock+1).Msg("received block is greater than expected, recovering skipped blocks")
					for i := ef.lastBlock + 1; i <= header; i++ {
						q.Append(i)
					}
				} else { // received is less than expected
					ef.log.Debug().Int64("received", header).Int64("expected", ef.lastBlock+1).Msg("received block is less than expected, ignoring")
					// do nothing
				}

				// now we need to check if we have any finalized blocks in the queue

				for {
					if q.Len() > ef.reqConfs {
						// we have enough blocks, send the oldest one
						retChan <- q.Pop()
					} else {
						// we don't have enough blocks, break out of the loop
						break
					}
				}
			}
		}
	}()

	return retChan, errs, nil
}

var ErrReconnect = errors.New("failed to reconnect")

// resubscribe will resubscribe to the block headers, and return the new subscription
func (ef *eventFeed) resubscribe(ctx context.Context, headers chan int64) (ct.BlockSubscription, error) {
	// I definitely need to refactor this
	ef.log.Debug().Msg("resubscribing to eth client")
	sub, err := ef.client.SubscribeBlocks(ctx, headers)
	if err != nil {
		ef.log.Warn().Err(err).Msg("error resubscribing to eth client, retrying in 1 second")
		time.Sleep(time.Second)
		sub, err = ef.client.SubscribeBlocks(ctx, headers)
		if err != nil {
			ef.log.Warn().Err(err).Msg("error resubscribing to eth client, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
			sub, err = ef.client.SubscribeBlocks(ctx, headers)
			if err != nil {
				ef.log.Warn().Err(err).Msg("error resubscribing to eth client, retrying in 10 seconds")
				time.Sleep(10 * time.Second)
				sub, err = ef.client.SubscribeBlocks(ctx, headers)
				if err != nil {
					ef.log.Warn().Err(err).Msg("error resubscribing to eth client, shutting down listener...")
					return nil, ErrReconnect
				}
			}
		}
	}

	return sub, nil
}

func (ef *eventFeed) Contract() ct.Contract {
	return ef.sc
}
