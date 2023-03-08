package tools

import (
	"bytes"
	"net/http"

	"github.com/alejoacosta74/ethcli/log"
)

// CreateHTTPRPCRequest is a helper function to create a new JSON RPC http request
// with the given method and arguments.
func CreateHTTPRPCRequest(url string, method string, args ...string) (*http.Request, error) {
	params := []string{}
	for _, arg := range args {
		arg = `"` + arg + `"`
		params = append(params, arg)
	}
	log.With("module", "rpc").Debugf("Creating json request with method: %s and params: %+v", method, params)
	jsonReq, err := CreateJSONRequest(method, params...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
