package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivToPubKey(t *testing.T) {
	tests := []struct {
		name    string
		privKey string
		want    string
		error   bool
	}{
		{"empty", "", "", true},
		{"short", "123456789012345678901234567890123456789012345678901234567890123", "", true},
		{"long", "12345678901234567890123456789012345678901234567890123456789012345", "", true},
		{"invalid", "123456789012345678901234567890123456789012345678901234567890123g", "", true},
		{"valid", "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19", "049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrivToPubKey(tt.privKey)
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPubKeyToAddress(t *testing.T) {
	tests := []struct {
		name   string
		pubKey string
		want   string
		error  bool
	}{
		{"empty", "", "", true},
		{"short", "123456789012345678901234567890123456789012345678901234567890123", "", true},
		{"long", "12345678901234567890123456789012345678901234567890123456789012345", "", true},
		{"invalid", "123456789012345678901234567890123456789012345678901234567890123g", "", true},
		{"valid", "041dac3ce6613f97afc4e077d0830b16993b7d0d4119b2acc562d11ff893d86071264a4053458c533a7d3a5670a591b88677fef9435ebd4e120ec9a0b82216637a", "0xA6d2799a4b465805421bd10247386a708F01DB03", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PubKeyToAddress(tt.pubKey)
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIsValidHexAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{"empty", "", false},
		{"short", "0x123456789012345678901234567890123456789", false},
		{"long", "0x12345678901234567890123456789012345678901", false},
		{"invalid", "0x1234567890123456789012345678901234567890g", false},
		{"valid", "0xA6d2799a4b465805421bd10247386a708F01DB03", true},
		{"valid(no prefix)", "A6d2799a4b465805421bd10247386a708F01DB03", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsValidHexAddress(tt.address))
		})
	}
}
