package tools

import (
	"strings"
	"testing"
)

func TestGetAccAddr(t *testing.T) {
	pk := GenPriKey()

	type args struct {
		privKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{pk}, "cosmos", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccAddr(tt.args.privKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotStr := got.String()
			if strings.Contains(gotStr, tt.want) {
				t.Log(gotStr)
			} else {
				t.Errorf("GetAccAddr() not contains got = %v, want %v", got, tt.want)
			}
		})
	}
}
