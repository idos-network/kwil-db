package gatewayclient

import (
	"net/http"

	clientType "github.com/kwilteam/kwil-db/core/types/client"
)

// GatewayOptions are options that can be set for the gateway client
type GatewayOptions struct {
	clientType.Options

	// AuthSignFunc is a function that will be used to sign gateway authentication messages.
	AuthSignFunc GatewayAuthSignFunc

	AuthCookieHandler func(*http.Cookie) error
}

// DefaultOptions returns the default options for the gateway client.
func DefaultOptions() *GatewayOptions {
	return &GatewayOptions{
		Options: *clientType.DefaultOptions(),

		AuthSignFunc: defaultGatewayAuthSignFunc,
	}
}

// Apply applies the passed options to the receiver.
func (c *GatewayOptions) Apply(opt *GatewayOptions) {
	if opt == nil {
		return
	}

	c.Options.Apply(&opt.Options)

	if opt.AuthSignFunc != nil {
		c.AuthSignFunc = opt.AuthSignFunc
	}
	if opt.AuthCookieHandler != nil {
		c.AuthCookieHandler = opt.AuthCookieHandler
	}
}
