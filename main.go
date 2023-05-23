package main

import (
	"fmt"
	"go.uber.org/zap"
	"me-test/client/query"
	"me-test/testsuite"
)

func work1() {
	nodeID, err := query.GetChainNotExistNodeID()
	if err != nil {
		zap.S().Info("GetChainNotExistNodeID error", err)
	}

	regionID, ok := testsuite.TestNewValidatorRegion(nodeID)
	if !ok {
		fmt.Println("TestNewValidatorRegion error")
		return
	}
	fmt.Println("----------case1/PASS----------:", regionID)
}

func main() {
	work1()

	//kycDel, ok := testsuite.TestKycDelegate("73afbfa2-d9b2-41c8-92f7-da0037185338")
	//if !ok {
	//	fmt.Println("TestKycDelegate error")
	//}
	//fmt.Println("------------case2-------:", kycDel)

	//del, ok := testsuite.TestDelegate()
	//if !ok {
	//	fmt.Println("TestDelegate error")
	//}
	//fmt.Println("------------case3-------:", del)
	//
	//mixDel, ok := testsuite.TestMixDelegate(regionID)
	//if !ok {
	//	fmt.Println("TestMixDelegate error")
	//}
	//fmt.Println("------------case4-------:", mixDel)

	//kycUnDel, ok := testsuite.TestKycUnDelegate(regionID)
	//if !ok {
	//	fmt.Println("TestKycUnDelegate error")
	//	return
	//}
	//fmt.Println("------------case5-------:", kycUnDel)

	//kycUnDel, ok := testsuite.TestUnDelegate()
	//if !ok {
	//	fmt.Println("TestUnDelegate error")
	//	return
	//}
	//fmt.Println("------------case6-------:", kycUnDel)

	// mixUnDel, ok := testsuite.TestMixUnDelegate("73afbfa2-d9b2-41c8-92f7-da0037185338")
	// if !ok {
	// 	fmt.Println("TestMixDelegate error")
	// 	return
	// }
	// fmt.Println("------------case7-------:", mixUnDel)

	// depAccAddr, ok := testsuite.TestKycDeposit("73afbfa2-d9b2-41c8-92f7-da0037185338")
	// if !ok {
	// 	fmt.Println("TestNotKycDeposit error")
	// 	return
	// }
	// fmt.Println("------------case8-------:", depAccAddr)

	// depAccAddr, ok := testsuite.TestExpireDepositWithdraw("ba51de1d-c17a-4e43-bc19-0da82083bf9c")
	// if !ok {
	// 	fmt.Println("TestExpireDepositWithdraw error")
	// 	return
	// }
	// fmt.Println("------------case9-------:", depAccAddr)

	// depAccAddr, ok := testsuite.TestNotExpireDepositWithdraw("ba51de1d-c17a-4e43-bc19-0da82083bf9c")
	// if !ok {
	// 	fmt.Println("TestNotExpireDepositWithdraw error")
	// 	return
	// }
	// fmt.Println("------------case10-------:", depAccAddr)

	//notKycFees, ok := testsuite.TestNotKycFees()
	//if !ok {
	//	fmt.Println("TestNotKycFees error")
	//	return
	//}
	//
	//fmt.Println("------------case11-------:", notKycFees)

	//kycFeesIsPm, ok := testsuite.TestKycFeesValidatorOwnerIsPM("beb8095f-02b9-40cd-9727-7f5bfa32f119")
	//if !ok {
	//	fmt.Println("TestKycFeesValidatorIsPM error")
	//	return
	//}
	//fmt.Println("------------case12-------:", kycFeesIsPm)

	//kycFeesIsUser, ok := testsuite.TestKycFeesValidatorOwnerIsUser("beb8095f-02b9-40cd-9727-7f5bfa32f119")
	//if !ok {
	//	fmt.Println("TestKycFeesValidatorOwnerIsUser error")
	//	return
	//}
	//fmt.Println("------------case13-------:", kycFeesIsUser)

	//acc, err := tools.GenWalletAcc()
	//if err != nil {
	//	return
	//}
	//fmt.Println("accAddr info: ", acc)
	//validatorID, ok := testsuite.TestEditValidator(acc["Addr"], "node6")
	//if !ok {
	//	fmt.Println("TestEditValidator error")
	//	return
	//}
	//fmt.Println("------------case14-------:", validatorID)

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
