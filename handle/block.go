package handle

import (
	"fmt"
	"me-test/api"
)

func GetBlockHeight() {
	blockReqData := api.Api["blockHeight"]
	resp, err := NewRequest(blockReqData.Method, blockReqData.Url, blockReqData.Body)
	if err != nil {
		fmt.Println("Error response data:", err)
		return
	}
	fmt.Println(resp)
}
