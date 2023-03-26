package ethcli

import (
	"math/big"

	"github.com/alejoacosta74/ethcli/log"
	"github.com/alejoacosta74/ethcli/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (c *EthClient) Send(key string, receiver string, weiAmount string) (signedTx *types.Transaction, err error) {

	// check if key is valid
	key = tools.RemoveHexPrefix(key)
	if !tools.IsValidHexPrivKey(key) {
		log.With("module", "ethcli").Debugf("invalid private key: %s", key)
		return nil, errors.New("invalid private key: " + key + " (must be 64 hex characters)")
	}
	// check if receiver is valid
	if !tools.IsValidHexAddress(receiver) {
		log.With("module", "ethcli").Debugf("invalid receiver address: %s", receiver)
		return nil, errors.New("invalid receiver address: " + receiver + " (must be 40 hex characters)")
	}
	// check if amount is valid
	if !tools.IsValidWeiAmount(weiAmount) {
		log.With("module", "ethcli").Debugf("invalid wei amount: %s", weiAmount)
		return nil, errors.New("invalid wei amount: " + weiAmount + " (must be a number)")
	}

	// convert key to ecdsa
	privateKey, err := tools.ConvertPrivateKey(key)
	if err != nil {
		log.With("module", "ethcli").Debugf("error converting private key %s to ecdsa: %s", key, err.Error())
		return nil, err
	}

	// get "from" address from private key
	fromAddress, err := tools.GetAddressFromPrivKey(privateKey)
	if err != nil {
		log.With("module", "ethcli").Debugf("error getting address from privkey %s : %s", key, err.Error())
		return nil, err
	}

	// get nonce
	nonce, err := c.PendingNonceAt(c.ctx, fromAddress)
	if err != nil {
		log.With("module", "ethcli").Debugf("Could not get nonce for address %s (error %s)", fromAddress, err.Error())
		return nil, err
	}

	// create transaction
	amount, ok := new(big.Int).SetString(weiAmount, 10)
	if !ok {
		log.With("module", "ethcli").Debugf("Could not convert wei amount of %s to big.Int", weiAmount)
		return nil, errors.New("Could not convert wei amount to big.Int")
	}

	// TODO: get gas limit from config
	gasLimit := uint64(21000) // in units

	// get gas price
	gasPrice, err := c.SuggestGasPrice(c.ctx)
	if err != nil {
		log.With("module", "ethcli").Debugf("could not get gas price: %s", err.Error())
		return nil, errors.Wrapf(err, "could not get gas price")
	}
	log.With("module", "ethcli").Debugf("gasPrice: %d", gasPrice)

	to := common.HexToAddress(receiver)

	// create transaction
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	// get network id
	netID, err := c.NetworkID(c.ctx)
	if err != nil {
		log.With("module", "ethcli").Debugf("could not get network id: %s ", err.Error())
		return nil, err

	}
	log.With("module", "ethcli").Debugf("netID: %d", netID)

	// sign transaction
	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(netID), privateKey)
	if err != nil {
		log.With("module", "ethcli").Debugf("error signing tx: %s ", err.Error())
		return nil, err
	}

	// send transaction
	err = c.SendTransaction(c.ctx, signedTx)
	if err != nil {
		log.With("module", "ethcli").Debugf("error sending transaction: %s", err.Error())
		return nil, err
	}

	log.With("module", "ethcli").Debugf("tx sent: %s\n", signedTx.Hash().Hex())

	return signedTx, nil
}
