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

func TestGetWhichYearBlockReward(t *testing.T) {
	type args struct {
		year uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "case1", args: args{1}, want: 793},
		{name: "case2", args: args{2}, want: 397},
		{name: "case3", args: args{3}, want: 199},
		{name: "case4", args: args{4}, want: 100},
		{name: "case5", args: args{5}, want: 50},
		{name: "case6", args: args{6}, want: 25},
		{name: "case7", args: args{7}, want: 13},
		{name: "case8", args: args{8}, want: 7},
		{name: "case9", args: args{9}, want: 4},
		{name: "case10", args: args{10}, want: 2},
		{name: "case11", args: args{11}, want: 1},
		{name: "case12", args: args{12}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWhichYearBlockReward(tt.args.year); got != tt.want {
				t.Errorf("GetWhichYearBlockReward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockRewardEqualZero(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BlockRewardEqualZero()
		})
	}
}
