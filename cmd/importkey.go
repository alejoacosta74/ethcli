/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	qcommon "github.com/qtumproject/qtool/lib/common"
	"github.com/spf13/cobra"
)

// importkeyCmd represents the importkey command
var importkeyCmd = &cobra.Command{
	Use:   "importkey",
	Short: "imports a private key into the keystore of the node",
	Long: `Imports an ethereum private key (in hexadecimal format), into 
the keystore of the node.
For example:
ethcli importkey 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
`,
	RunE: runImportKey,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(importkeyCmd)

}

func runImportKey(cmd *cobra.Command, args []string) error {
	privateKeyHex := args[0]
	privateKeyHex = qcommon.RemoveHexPrefix(privateKeyHex)

	result, err := client.ImportRawKey(privateKeyHex)
	if err != nil {
		return err
	}

	fmt.Printf("Result: \n%s\n", result)

	return nil
}
