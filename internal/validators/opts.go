package validators

import "github.com/kwilteam/kwil-db/core/log"

type ValidatorMgrOpt func(*ValidatorMgr)

// WithLogger sets the logger
func WithLogger(logger log.Logger) ValidatorMgrOpt {
	return func(v *ValidatorMgr) {
		v.log = logger
	}
}

func WithJoinExpiry(joinExpiry int64) ValidatorMgrOpt {
	return func(v *ValidatorMgr) {
		v.joinExpiry = joinExpiry
	}
}

func WithFeeMultiplier(multiplier int64) ValidatorMgrOpt {
	return func(v *ValidatorMgr) {
		v.feeMultiplier = multiplier
	}
}