package staking

import "testing"

func TestNewRegion(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewRegion"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRegion()
		})
	}
}
