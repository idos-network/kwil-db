package dataset

import (
	"github.com/kwilteam/kwil-db/core/log"
)

// newTxOpts creates a new TxOpts with the default values
func newTxOpts() *TxOpts {
	return &TxOpts{
		Caller: &noCaller{},
	}
}

type TxOpts struct {
	// Caller is the user that is calling the transaction
	Caller User
}

type OpenOpt func(*Dataset)

func WithAvailableExtensions(exts map[string]Initializer) OpenOpt {
	return func(opts *Dataset) {
		opts.initializers = exts
	}
}

func OwnedBy(owner User) OpenOpt {
	return func(opts *Dataset) {
		opts.owner = owner
	}
}

func Named(name string) OpenOpt {
	return func(opts *Dataset) {
		opts.name = name
	}
}

func WithLogger(logger log.Logger) OpenOpt {
	return func(opts *Dataset) {
		opts.log = logger
	}
}

// TODO: test that this works
// OpenWithMissingExtensions will open the dataset,
// even if the server does not have the correct extensions
// installed.  This should only be used when failing to have an extension
// effects the entire server (like startup)
func OpenWithMissingExtensions() OpenOpt {
	return func(opts *Dataset) {
		opts.allowMissingExtensions = true
	}
}