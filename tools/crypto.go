package tools

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

// PrivToPubKey converts a ECDSA private key to a Ethereum public key.
func PrivToPubKey(privKeyHex string) (string, error) {
	privKeyHex, err := sanitizePrivateKey(privKeyHex)
	if err != nil {
		return "", err
	}
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return "", errors.Wrapf(err, "error converting private key to ECDSA: %s", privKeyHex)
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	return hex.EncodeToString(pubKeyBytes), nil
}

// PubKeyToAddress converts a public key to an Ethereum address.
func PubKeyToAddress(pubKeyHex string) (string, error) {
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", err
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*pubKey)
	return address.Hex(), nil
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

// ConvertPrivateKey() converts a private key (hex) string to an ecdsa.PrivateKey
func ConvertPrivateKey(key string) (privateKey *ecdsa.PrivateKey, err error) {

	key = RemoveHexPrefix(key)
	privateKey, err = crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
