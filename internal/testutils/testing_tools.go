package testutils

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/alejoacosta74/ethcli/log"
	"github.com/alejoacosta74/gologger"
)

var (
	Debug = false
)

func init() {
	// Read environment DEBUG variable
	environ := os.Environ()
	for _, env := range environ {
		if env == "DEBUG=true" {
			Debug = true
		}
	}
	SetLogger()
	log.With("module", "testutils").Debugf("Debug mode: %t", Debug)
}

// setLogger sets the logger to use while unit testing
func SetLogger() {
	var logger *gologger.Logger
	if Debug {
		logger, _ = gologger.NewLogger(gologger.WithDebugLevel(true))
	} else {
		logger, _ = gologger.NewLogger(gologger.WithNullLogger())
	}
	log.SetLogger(logger)
}

// LoadFile loads a file into a string for testing
func LoadFile(t *testing.T, filename string) []byte {
	fullpath := path.Join("../internal/testdata", filename)
	content, err := ioutil.ReadFile(fullpath)
	if err != nil {
		t.Fatal(err)
	}
	return content
}
