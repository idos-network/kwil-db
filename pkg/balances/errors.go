package balances

import "fmt"

var (
	ErrInsufficientFunds = fmt.Errorf("insufficient funds")
	ErrConvertToBigInt   = fmt.Errorf("could not convert to big int")
	ErrInvalidNonce      = fmt.Errorf("invalid nonce")
	errAccountNotFound   = fmt.Errorf("account not found")
)
