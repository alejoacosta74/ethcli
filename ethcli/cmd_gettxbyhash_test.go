package ethcli

import (
	"reflect"
	"testing"

	"github.com/alejoacosta74/ethcli/internal/mocks"
	utils "github.com/alejoacosta74/ethcli/internal/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestGetTxByHash(t *testing.T) {
	mocks.SetResponse(t, "eth_getTransactionByHash", "eth_getTransactionByHash_response_1.json")
	responses := mocks.GetMockedResponses()
	mockedNode := mocks.NewMockEthNode(&reqRecorder, responses)
	defer mockedNode.Close()

	c, err := NewEthClient(mockedNode.URL)
	handleFatalError(t, err)

	input := utils.LoadFile(t, "cmd_gettx_input_1.json")
	hash := common.BytesToHash(input)
	hashes := []common.Hash{hash}
	txs, err := c.GetTransactionsByHash(hashes)
	handleFatalError(t, err)
	if len(txs) != 1 {
		t.Fatalf("Expected 1 transaction, got %d", len(txs))
	}
	tx := txs[0]
	jsonResp := responses["eth_getTransactionByHash"]
	var wantTx = new(types.Transaction)
	err = wantTx.UnmarshalJSON(jsonResp.Result)

	handleFatalError(t, err)
	compareTXs(t, tx, wantTx)
}

func compareTXs(t *testing.T, got, want *types.Transaction) {
	if got.Hash() != want.Hash() {
		t.Fatalf("Hash mismatch: got %s, want %s", got.Hash(), want.Hash())
	}
	if got.Nonce() != want.Nonce() {
		t.Fatalf("Nonce mismatch: got %d, want %d", got.Nonce(), want.Nonce())
	}
	if got.Gas() != want.Gas() {
		t.Fatalf("Gas mismatch: got %d, want %d", got.Gas(), want.Gas())
	}
	if got.GasPrice().Cmp(want.GasPrice()) != 0 {
		t.Fatalf("GasPrice mismatch: got %d, want %d", got.GasPrice(), want.GasPrice())
	}
	if got.Value().Cmp(want.Value()) != 0 {
		t.Fatalf("Value mismatch: got %d, want %d", got.Value(), want.Value())
	}
	if got.Data() != nil && want.Data() != nil {
		if !reflect.DeepEqual(got.Data(), want.Data()) {
			t.Fatalf("Data mismatch: got %s, want %s", got.Data(), want.Data())
		}
	}
	if got.To() != nil && want.To() != nil {
		if got.To().Hex() != want.To().Hex() {
			t.Fatalf("To mismatch: got %s, want %s", got.To(), want.To())
		}
	}

}
