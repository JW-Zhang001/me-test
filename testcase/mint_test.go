package testcase

import (
	"fmt"
	"testing"
)

func TestOneYearTotalBlocks1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := TestOneYearTotalBlocks()
			if !ok {
				panic("TestOneYearTotalBlocks1 error")
			}
			fmt.Println(got)
		})
	}
}

func TestInitBlockReward1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := TestInitBlockReward()
			if !ok {
				panic("TestOneYearTotalBlocks1 error")
			}
			fmt.Println(got)
		})
	}
}
