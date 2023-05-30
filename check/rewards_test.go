package check

import (
	"fmt"
	"testing"
)

func TestCalculateUserReward(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCalculateUserReward"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateAllUserReward()
			if err != nil {
				panic("TestCalculateUserReward error")
			}
			fmt.Println(got)
		})
	}
}

func TestCalculateTreasuryReward(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCalculateTreasuryReward"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTreasuryReward()
			if err != nil {
				panic("TestCalculateTreasuryReward error")
			}
			fmt.Println(got)
		})
	}
}
