package lib

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ConvertPrivateKey() converts a private key (hex) string to an ecdsa.PrivateKey
func ConvertPrivateKey(key string) (privateKey *ecdsa.PrivateKey, err error) {

	key = RemoveHexPrefix(key)
	privateKey, err = crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// GetAddressFromPrivKey() returns the ethereum common.Address from a ecdsa.PrivateKey
func GetAddressFromPrivKey(privateKey *ecdsa.PrivateKey) (address common.Address, err error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA: " + privateKey.D.String())
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return

}

// PrivToPubKey() converts a ECDSA private key hex string to a public key hex string.
func PrivToPubKey(privKeyHex string) (string, error) {
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return "", err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	return hex.EncodeToString(pubKeyBytes), nil
}
