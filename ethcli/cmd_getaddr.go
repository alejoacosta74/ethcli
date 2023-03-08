package ethcli

import (
	"github.com/alejoacosta74/ethcli/lib"
)

// GetAdressFromPrivateKey returns the ethereum address from
// a private key encoded as a hex string
func (c *EthClient) GetAdressFromPrivateKey(privKeyHex string) (string, error) {
	privKey, err := lib.ConvertPrivateKey(privKeyHex)
	if err != nil {
		return "", err
	}
	address, err := lib.GetAddressFromPrivKey(privKey)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}
