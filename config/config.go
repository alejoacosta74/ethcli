package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	EthereumNode     string `mapstructure:"node"`
	ConfigPath       string `mapstructure:"configpath"`
	PrivateKey       string `mapstructure:"privkey"`
	KeystorePath     string `mapstructure:"keystorepath"`
	KeystorePassword string `mapstructure:"keystorepassword"`
	ContractAddress  string `mapstructure:"contractaddress"`
	ContractABI      string `mapstructure:"contractabi"`
	Debug            bool   `mapstructure:"debug"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("ethcli")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // <--- read values from environment and override

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
