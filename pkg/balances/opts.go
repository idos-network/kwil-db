package balances

import "kwil/pkg/log"

type balancesOpts func(*AccountStore)

func WithPath(path string) balancesOpts {
	return func(ar *AccountStore) {
		ar.path = path
	}
}

func Wipe() balancesOpts {
	return func(ar *AccountStore) {
		ar.wipe = true
	}
}

func WithLogger(logger log.Logger) balancesOpts {
	return func(ar *AccountStore) {
		ar.log = logger
	}
}