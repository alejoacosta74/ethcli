package ethcli

import (
	"net/http"

	"github.com/alejoacosta74/ethcli/log"
	"github.com/alejoacosta74/ethcli/tools"
	"github.com/pkg/errors"
)

// ImportRawKey() imports a private key into the keystore of the node
func (c *EthClient) ImportRawKey(key string) (string, error) {

	// check if key is valid
	key = tools.RemoveHexPrefix(key)
	if !tools.IsValidHexPrivKey(key) {
		log.With("module", "ethcli").Debugf("invalid private key: %s", key)
		return "", errors.New("invalid private key")
	}

	// 1. create http client
	client := &http.Client{}

	// 2. create request
	httpReq, err := tools.CreateHTTPRPCRequest(c.url, "personal_importRawKey", key, "password")
	if err != nil {
		return "", err
	}

	// 3. Send request
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return "", errors.New("Error sending request: " + err.Error())
	}

	var jsonRespResult string
	err = tools.ReadJSONResult(httpResp, &jsonRespResult)
	if err != nil {
		return "", err
	}

	// 5. Return result
	return jsonRespResult, nil
}
