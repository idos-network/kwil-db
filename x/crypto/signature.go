package crypto

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ec "github.com/ethereum/go-ethereum/crypto"
)

func Sign(data []byte, k *ecdsa.PrivateKey) (string, error) {
	hash := ec.Keccak256Hash(data)
	sig, err := ec.Sign(hash.Bytes(), k)
	if err != nil {
		return "", err
	}
	hx := hexutil.Encode(sig)
	return "00" + hx, nil
}

func ECDSAFromHex(hex string) (*ecdsa.PrivateKey, error) {
	return ec.HexToECDSA(hex)
}

func AddressFromPrivateKey(key string) (string, error) {
	ecdk, err := ECDSAFromHex(key)
	if err != nil {
		return "", err
	}

	caddr := ec.PubkeyToAddress(ecdk.PublicKey)
	return caddr.Hex(), nil
}

func CheckSignature(addr, sig string, data []byte) (bool, error) {
	switch sig[0:2] {
	case "00": // PK uncompressed
		return checkSignaturePkSECP256k1Uncompressed(addr, sig[2:], data)
	case "01": // Account uncompressed
		return checkSignatureAccountSECP256k1Uncompressed(addr, sig[2:], data)
	case "02": // Account compressed
		return checkSignatureAccountSECP256k1Compressed(addr, sig[2:], data)
	default:
		return false, fmt.Errorf("unknown signature indicator: %s", sig[0:2])
	}
}

// This would be for a signature generated by a private key.  Likely EVM chains
func checkSignaturePkSECP256k1Uncompressed(addr, sig string, data []byte) (bool, error) {
	hash := ec.Keccak256Hash(data)
	sb, err := hexutil.Decode(sig)
	if err != nil {
		return false, err
	}

	pubBytes, err := ec.Ecrecover(hash.Bytes(), sb)
	if err != nil {
		return false, err // I don't believe this can be reached
	}

	pub, err := ec.UnmarshalPubkey(pubBytes)
	if err != nil {
		return false, err // I don't believe this can be reached
	}
	derAddr := ec.PubkeyToAddress(*pub)

	return derAddr == common.HexToAddress(addr), nil
}

// checkSignatureAccount checks a signature that was generated by an account instead of a private key
// generally this would mean a signature from MetaMask / equivalent for EVM chains
func checkSignatureAccountSECP256k1Uncompressed(addr, sig string, data []byte) (bool, error) {
	sb, err := hexutil.Decode(sig)
	if err != nil {
		return false, err
	}
	msg := accounts.TextHash(data)
	if sb[ec.RecoveryIDOffset] == 27 || sb[ec.RecoveryIDOffset] == 28 {
		sb[ec.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recoveredPubKey, err := ec.SigToPub(msg, sb)
	if err != nil {
		return false, err
	}

	derAddr := ec.PubkeyToAddress(*recoveredPubKey)
	return addr == derAddr.Hex(), nil
}

// checkSignatureAccountSECP256k1Compressed checks a signature that was generated by an account instead of a private key
// this generally means that a cosmos chain was used (wallets like Keplr)
func checkSignatureAccountSECP256k1Compressed(addr, sig string, data []byte) (bool, error) {
	// TODO: implement
	return false, nil
}
