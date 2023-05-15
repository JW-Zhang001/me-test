package staking

import "testing"

func TestDelegate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDelegate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delegate()
		})
	}
}

func Test_initStakeClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_initStakeClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initStakeClient()
		})
	}
}

func TestUndelegate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestUndelegate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Undelegate()
		})
	}
}
