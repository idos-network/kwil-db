package balances

import "github.com/kwilteam/kwil-db/pkg/log"

type AccountStoreOpts func(*AccountStore)

func WithPath(path string) AccountStoreOpts {
	return func(ar *AccountStore) {
		ar.path = path
	}
}

func WithLogger(logger log.Logger) AccountStoreOpts {
	return func(ar *AccountStore) {
		ar.log = logger
	}
}

func WithGasCosts(gas_enabled bool) AccountStoreOpts {
	return func(ar *AccountStore) {
		ar.gasEnabled = gas_enabled
	}
}

func WithNonces(nonces_enabled bool) AccountStoreOpts {
	return func(ar *AccountStore) {
		ar.noncesEnabled = nonces_enabled
	}
}

func WithDatastore(db Datastore) AccountStoreOpts {
	return func(ar *AccountStore) {
		ar.db = db
	}
}
