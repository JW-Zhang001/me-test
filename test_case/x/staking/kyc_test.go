package staking

import (
	"testing"

	_ "me-test/initialize"
)

func TestNewKyc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewKyc()
		})
	}
}
