package ethcli

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"

	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type EthClient struct {
	*ethclient.Client
	ctx context.Context
	url string
}

func NewEthClient(url string) (*EthClient, error) {
	url = strings.TrimSpace(url)
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, errors.New("Error connecting to ethereum node: " + err.Error())
	}

	c := &EthClient{
		Client: client,
		ctx:    context.Background(),
		url:    url,
	}
	return c, nil

}

func (c *EthClient) PrintPretty(msg string, source interface{}) {
	output, err := json.Marshal(source)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
	}
	if len(output) > 0 {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, output, "", "    "); err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
		} else {
			fmt.Printf("\n%s :\n%s\n", msg, prettyJSON.String())

		}
	}
}
