package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	jsonrpcVersion = "2.0"
)

// http://www.jsonrpc.org/specification#request_object
type JSONRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	// Params  []interface{} `json:"params"`
	Params []json.RawMessage `json:"params"`
	ID     int               `json:"id"`
}

func NewJSONRPCRequest(id int, method string, params []json.RawMessage) *JSONRPCRequest {
	return &JSONRPCRequest{
		JSONRPC: jsonrpcVersion,
		ID:      id,
		Method:  method,
		Params:  params,
	}
}

// http://www.jsonrpc.org/specification#response_object
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	// Error   string `json:"error"`
	Error *JSONRPCError `json:"error"`
	ID    int           `json:"id"`
}

// http://www.jsonrpc.org/specification#error_object
type JSONRPCError struct {
	JSONRPC string `json:"jsonrpc"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data    string `json:"data"`
	Data json.RawMessage `json:"data"`
	ID   int             `json:"id"`
}

func NewJSONRPCResponse(id int, result json.RawMessage, err *JSONRPCError) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: jsonrpcVersion,
		Result:  result,
		Error:   err,
		ID:      id,
	}
}

func NewJSONRPCResponseError(id, code int, message string, data json.RawMessage) *JSONRPCError {
	return &JSONRPCError{
		JSONRPC: jsonrpcVersion,
		ID:      id,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// CreateJSONRequest creates a JSON RPC request with the given method and params
// and returns the JSON encoded request as a byte array.
func CreateJSONRequest(method string, paramsStr ...string) ([]byte, error) {
	params := make([]json.RawMessage, len(paramsStr))
	for i, param := range paramsStr {
		params[i] = json.RawMessage(param)
	}
	// TODO: use btcjson function
	//btcjson.NewRequest(1, method, params)
	jsonReq := NewJSONRPCRequest(1, method, params)
	return json.Marshal(jsonReq)
}

// ReadJSONResult reads the JSON response from the given http response and
// unmarshals the result into the given result interface.
func ReadJSONResult(resp *http.Response, result interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error reading body: " + err.Error())
	}
	var jsonResp *JSONRPCResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return errors.New("error unmarshaling JSON response: " + err.Error())
	}
	if jsonResp.Error != nil {
		return errors.New("JSON Response with error: " + jsonResp.Error.Message + " (code: " + fmt.Sprint(jsonResp.Error.Code) + ")")
	}
	err = json.Unmarshal(jsonResp.Result, result)
	if err != nil {
		return errors.New("error unmarshaling JSON result: " + err.Error())
	}
	return nil
}
