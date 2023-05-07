//package main
//
//import (
//	"math/rand"
//	"time"
//)
//
////选择winner，通过随机选择块来选择验证者来伪造一个区块链，并通过标记的数量加权
//func PickWinner() {
//	time.Sleep(time.Second * 30)
//	mutex.Lock()
//	temp := tempBlocks
//	mutex.Unlock()
//
//	lotteryPool := []string{}
//	if len(temp) > 0 {
//
//	OUTER:  //使用这个循环来判断是否已经存在相同的验证在temp当中
//		for _, block := range temp { //索引值不用，设"_"
//			// if already in lottery pool, skip
//			for _, node := range lotteryPool {
//				if block.Validator == node {
//					continue OUTER
//				}
//			}
//
//			mutex.Lock()
//			setValidators := validators
//			mutex.Unlock()
//
//			k, ok := setValidators[block.Validator]
//			if ok {
//				for i := 1; i < k; i++ {
//					lotteryPool = append(lotteryPool, block.Validator)
//				}
//			}
//		}
//
//		//从池(lotteryPool)中随机选取winner
//		s := rand.NewSource(time.Now().Unix())
//		r := rand.New(s)
//		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]
//
//		//添加winner中的块block 并让所有节点知道
//		for _, block := range temp {
//			if block.Validator == lotteryWinner {
//				mutex.Lock()
//				Blockchain = append(Blockchain, block)
//				mutex.Unlock()
//				for range validators {
//					announcements <- "\n winning validator: " + lotteryWinner + "\n"
//				}
//				break
//			}
//		}
//
//		mutex.Lock()
//		tempBlocks = []Block{}
//		mutex.Unlock()
//	}
//}