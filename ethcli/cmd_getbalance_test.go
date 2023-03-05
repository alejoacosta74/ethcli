package ethcli

import (
	"testing"

	"github.com/alejoacosta74/ethereum-client/internal/mocks"
)

func TestGetBalance(t *testing.T) {
	address := "0x96216849c49358B10257cb55b28eA603c874b05E"
	mocks.SetResponse(t, "eth_getBalance", "eth_getBalance_response.json")
	responses := mocks.GetMockedResponses()
	mockedNode := mocks.NewMockEthNode(&reqRecorder, responses)
	defer mockedNode.Close()

	c, err := NewEthClient(mockedNode.URL)
	handleFatalError(t, err)

	balance, err := c.GetBalance(address, nil)
	if err != nil {
		t.Fatalf("Error getting balance: %s", err)
	}
	want := "842796652117371" // 0x2fe84e3113d7b
	if balance != want {
		t.Fatalf("Balance mismatch: got %s, want %s", balance, want)
	}

}
