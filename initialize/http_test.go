package initialize

import "testing"

func Test_updateInitConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_updateInitConfig"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getAccounts()
		})
	}
}
