package ethcli

import (
	"testing"
)

// reqRecorder is a slice of bytes used to record requests
// sent to the mocked ethereum node by ethcli
var reqRecorder []byte

func handleFatalError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Fatal error: %+v", err)
	}
}
