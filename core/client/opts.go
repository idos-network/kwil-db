package client

import (
	"github.com/kwilteam/kwil-db/core/crypto/auth"
	"github.com/kwilteam/kwil-db/core/log"
)

type Option func(*Client)

func WithLogger(logger log.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithSigner sets a signer to use when authoring transactions. The chain ID
// will be used in all transactions, which helps prevent replay attacks on
// different chains. On the initial connection, the remote node's chain ID is
// checked against ours to ensure were are on the right network. If the chain ID
// is empty, we will create and sign transactions for whatever network the
// remote node claims, which should only be done for testing or when in secure
// (TLS) communication with a trusted node.
func WithSigner(signer auth.Signer, chainID string) Option {
	return func(c *Client) {
		c.Signer = signer
		c.chainID = chainID
	}
}

func WithTLSCert(certFile string) Option {
	return func(c *Client) {
		c.tlsCertFile = certFile
	}
}

func WithTransportClient(tc TransportClient) Option {
	return func(c *Client) {
		c.transportClient = tc
	}
}

type callOptions struct {
	// forceAuthenticated is used to force the client to authenticate
	// if nil, the client will use the default value
	// if false, it will not authenticate
	// if true, it will authenticate
	forceAuthenticated *bool // is pointer necessary here?
}

type CallOpt func(*callOptions)

// Authenticated can be used to force the client to authenticate (or not)
// if true, the client will authenticate. if false, it will not authenticate
// if nil, the client will decide itself
func Authenticated(shouldSign bool) CallOpt {
	return func(o *callOptions) {
		copied := shouldSign
		o.forceAuthenticated = &copied
	}
}

type txOptions struct {
	nonce int64
}

type TxOpt func(*txOptions)

func WithNonce(nonce int64) TxOpt {
	return func(o *txOptions) {
		o.nonce = nonce
	}
}