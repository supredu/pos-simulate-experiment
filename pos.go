package main

import (
	"crypto/sha256"
	"math/big"
)

// 函数名：Pos,传入Miners数组，当前难度值Dif和一个string类型变量tradeData，内设一个int变量timeCounter, 从0递增到Intmax，
//hash值为SHA256(SHA256(tradeData|timeCounter)),循环内遍历Miners数组，目标值target=Dif乘当前Miner的币龄，
//要求hash小于target，返回满足要求的第一个Miner的序号并清空这个Miner的币龄，一旦满足要求则退出整个循环
func Pos(Miners *[]Miner, Dif int64, tradeData string) int {
	var timeCounter int
	target := big.NewInt(1)

	for timeCounter = 0; timeCounter < INT64_MAX; timeCounter++ {
		hash := sha256.Sum256([]byte(tradeData + string(timeCounter)))
		hash = sha256.Sum256(hash[:])
		var hashInt big.Int
		hashInt.SetBytes(hash[:])
		for i := 0; i < len(*Miners); i++ {
			//最小持币量为3才能挖矿
			if (*Miners)[i].num <= 2 {
				continue
			}
			// 数据长度为8位
			//需求：需要满足前两位为0，才能解决问题
			//1 * 2 << (8-2) = 64
			// 0100 0000
			// 00xx xxxx
			// 32 * 8
			target.Lsh(target, uint(Dif)*uint((*Miners)[i].coinAge))
			if hashInt.Cmp(target) < 0 {
				(*Miners)[i].coinAge = 0
				return i
			}
		}
	}
	return -1
}
