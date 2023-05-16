package staking

import (
	"context"
	"fmt"
	"me-test/client"
	"me-test/config"
	"testing"
)

var Ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

// defer cancel()
var C, err = client.NewCmClient("")

func TestNewKeeper(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewKeeper"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewKeeper()
			fmt.Println(got)
			defer got1()
		})
	}
}
