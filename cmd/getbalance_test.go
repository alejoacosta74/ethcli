package cmd

import (
	"bytes"
	"testing"

	"github.com/alejoacosta74/ethereum-client/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	mocks.SetResponse(t, "eth_getBalance", "eth_getBalance_response.json")
	responses := mocks.GetMockedResponses()
	mockedNode := mocks.NewMockEthNode(nil, responses)
	defer mockedNode.Close()

	address := "0x96216849c49358B10257cb55b28eA603c874b05E"
	want := "Balance: 842796652117371 wei"

	tests := []struct {
		name    string
		args    []string
		output  string
		isError bool
	}{
		{
			name:    "Valid block",
			args:    []string{address, "12345", "-n " + mockedNode.URL},
			output:  want,
			isError: false,
		},
		{
			name:    "No block",
			args:    []string{address},
			output:  want,
			isError: false,
		},
		{
			name:    "Invalid block number",
			args:    []string{address, "I2E45"},
			output:  "",
			isError: true,
		},
		{
			name:    "Invalid block number",
			args:    []string{address, "0x11223344"},
			output:  "",
			isError: true,
		},
		{
			name:    "Invalid address",
			args:    []string{"invalid", "12345"},
			output:  "",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			args := append([]string{"getbalance"}, tt.args...)
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
