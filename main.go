package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

//创建币池数组Coins
var Coins []Coin
//调用InitBlockChain函数，生成一个区块数组
var BlockChain []Block
//默认难度值dif为1
var Dif int64 = 1
//创建矿工数组Miners
var Miners []Miner


func main() {
	//默认难度值dif为1
	//var Dif int64 = 1
	//创建矿工数组Miners
	//var Miners []Miner
	Miners = InitMiners()
	//添加矿工
	AddMiners()

	//创建币池数组Coins
	//var Coins []Coin
	//给矿工数组中的矿工添加币
	Coins = InitCoins(Miners)
	for i := 0; i < len(Miners); i++ {
		AddCoin(NewCoin(int64(i), Miners), &Coins)
	}
	//调用InitBlockChain函数，生成一个区块数组
	//var BlockChain []Block
	BlockChain = InitBlockChain(Miners, Coins)
	//fmt.Println("创建第二个区块")
	//GenerateBlock(0, Miners, Coins[0], "second block", &BlockChain)
	//fmt.Println("创建结束")
	//时间延迟，给出币龄
	time.Sleep(5*time.Second)
	UpdateMiners(&Coins, &Miners)
	PrintMiners(Miners)

	//挖矿
	IsContinueMining()

	//打印区块
	PrintBlockChain()

}

//传入Miners数组，打印矿工数组每个矿工信息的函数

func PrintMiners(Miners []Miner) {
	for i := 0; i <= len(Miners)-1; i++ {
		fmt.Println("Miner", i, ":", hex.EncodeToString(Miners[i].addr), Miners[i].num, Miners[i].coinAge)
	}
}
//打印区块
func PrintBlockChain()  {
	for i, block := range BlockChain {
		prevBlockHash := hex.EncodeToString(block.PrevHash)
		currentHash := hex.EncodeToString(block.Hash)
		if i == 0{
			fmt.Printf("prevBlockHash: %s, currentHash : 0x%s \n",prevBlockHash,currentHash)
		} else{
			fmt.Printf("prevBlockHash: 0x%s, currentHash : 0x%s \n",prevBlockHash,currentHash)
		}
	}
}

//挖矿

func Mine(Miners []Miner,Dif int64, tradeData string,BlockChain *[]Block)  {
	fmt.Println("开始挖矿")
	winnerIndex := CorrectMiner(&Miners , Dif, tradeData)
	if winnerIndex == -1 {
		panic("挖矿失败")
	}
	fmt.Println("挖矿成功")
	AddCoin(NewCoin(int64(winnerIndex), Miners), &Coins)
	GenerateBlock(winnerIndex, Miners, Coins[len(Coins)-1], tradeData, BlockChain)
	time.Sleep(5*time.Second)
	UpdateMiners(&Coins, &Miners)
	PrintMiners(Miners)
}

func IsContinueMining() {
	var isContinue string
	for  {
		fmt.Println("开始or结束 挖矿?y or n")
		fmt.Scanf("%s",&isContinue)
		if isContinue == "y" {
			Mine(Miners, Dif, "New block",&BlockChain)
			continue
		}else if isContinue == "n" {
			break
		}else{
			fmt.Println("输入错误")
			continue
		}
	}
}
