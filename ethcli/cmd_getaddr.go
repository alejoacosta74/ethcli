package ethcli

import (
	"github.com/alejoacosta74/ethcli/tools"
)

// GetAdressFromPrivateKey returns the ethereum address from
// a private key encoded as a hex string
func (c *EthClient) GetAdressFromPrivateKey(privKeyHex string) (string, error) {
	privKey, err := tools.ConvertPrivateKey(privKeyHex)
	if err != nil {
		return "", err
	}
	address, err := tools.GetAddressFromPrivKey(privKey)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}
