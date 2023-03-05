/*
Copyright © 2022 Alejo Acosta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alejoacosta74/ethereum-client/ethcli"
	"github.com/alejoacosta74/ethereum-client/log"
	"github.com/alejoacosta74/gologger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ethcli",
	Short: "Ethereum CLI",
	Long: `Use this CLI to interact with the Ethereum blockchain.
Use ethcli --help to see the list of available sub commands.
Must be used with a running Ethereum node with JSON-RPC API enabled.`,
	PersistentPreRunE: runPersistenPreRunE,
}

var (
	ethNode  string
	debug    bool
	cfgFile  string
	client   *ethcli.EthClient
	logger   *gologger.Logger
	logLevel string
)

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	// set viper default values
	// viper.SetDefault("node", "http://http://127.0.0.1:8080/proxy")
	viper.SetDefault("debug", false)
	viper.SetDefault("loglevel", "info")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&ethNode, "node", "n", "http://127.0.0.1:8080/proxy", "Ethereum node URL")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "", "log level (trace, debug, info, warn, error, fatal, panic")

	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))

	cobra.MarkFlagFilename(rootCmd.PersistentFlags(), "config")

	viper.BindPFlag("node", rootCmd.PersistentFlags().Lookup("node"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

}

func runPersistenPreRunE(cmd *cobra.Command, args []string) error {
	err := loadConfig()
	if err != nil {
		return err
	}
	if !viper.IsSet("node") {
		return errors.New("ethereum node URL is not set")
	}

	debug = viper.GetBool("debug")
	err = setLogger()
	if err != nil {
		return errors.Wrap(err, "error setting logger")
	}
	client, err = ethcli.NewEthClient(viper.GetString("node"))
	if err != nil {
		return errors.Wrap(err, "error creating eth client")
	}
	log.With("module", "rootcmd").Debugf("using config file: %s\n", viper.ConfigFileUsed())

	return nil

}

// loadConfig loads the ethcli configuration from a
// file or from env variables
func loadConfig() error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file - %s", err)
		}
	} else {
		// Default location for config file is $HOME/.ethcli
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return errors.Wrap(err, "error getting user home directory")
		} else {
			configFile := homeDir + "/.ethcli/config.yaml"
			viper.SetConfigFile(configFile)

			if err := viper.ReadInConfig(); err != nil {
				return errors.Wrap(err, "failed to read config file")
			}
		}
	}

	viper.AutomaticEnv()
	return nil
}

func setLogger() error {
	var err error
	level := viper.GetString("loglevel")
	switch level {
	case "trace":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.TraceLevel))
	case "debug":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.DebugLevel))
	case "info":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.InfoLevel))
	case "warn":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.WarnLevel))
	case "error":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.ErrorLevel))
	case "fatal":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.FatalLevel))
	case "panic":
		logger, err = gologger.NewLogger(gologger.WithLevel(gologger.PanicLevel))
	default:
		logger, err = gologger.NewLogger(gologger.WithNullLogger())
	}
	if err != nil {
		return err
	}
	log.SetLogger(logger)
	return nil
}
