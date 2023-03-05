/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// getbalanceCmd represents the getbalance command
var getbalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "Returns the balance of the account of given address",
	Long: `Returns the balance of the account of given address.
The balance is returned in units of wei.
If a block number is given, it returns the balance of the account at the block height.
Otherwise, it returns the balance of the account at the 'latest' block. 
Usage: ethcli getbalance <address> [<block number>]
`,
	RunE:                       runGetBalance,
	Args:                       cobra.RangeArgs(1, 2),
	Example:                    `ethcli getbalance 0x71517f86711B4BFf4D789Ad6FEE9a58D8AF1c6bB 2540320`,
	SilenceUsage:               true,
	SuggestionsMinimumDistance: 2,
}

func init() {
	rootCmd.AddCommand(getbalanceCmd)
}

func runGetBalance(cmd *cobra.Command, args []string) error {

	address := args[0]
	var blockNumber *big.Int
	if len(args) == 2 {
		bkNum, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return errors.Wrapf(err, "Error parsing block number %s", args[1])
		}
		blockNumber = big.NewInt(bkNum)

	}

	result, err := client.GetBalance(address, blockNumber)
	defer client.Close()
	if err != nil {
		return errors.Wrapf(err, "Error getting balance for address %s", address)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Balance: %s wei", result)
	println()

	return nil
}
