package api

import (
	"fmt"
	"io"
	"net/http"
)

// ApiUrl HTTP API Console
var (
	Host       = "http://192.168.0.206:10001"
	HttpClient = &http.Client{}
	Api        = map[string]RequestData{}
)

type RequestData struct {
	Url    string
	Method string
	Body   io.Reader
}

func init() {

	Api["accounts"] = RequestData{
		Url:    Host + "/cosmos/auth/v1beta1/accounts",
		Method: "get",
		Body:   nil,
	}

	Api["blockHeight"] = RequestData{
		Url:    Host + "/cosmos/base/tendermint/v1beta1/blocks/latest",
		Method: "get",
		Body:   nil,
	}

	fmt.Println("Api init success")
}
