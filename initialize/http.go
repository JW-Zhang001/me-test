package initialize

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"me-test/config"
)

// HttpClient ApiUrl HTTP API Console
var HttpClient = &http.Client{}

// NewRequest Encapsulate the request and response
func NewRequest(method, url string, body io.Reader) (string, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("Request creation failure:", err)
		return "", err
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response data:", err)
		return "", err
	}

	return string(respBody), nil
}

func getAccounts() {
	accounts := config.OpenAPI["accounts"]
	response, err2 := NewRequest(accounts.Method, accounts.Url, nil)
	if err2 != nil {
		return
	}
	fmt.Println("response: ", response)
}
