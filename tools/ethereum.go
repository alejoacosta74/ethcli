package tools

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// VerifyTxSignature() verifies the signature of a signed ethereum transaction
// against a public key.
//
// Params:
// - tx: the signed transaction types.Transaction
// - pubKeyHex: the public key hex string of the signer
func VerifyTxSignature(tx *types.Transaction, pubKeyHex string) bool {

	//1. Get the signature and signer from the signed transaction
	signature, signer := GetSignatureAndSigner(tx)

	//2. Recreate the raw transaction from the signed transaction
	var rawTx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data())

	//3. RLP encode the raw transaction
	var buf bytes.Buffer
	if err := rlp.Encode(&buf, rawTx); err != nil {
		fmt.Printf("error encoding raw tx: %v\n", err)
		return false
	}
	//4. Hash the raw transaction
	rawTxHashed := signer.Hash(rawTx)
	digest := rawTxHashed.Bytes()

	//6. Recover the public key from the signature and message digest
	recoveredPubKey, err := crypto.SigToPub(digest, signature)
	if err != nil {
		return false
	}

	//7. Convert the recovered public key to a hex string
	recoveredPubKeyHex := hex.EncodeToString(crypto.FromECDSAPub(recoveredPubKey))

	//8. Compare the recovered public key hex string with the expected public key hex string
	return recoveredPubKeyHex == pubKeyHex
}

// DecodeRawTX decodes a raw ethereum RLP encoded transaction
// and returns a go-ethereum types.Transaction
func DecodeRawTx(rawtx string) (*types.Transaction, error) {
	rawtx = RemoveHexPrefix(rawtx)

	rawtxBytes, err := hex.DecodeString(rawtx)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding hex rawtx")
	}

	var tx = &types.Transaction{}

	err = rlp.DecodeBytes(rawtxBytes, tx)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding rlp rawtx")
	}

	return tx, nil
}

// GetFromAddress extracts the signer address from the public key
// of a signed ethereum transaction
func GetFromAddress(tx *types.Transaction) (*common.Address, error) {

	signature, signer := GetSignatureAndSigner(tx)

	recoveredPubKey, err := crypto.SigToPub(signer.Hash(tx).Bytes(), signature)
	if err != nil {
		return nil, errors.Wrap(err, "error recovering pubkey")
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)

	return &recoveredAddress, nil

}

// GetSignatureAndSigner returns the signature and signer from
// a signed ethereum transaction
func GetSignatureAndSigner(tx *types.Transaction) (signature []byte, signer types.Signer) {
	v, r, s := tx.RawSignatureValues()

	signature = make([]byte, 65)
	copy(signature[32-len(r.Bytes()):32], r.Bytes())
	copy(signature[64-len(s.Bytes()):64], s.Bytes())

	if tx.Protected() {
		signer = types.NewEIP155Signer(tx.ChainId())
		signature[64] = byte(v.Uint64() - 35 - 2*tx.ChainId().Uint64())
	} else {
		signer = types.HomesteadSigner{}
		signature[64] = byte(v.Uint64() - 27)
	}

	return signature, signer
}

// NewEthereumTx() creates a new unsigned ethereum transaction
// with default parameters.
func NewEthereumTx() *types.Transaction {
	nonce := uint64(0)
	gasPrice := big.NewInt(20000000000)
	gasLimit := uint64(21000)
	toAddress := common.HexToAddress("0x71517f86711b4bff4d789ad6fee9a58d8af1c6bb")
	amount := big.NewInt(1000000)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	return tx
}

// SignEthereumTx() signs an ethereum transaction with a private key.
//
// Params:
// - tx: the unsigned transaction types.Transaction
// - signer: the signer types.Signer
// - privKeyHex: the private key hex string of the signer
func SignEthereumTx(tx *types.Transaction, signer types.Signer, privKeyHex string) (*types.Transaction, error) {
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

// IsValidWeiAmount() checks if a string represents a valid wei amount.
func IsValidWeiAmount(weiAmount string) bool {
	if len(weiAmount) == 0 || weiAmount == "0" {
		return false
	}
	_, ok := new(big.Int).SetString(weiAmount, 10)
	return ok
}
