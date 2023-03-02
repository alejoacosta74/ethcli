package ethcli

import "testing"

func TestGetAddressFromPrivKey(t *testing.T) {
	c, err := NewEthClient("http://localhost:8545")
	handleFatalError(t, err)
	privKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	want := "0x96216849c49358B10257cb55b28eA603c874b05E"
	addr, err := c.GetAdressFromPrivateKey(privKey)
	handleFatalError(t, err)
	if addr != want {
		t.Fatalf("Address mismatch: got %s, want %s", addr, want)
	}
}
