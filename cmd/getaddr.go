/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getaddrCmd represents the getaddr command
var getaddrCmd = &cobra.Command{
	Use:   "getaddr",
	Short: "Extracts the ethereum address from the private key",
	Long: `Extracts the ethereum address from the private key.
Usage: ethcli getaddr <priv key>`,
	RunE: runGetaddr,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(getaddrCmd)
}

func runGetaddr(cmd *cobra.Command, args []string) error {
	privateKey := args[0]
	addr, err := client.GetAdressFromPrivateKey(privateKey)
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Address: %s", addr)
	println()

	return nil
}
