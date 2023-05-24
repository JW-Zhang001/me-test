package testload

import "testing"

func TestDelegateBenchmark(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDelegateBenchmark"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelegateBenchmark()
		})
	}
}

func TestMixDelegateBenchmark(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMixDelegateBenchmark"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MixDelegateBenchmark()
		})
	}
}

func TestMixDelegateBenchmark2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMixDelegateBenchmark2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MixDelegateBenchmark2()
		})
	}
}
