/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send ether to an ethereum address",
	Long: `Sends the amount of ether to the specified receiver address. 
Usage: ethcli send <amount> <receiver address> {<sender private key> | <keystore filename>}

- Amount shall be specified in units of wei. 
- A private key must be provided with the --key flag or alternatively via a keystore file
with the flag --keystore <filename> .`,
	RunE:                       runSend,
	Args:                       cobra.ExactArgs(2),
	Example:                    `ethcli send 100 0x71517f86711B4BFf4D789Ad6FEE9a58D8AF1c6bB --key 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef`,
	SilenceUsage:               true,
	SuggestionsMinimumDistance: 2,
}

var (
	key      string
	keystore string
)

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&key, "key", "k", "", "Private key to sign the transaction")
	sendCmd.Flags().StringVarP(&keystore, "keystore", "s", "", "Keystore file that holds the private key")
	cobra.MarkFlagFilename(sendCmd.Flags(), "keystore")
	viper.BindPFlag("key", sendCmd.Flags().Lookup("key"))
	viper.BindPFlag("keystore", sendCmd.Flags().Lookup("keystore"))

}

func runSend(cmd *cobra.Command, args []string) error {

	if !viper.IsSet("key") && !viper.IsSet("keystore") {
		fmt.Println("Error: Private key not provided")
		os.Exit(1)
	}

	amount := args[0]
	receiver := args[1]

	// TODO: get the private key from the keystore file
	privKey := viper.GetString("key")

	tx, err := client.Send(privKey, receiver, amount)
	defer client.Close()

	if err != nil {
		return errors.Wrapf(err, "Error sending transaction")
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Succes!! Transaction hash: %s", tx.Hash().Hex())
	println()
	return nil

}
