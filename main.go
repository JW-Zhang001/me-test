package main

import (
	"fmt"
	"me-test/testsuite"
	"me-test/tools"
)

func main() {
	nodeID := tools.RandNodeID()
	fmt.Println(nodeID)
	regionID, ok := testsuite.TestNewValidatorRegion(nodeID)
	if !ok {
		fmt.Println("TestNewValidatorRegion error")
	}
	fmt.Println("------------case1-------:", regionID)

	kycDel, ok := testsuite.TestKycDelegate("b57407f6-e92b-4a01-aa66-562167e51e6f")
	if !ok {
		fmt.Println("TestKycDelegate error")
	}
	fmt.Println("------------case2-------:", kycDel)

	del, ok := testsuite.TestDelegate()
	if !ok {
		fmt.Println("TestDelegate error")
	}
	fmt.Println("------------case3-------:", del)

	mixDel, ok := testsuite.TestMixDelegate(regionID)
	if !ok {
		fmt.Println("TestMixDelegate error")
	}
	fmt.Println("------------case4-------:", mixDel)

	//goroutineNum := 4
	//wg := sync.WaitGroup{}
	//for i := 0; i < goroutineNum; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < 4; j++ {
	//			ok, regionID := testsuite.TestNewValidatorRegion()
	//			if !ok {
	//				fmt.Println("TestNewValidatorRegion error")
	//			}
	//			fmt.Println(regionID)
	//		}
	//	}()
	//}
	//wg.Wait()
	//fmt.Println("Create validator and region success")
}
