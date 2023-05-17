package main

import (
	"fmt"
	"me-test/testsuite"
)

func main() {
	//res := validator.GetTmPubKey()
	//fmt.Println(res)
	//res, _ := initialize.GetValidatorPubKey("node3")
	//fmt.Println(res)
	//data.Test1()
	//res, _ := testsuite.NewValidatorRegion()
	//fmt.Println(res)

	ok := testsuite.TestKycDelegate()
	if ok {
		fmt.Println("pass")
	}

}
