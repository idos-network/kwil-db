package node

import (
	"p2p/node/types"

	"github.com/libp2p/go-libp2p/core/host"
)

type options struct {
	port    uint64
	privKey []byte
	leader  bool
	pex     bool
	host    host.Host
	bs      types.BlockStore
	mp      types.MemPool
	ce      ConsensusEngine
}

func (o *options) set(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(*options)

func WithPort(port uint64) Option {
	return func(o *options) {
		o.port = port
	}
}
func WithPrivKey(privKey []byte) Option {
	return func(o *options) {
		o.privKey = privKey
	}
}
func WithLeader(leader bool) Option {
	return func(o *options) {
		o.leader = leader
	}
}
func WithPex(pex bool) Option {
	return func(o *options) {
		o.pex = pex
	}
}
func WithHost(host host.Host) Option {
	return func(o *options) {
		o.host = host
	}
}
func WithBlockStore(bs types.BlockStore) Option {
	return func(o *options) {
		o.bs = bs
	}
}
func WithMemPool(mp types.MemPool) Option {
	return func(o *options) {
		o.mp = mp
	}
}
func WithConsensusEngine(ce ConsensusEngine) Option {
	return func(o *options) {
		o.ce = ce
	}
}
