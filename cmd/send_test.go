package cmd

import (
	"bytes"
	"testing"

	"github.com/alejoacosta74/ethcli/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	mocks.SetResponse(t, "eth_getTransactionCount", "eth_getTransactionCount_response.json")
	mocks.SetResponse(t, "eth_gasPrice", "eth_gasPrice_response.json")
	mocks.SetResponse(t, "net_version", "net_version_response.json")
	mocks.SetResponse(t, "eth_sendRawTransaction", "eth_sendRawTransaction_response_1.json")

	responses := mocks.GetMockedResponses()
	mockedNode := mocks.NewMockEthNode(nil, responses)
	defer mockedNode.Close()
	tests := []struct {
		name    string
		args    []string
		output  string
		isError bool
	}{
		{
			name:    "With mocked node and valid arguments",
			args:    []string{"2000", "0x96216849c49358B10257cb55b28eA603c874b05E", "--node=" + mockedNode.URL, "--key=85cbc7b1adfe877051d746c3996a01c2bc3e7a6988490439b1f4b4c2b465322d"},
			output:  "Succes!! Transaction hash: 0x5869e9bab50d0ecf02a83760deca63c1229ad9040dc552f1aa8b7c549a2062e6",
			isError: false,
		},
		// {
		// 	name:    "Missing private key",
		// 	args:    []string{"2000", "0x96216849c49358B10257cb55b28eA603c874b05E"},
		// 	output:  "",
		// 	isError: true,
		// },
		{
			name:    "Invalid arguments",
			args:    []string{"invalid"},
			output:  "",
			isError: true,
		},
		{
			name:    "Invalid number of arguments",
			args:    []string{"invalid"},
			output:  "",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			args := append([]string{"send"}, tt.args...)
			rootCmd.SetArgs(args)
			err := rootCmd.Execute()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.output, actual.String())
			}
		})
	}

}
