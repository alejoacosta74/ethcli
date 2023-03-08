package ethcli

import (
	"math/big"

	"github.com/alejoacosta74/ethcli/lib"
	"github.com/alejoacosta74/ethcli/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// GetBalance returns a string with the balance of
// an address in units of wei
func (c *EthClient) GetBalance(address string, blockNumber *big.Int) (string, error) {
	// check if address is valid
	if !lib.IsValidHexAddress(address) {
		log.With("module", "ethcli").Debugf("invalid address: %s", address)
		return "", errors.New("invalid address: " + address)
	}

	// get balance
	balance, err := c.BalanceAt(c.ctx, common.HexToAddress(address), nil)
	if err != nil {
		log.With("module", "ethcli").Debugf("Could not get balance for address %s (error %s)", address, err.Error())
		return "", err
	}

	return balance.String(), nil
}
