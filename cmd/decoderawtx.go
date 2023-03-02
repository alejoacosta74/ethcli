/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// decoderawtxCmd represents the decoderawtx command
var decoderawtxCmd = &cobra.Command{
	Use:   "decoderawtx",
	Short: "Decode a ethereum RLP encoded raw transaction",
	Long: `Decodes a ethereum RLP encoded raw transaction and prints it
in a json style format. 
Usage: ethcli decoderawtx <raw tx hex string>
`,
	RunE:    runDecoderawtx,
	Args:    cobra.ExactArgs(1),
	Example: "ethcli decoderawtx 0xf86781f98504a817c8008252089471517f86711b4bff4d789ad6fee9a58d8af1c6bb6480822d45a06e5dbcace970fccd85faf1d7c940bc72b122ff8702327d8c88f3bcbf8daf4ac0a00d23cac1ce9515bf31e9c5b8ac2691c12517478d395dde0200fe45fa3447fff2",
}

func init() {
	rootCmd.AddCommand(decoderawtxCmd)
}

func runDecoderawtx(cmd *cobra.Command, args []string) error {
	decoded, err := client.DecodeRawTx(args[0])
	if err != nil {
		return err
	}

	client.PrintPretty("decoded tx", decoded)
	return nil
}
