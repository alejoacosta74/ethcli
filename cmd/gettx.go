/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	qcommon "github.com/qtumproject/qtool/lib/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gettxbyhashCmd represents the gettxbyhash command
var gettxCmd = &cobra.Command{
	Use:   "gettx",
	Short: "Get transaction details",
	Long: `Gets the details of an ethereum transaction by issuing eth_getTransactionByHash . 
The hash of the transaction should be passed as an argument or in a file.
Usage: `,
	Example: `ethcli gettx 0x71517f86711B4BFf4D789Ad6FEE9a58D8AF1c6bB
ethcli gettx --file txHashes.json`,
	Args: cobra.MinimumNArgs(1),
	RunE: runGetTxByHash,
}

var (
	txHashesFile string
)

func init() {
	rootCmd.AddCommand(gettxCmd)
	gettxCmd.Flags().StringVarP(&txHashesFile, "file", "f", "", "File with the hash(es) of the transaction to retrieve")
	// viper.BindPFlag("hashesFile", gettxbyhashCmd.Flags().Lookup("file"))

}

func runGetTxByHash(cmd *cobra.Command, args []string) error {
	var txHashesStr []string
	if len(args) > 0 {
		txHashesStr = args
	} else {
		if txHashesFile != "" {
			viper.AddConfigPath(txHashesFile)
			viper.SetConfigName(txHashesFile)
			viper.SetConfigType("json")
			err := viper.ReadInConfig()
			if err != nil {
				return errors.New("error reading config file: " + err.Error())
			}
			txHashesStr = viper.GetStringSlice("hashes")
			// fmt.Println("txHashesStr from viper: ", txHashesStr)
		} else {
			return errors.New("no hashes provided")
		}
	}
	txHashes, err := readTxHashes(txHashesStr)
	if err != nil {
		return errors.New("error reading hashes: " + err.Error())
	}
	txs, err := client.GetTransactionsByHash(txHashes)
	if err != nil {
		return errors.New("error getting transactions: " + err.Error())
	}
	for _, tx := range txs {
		fmt.Printf("Tx Hash: %s, Gas :%v, GasPrice: %s,  Nonce: %v, Value: %s, to: %s \n", tx.Hash().Hex(), tx.Gas(), tx.GasPrice(), tx.Nonce(), tx.Value(), tx.To().Hex())
	}

	return nil
}

func readTxHashes(args []string) ([]common.Hash, error) {
	var hashes []common.Hash
	for _, arg := range args {
		hashStr := qcommon.RemoveHexPrefix(arg)
		if len(hashStr) != 64 {
			return nil, fmt.Errorf("invalid hash: %s", arg)
		}
		hash := common.HexToHash(arg)
		hashes = append(hashes, hash)
	}
	return hashes, nil
}
