package utils

import (
	"encoding/hex"
	"errors"
	"math/big"
	"testing"

	"github.com/kwilteam/kwil-db/cmd/internal/display"
	"github.com/kwilteam/kwil-db/cmd/kwil-cli/config"
	"github.com/kwilteam/kwil-db/pkg/client/types"
	"github.com/kwilteam/kwil-db/pkg/crypto"
	"github.com/kwilteam/kwil-db/pkg/transactions"
	"github.com/stretchr/testify/assert"
)

func Test_respStr(t *testing.T) {
	s := respStr("pong")

	bs, err := s.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "pong", string(bs))

	sb, err := s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"message":"pong"}`, string(sb))
}

func Example_respStr_text() {
	display.Print(respStr("pong"), nil, "text")
	// Output:
	// pong
}

func Example_respStr_json() {
	display.Print(respStr("pong"), nil, "json")
	// Output:
	// {
	//   "result": {
	//     "message": "pong"
	//   },
	//   "error": ""
	// }
}

func Example_respStr_json_withError() {
	err := errors.New("an error")
	display.Print(respStr("pong"), err, "json")
	// Output:
	// {
	//   "result": "",
	//   "error": "an error"
	// }
}

func getExampleTxQueryResponse() *types.TcTxQueryResponse {
	secp256k1EpSigHex := "cb3fed7f6ff36e59054c04a831b215e514052753ee353e6fe31d4b4ef736acd6155127db555d3006ba14fcb4c79bbad56c8e63b81a9896319bb053a9e253475800"
	secp256k1EpSigBytes, _ := hex.DecodeString(secp256k1EpSigHex)
	secpSig := crypto.Signature{
		Signature: secp256k1EpSigBytes,
		Type:      crypto.SignatureTypeSecp256k1Personal,
	}

	rawPayload := transactions.ActionExecution{
		DBID:   "xf617af1ca774ebbd6d23e8fe12c56d41d25a22d81e88f67c6c6ee0d4",
		Action: "create_user",
		Arguments: [][]string{
			{"foo", "32"},
		},
	}

	payloadRLP, err := rawPayload.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return &types.TcTxQueryResponse{
		Hash:   []byte("1024"),
		Height: 10,
		Tx: transactions.Transaction{
			Body: &transactions.TransactionBody{
				Payload:     payloadRLP,
				PayloadType: rawPayload.Type(),
				Fee:         big.NewInt(100),
				Nonce:       10,
				Salt:        []byte("salt"),
				Description: "This is a test transaction for cli",
			},
			Serialization: transactions.SignedMsgConcat,
			Signature:     &secpSig,
		},
		TxResult: transactions.TransactionResult{
			Code:      0,
			Log:       "This is log",
			GasUsed:   10,
			GasWanted: 10,
			Data:      nil,
			Events:    nil,
		},
	}
}

func Example_respTxQuery_text() {
	display.Print(&respTxQuery{Msg: getExampleTxQueryResponse()}, nil, "text")
	// Output:
	// Transaction ID: 31303234
	// Status: success
	// Height: 10
	// Log: This is log
}

func Example_respTxQuery_json() {
	display.Print(&respTxQuery{Msg: getExampleTxQueryResponse()}, nil, "json")
	// Output:
	// {
	//   "result": {
	//     "hash": "31303234",
	//     "height": 10,
	//     "tx": {
	//       "Signature": {
	//         "signature_bytes": "yz/tf2/zblkFTASoMbIV5RQFJ1PuNT5v4x1LTvc2rNYVUSfbVV0wBroU/LTHm7rVbI5juBqYljGbsFOp4lNHWAA=",
	//         "signature_type": "secp256k1_ep"
	//       },
	//       "Body": {
	//         "Description": "This is a test transaction for cli",
	//         "Payload": "AAH4ULg5eGY2MTdhZjFjYTc3NGViYmQ2ZDIzZThmZTEyYzU2ZDQxZDI1YTIyZDgxZTg4ZjY3YzZjNmVlMGQ0i2NyZWF0ZV91c2VyyMeDZm9vgjMy",
	//         "PayloadType": "execute_action",
	//         "Fee": 100,
	//         "Nonce": 10,
	//         "Salt": "c2FsdA=="
	//       },
	//       "Serialization": "concat",
	//       "Sender": null
	//     },
	//     "tx_result": {
	//       "log": "This is log",
	//       "gas_used": 10,
	//       "gas_wanted": 10
	//     }
	//   },
	//   "error": ""
	// }
}

func Example_respKwilCliConfig_text() {
	pk, err := crypto.GenerateSecp256k1Key()
	if err != nil {
		panic(err)
	}

	display.Print(&respKwilCliConfig{
		cfg: &config.KwilCliConfig{
			PrivateKey:  pk,
			GrpcURL:     "localhost:9090",
			TLSCertFile: "",
		},
	}, nil, "text")
	// Output:
	// PrivateKey: ***
	// GrpcURL: localhost:9090
	// TLSCertFile:
}

func Example_respKwilCliConfig_json() {
	pk, err := crypto.GenerateSecp256k1Key()
	if err != nil {
		panic(err)
	}

	display.Print(&respKwilCliConfig{
		cfg: &config.KwilCliConfig{
			PrivateKey:  pk,
			GrpcURL:     "localhost:9090",
			TLSCertFile: "",
		},
	}, nil, "json")
	// Output:
	// {
	//   "result": {
	//     "private_key": "***",
	//     "grpc_url": "localhost:9090",
	//     "tls_cert_file": ""
	//   },
	//   "error": ""
	// }
}