package ethcli

import (
	"github.com/alejoacosta74/ethcli/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// Returns a slice of transactions given a slice of transaction hashes
func (c *EthClient) GetTransactionsByHash(hashes []common.Hash) (txs []*types.Transaction, err error) {
	for _, hash := range hashes {
		log.With("cmd", "gettx").Debugf("Attempting to get tx: %s ", hash.Hex())
		tx, isPending, err := c.TransactionByHash(c.ctx, hash)
		if err != nil {
			log.With("cmd", "gettx").Debugf("error getting tx for hash: %s, error: %s", hash.Hex(), err.Error())
			return nil, errors.New("error getting tx for hash: " + hash.Hex() + ", error: " + err.Error())
		}
		if isPending {
			log.With("cmd", "gettx").Debugf("Transaction %s is pending", hash.Hex())
		}
		txs = append(txs, tx)
		log.With("cmd", "gettx").Debugf("Tx Gas :%v, GasPrice: %s, Hash: %s, Nonce: %v, Value: %s, to: %s \n", tx.Gas(), tx.GasPrice(), tx.Hash().Hex(), tx.Nonce(), tx.Value(), tx.To().Hex())
	}
	return txs, nil
}
