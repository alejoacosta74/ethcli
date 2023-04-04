package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddr(t *testing.T) {
	privKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	address := "Address: 0x96216849c49358B10257cb55b28eA603c874b05E"

	tests := []struct {
		name    string
		args    []string
		output  string
		isError bool
	}{
		{
			name:    "Valid key",
			args:    []string{privKey},
			output:  address,
			isError: false,
		},
		{
			name:    "Invalid key",
			args:    []string{"invalid"},
			output:  "",
			isError: true,
		},
		{
			name:    "Invalid number of arguments",
			args:    []string{"invalid", "invalid"},
			output:  "",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := new(bytes.Buffer)
			rootCmd.SetOut(actual)
			rootCmd.SetErr(actual)
			args := append([]string{"getaddr"}, tt.args...)
			args = append(args, "--node", "http://127.0.0.1:5777")
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
