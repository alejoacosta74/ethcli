package mocks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "github.com/alejoacosta74/ethereum-client/internal/testutils"
	"github.com/alejoacosta74/ethereum-client/log"
	"github.com/btcsuite/btcd/btcjson"
)

// mockResponses is a map of method names to json responses
// that is used to mock a qtumd server
var mockResponses = make(map[string]*btcjson.Response)

func GetMockedResponses() map[string]*btcjson.Response {
	return mockResponses
}

func SetResponse(t *testing.T, method string, filename string) {
	response := utils.LoadFile(t, filename)
	var jsonResp = new(btcjson.Response)
	err := json.Unmarshal(response, jsonResp)
	if err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}
	mockResponses[method] = jsonResp
}

type MockedEthNode struct {
	*httptest.Server
}

// NewMockedEthNode creates a new mock qtumd server with the given handler.
// Params:
//   - request: a buffer used to record the request sent to the server
//   - responses: a map of method names to json responses
func NewMockEthNode(reqRecorder *[]byte, responses map[string]*btcjson.Response) *MockedEthNode {
	handlerFunc := responseHandler(reqRecorder, responses)
	handler := http.HandlerFunc(handlerFunc)
	return &MockedEthNode{
		Server: httptest.NewServer(handler),
	}
}

// ResponseHandler creates a handler function that for a given JSON RPC method
// returns the corresponding response. The handler function is used to create
// a mocked Ethereum server.
func responseHandler(reqRecorder *[]byte, responses map[string]*btcjson.Response) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// if reqRecorder is not nil, record the incoming request
		// for later inspection during testing
		if reqRecorder != nil {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			*reqRecorder = append(body, []byte{}...)

			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		var jsonreq btcjson.Request
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&jsonreq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// find response for method
		method := jsonreq.Method
		log.With("module", "mockedserver").Debugf("received method: %s", method)
		jResp, ok := responses[method]
		if !ok {
			http.Error(w, "method not found: "+method, http.StatusInternalServerError)
			return
		}
		// write response
		w.Header().Set("Content-Type", "application/json")

		resp, err := json.Marshal(jResp)
		if err != nil {
			panic(err)
		}
		w.Write(resp)
	}
}
