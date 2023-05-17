package main

import (
	"fmt"
	"me-test/testcase"
)

func main() {
	//res := validator.GetTmPubKey()
	//fmt.Println(res)
	//res, _ := initialize.GetValidatorPubKey("node3")
	//fmt.Println(res)
	//data.Test1()
	//testcase.NewValidatorRegion()
	userPk, _ := testcase.NewKyc()
	fmt.Println(userPk)
}
