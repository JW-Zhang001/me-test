package test_case

import "testing"

func Test_transferTx(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TransferTx()
		})
	}
}
