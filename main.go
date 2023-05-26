package main

import (
	"fmt"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/client/query"
	"me-test/testsuite"
)

func main() {
	smokeTest := []func(){case1, case2, case3, case4, case5, case6, case7, case8, case9, case10, case11, case12, case13, case14}
	for _, f := range smokeTest {
		f()
	}

	fmt.Println("All use cases have executed")
}

func test() {
	//ch := make(chan func(), len(caseList))
	//// 将函数添加到通道中
	//for _, c := range caseList {
	//	ch <- c
	//}
	//
	//goroutineNum := 10
	//wg := sync.WaitGroup{}
	//for i := 0; i < goroutineNum; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for c := range ch {
	//			c()
	//		}
	//	}()
	//}
	//
	//close(ch)
	//wg.Wait()
}

func case1() {
	nodeID, err := query.GetChainNotExistNodeID()
	if err != nil {
		zap.S().Error("GetChainNotExistNodeID error: ", err)
		return
	}

	regionID, ok := testsuite.TestNewValidatorRegion(nodeID)
	if !ok {
		zap.S().Error("TestNewValidatorRegion error")
		return
	}
	zap.S().Info("----------case1/PASS----------:", regionID)
}

func case2() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	kycDel, ok := testsuite.TestKycDelegate(regionID)
	if !ok {
		zap.S().Error("TestKycDelegate error")
		return
	}
	zap.S().Info("----------case2/PASS----------:", kycDel)
}

func case3() {
	del, ok := testsuite.TestDelegate()
	if !ok {
		zap.S().Error("TestDelegate error")
		return
	}
	zap.S().Info("----------case3/PASS----------:", del)
}

func case4() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	mixDel, ok := testsuite.TestMixDelegate(regionID)
	if !ok {
		zap.S().Error("TestMixDelegate error")
		return
	}
	zap.S().Info("----------case4/PASS----------:", mixDel)
}

func case5() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	kycUnDel, ok := testsuite.TestKycUnDelegate(regionID)
	if !ok {
		zap.S().Error("TestKycUnDelegate error")
		return
	}
	zap.S().Info("----------case5/PASS----------:", kycUnDel)
}

func case6() {
	kycUnDel, ok := testsuite.TestUnDelegate()
	if !ok {
		zap.S().Error("TestUnDelegate error")
		return
	}
	zap.S().Info("----------case6/PASS----------:", kycUnDel)
}

func case7() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	mixUnDel, ok := testsuite.TestMixUnDelegate(regionID)
	if !ok {
		zap.S().Error("TestMixDelegate error")
		return
	}
	zap.S().Info("----------case7/PASS----------:", mixUnDel)
}

func case8() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	depAccAddr, ok := testsuite.TestKycDeposit(regionID)
	if !ok {
		zap.S().Error("TestNotKycDeposit error")
		return
	}
	zap.S().Info("----------case8/PASS----------:", depAccAddr)
}

func case9() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	depAccAddr, ok := testsuite.TestExpireDepositWithdraw(regionID)
	if !ok {
		zap.S().Error("TestExpireDepositWithdraw error")
		return
	}
	zap.S().Info("----------case9/PASS----------:", depAccAddr)
}

func case10() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	depAccAddr, ok := testsuite.TestNotExpireDepositWithdraw(regionID)
	if !ok {
		zap.S().Error("TestNotExpireDepositWithdraw error")
		return
	}
	zap.S().Info("----------case10/PASS----------:", depAccAddr)
}

func case11() {
	notKycFees, ok := testsuite.TestNotKycFees()
	if !ok {
		zap.S().Error("TestNotKycFees error")
		return
	}
	zap.S().Info("----------case11/PASS----------:", notKycFees)
}

func case12() {
	regionID, err := query.GetChainExistRegionID()
	if err != nil {
		zap.S().Error("GetChainExistRegionID error", err)
		return
	}
	kycFeesIsPm, ok := testsuite.TestKycFeesValidatorOwnerIsPM(regionID)
	if !ok {
		zap.S().Error("TestKycFeesValidatorIsPM error")
		return
	}
	zap.S().Info("----------case12/PASS----------:", kycFeesIsPm)
}

func case13() {
	nodeID, err := query.GetChainNotExistNodeID()
	if err != nil {
		zap.S().Error("GetChainNotExistNodeID error: ", err)
		return
	}
	kycFeesIsUser, ok := testsuite.TestKycFeesValidatorOwnerIsUser(nodeID)
	if !ok {
		zap.S().Error("TestKycFeesValidatorOwnerIsUser error")
		return
	}
	zap.S().Info("----------case13/PASS----------:", kycFeesIsUser)
}

func case14() {
	acc, err := client.GenWalletAcc()
	if err != nil {
		return
	}
	zap.S().Info("userAccAddr info: ", acc)

	nodeID, err := query.GetChainNotExistNodeID()
	if err != nil {
		zap.S().Error("GetChainNotExistNodeID error: ", err)
		return
	}

	validatorID, ok := testsuite.TestEditValidator(acc["Addr"], nodeID)
	if !ok {
		zap.S().Error("TestEditValidator error")
		return
	}
	zap.S().Info("----------case14/PASS----------:", validatorID)
}
