package ethcli

import (
	"github.com/alejoacosta74/ethcli/lib"
	"github.com/ethereum/go-ethereum/core/types"
)

// DecodeRawTX decodes a raw ethereum RLP encoded transaction
// and returns a go-ethereum types.Transaction
func (c *EthClient) DecodeRawTx(rawtx string) (*types.Transaction, error) {
	return lib.DecodeRawTx(rawtx)
}
