package ethcli

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alejoacosta74/ethcli/internal/mocks"
	"github.com/alejoacosta74/ethcli/tools"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	/*
		Scope: test the Send method of the ethcli client

		What to test?
		1. The signed raw transaction is a valid tx and can be decoded
		2. The signed raw transaction contains the correct values (i.e. to, value, data)
		3. The signature is valid
		4. The signer is the correct one
	*/

	// 1. Create mocked responses and mocked ethereum node
	mocks.SetResponse(t, "eth_getTransactionCount", "eth_getTransactionCount_response.json")
	mocks.SetResponse(t, "eth_gasPrice", "eth_gasPrice_response.json")
	mocks.SetResponse(t, "net_version", "net_version_response.json")
	mocks.SetResponse(t, "eth_sendRawTransaction", "eth_sendRawTransaction_response_1.json")

	responses := mocks.GetMockedResponses()
	mockedNode := mocks.NewMockEthNode(&reqRecorder, responses)
	defer mockedNode.Close()

	// 2. Create ethcli client
	c, err := NewEthClient(mockedNode.URL)
	assert.NoError(t, err)

	// 3. Call Send
	privKey := "85cbc7b1adfe877051d746c3996a01c2bc3e7a6988490439b1f4b4c2b465322d"
	to := "0x71517f86711b4bff4d789ad6fee9a58d8af1c6bb"
	amount := "1000000000000000000"
	_, err = c.Send(privKey, to, amount)
	assert.NoError(t, err)

	// 4. Verify the received encoded raw tx can be decoded and is valid
	var req btcjson.Request
	json.Unmarshal(reqRecorder, &req)
	if len(req.Params) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(req.Params))
	}
	var params string
	err = json.Unmarshal(req.Params[0], &params)
	if err != nil {
		t.Fatalf("Error decoding params: %v", err)
	}
	encodedTx := string(params)
	decodedTx, err := c.DecodeRawTx(encodedTx)
	assert.NoError(t, err)

	// 5. Check the decoded tx contains the correct values
	gotToAddr := strings.ToLower(decodedTx.To().Hex())
	assert.Equal(t, to, gotToAddr)
	assert.Equal(t, amount, decodedTx.Value().String())

	// 6. Verify the decoded tx is signed and the signature is valid
	pubKey, _ := tools.PrivToPubKey(privKey)
	verified := tools.VerifyTxSignature(decodedTx, pubKey)
	assert.True(t, verified)

}
