package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// 定义常量设置最大最小概率，这里由于不可能等30天之后再计算币龄，我设置10秒之后开始计算，并且防止数据过大，按分钟计时
const (
	dif         = 2
	INT64_MAX   = math.MaxInt64
	MaxProbably = 255
	MinProbably = 235
	MaxCoinAge  = 10
	Minute      = 60
)

// 定义币的数据结构和一个币池，这里币中提出了地址Address概念，一般理解为钱包地址，一般一个钱包属于一个用户，表明这个币的所有权是这个地址对应的用户
type Coin struct {
	Time    int64
	Num     int
	Address string
}

var CoinPool []Coin

// 定义区块和链
type Block struct {
	PrevHash  []byte
	Hash      []byte
	Data      string
	Height    int64
	Timestamp int64
	Coin      Coin
	Nonce     int
	Dif       int64
}

type BlockChain struct {
	Blocks []Block
}

// init函数，设置随机数种子和币池初始化，会在main函数开始前自动执行
func init() {
	rand.Seed(time.Now().UnixNano())
	CoinPool = make([]Coin, 0)
}

// 生成创世块，传入的参数是区块上的数据data和挖矿地址addr，这里每个区块的币随机给出1到5，币池增加创世块的币
func GenesisBlock(data string, addr string) *BlockChain {
	var bc BlockChain
	bc.Blocks = make([]Block, 1)
	newCoin := Coin{
		Time:    time.Now().Unix(),
		Num:     1 + rand.Intn(5),
		Address: addr,
	}
	bc.Blocks[0] = Block{
		PrevHash:  []byte(""),
		Data:      data,
		Height:    1,
		Timestamp: time.Now().Unix(),
		Coin:      newCoin,
		Nonce:     0,
	}
	bc.Blocks[0].Hash, bc.Blocks[0].Nonce, bc.Blocks[0].Dif = ProofOfStake(dif, addr, bc.Blocks[0])
	CoinPool = append(CoinPool, newCoin)
	return &bc
}

// 生成新区块函数，还是一样，prevHash对应上一个节点的Hash，随机给出1-5币奖励，记录该币的所有者地址addr，币池增加新区块的币
func GenerateBlock(bc *BlockChain, data string, addr string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newCoin := Coin{
		Time:    time.Now().Unix(),
		Num:     1 + rand.Intn(5),
		Address: addr,
	}
	b := Block{
		PrevHash:  prevBlock.Hash,
		Data:      data,
		Height:    prevBlock.Height + 1,
		Timestamp: time.Now().Unix(),
	}
	b.Hash, b.Nonce, b.Dif = ProofOfStake(dif, addr, b)
	b.Coin = newCoin
	bc.Blocks = append(bc.Blocks, b)
	CoinPool = append(CoinPool, newCoin)
}
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

// SHA256(SHA256(tradeData|timeCounter))<= D x coinAge
// POS机制的核心算法，设置的是10s之后开始计算币龄。
func ProofOfStake(dif int, addr string, b Block) ([]byte, int, int64) {
	var coinAge int64
	var realDif int64
	realDif = int64(MinProbably)
	curTime := time.Now().Unix()

	for k, i := range CoinPool {
		if i.Address == addr && i.Time+MaxCoinAge < curTime {
			//币龄增加, 并设置上限
			var curCoinAge int64
			if curTime-i.Time < 3*MaxCoinAge {
				curCoinAge = curTime - i.Time
			} else {
				curCoinAge = 3 * MaxCoinAge
			}
			coinAge += int64(i.Num) * curCoinAge
			//参与挖矿的币龄置为0
			CoinPool[k].Time = curTime
		}
	}

	if realDif+int64(dif)*coinAge/Minute > int64(MaxProbably) {
		realDif = MaxProbably
	} else {
		realDif += int64(dif) * coinAge / Minute
	}

	target := big.NewInt(1)
	target.Lsh(target, uint(realDif))
	timeCounter := 0
	for ; timeCounter < INT64_MAX; timeCounter++ {
		check := bytes.Join(
			[][]byte{
				b.PrevHash,
				[]byte(b.Data),
				IntToHex(b.Height),
				IntToHex(b.Timestamp),
				IntToHex(int64(timeCounter)),
			},
			[]byte{})
		hash := sha256.Sum256(check)
		var hashInt big.Int
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			return hash[:], timeCounter, 255 - realDif
		}
	}

	return []byte(""), -1, 255 - realDif
}

// 打印区块和打印币池
func Print(bc *BlockChain) {
	for _, i := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", i.PrevHash)
		fmt.Printf("Hash: %x\n", i.Hash)
		fmt.Println("Block's Data: ", i.Data)
		fmt.Println("Current Height: ", i.Height)
		fmt.Println("Timestamp: ", i.Timestamp)
		fmt.Println("Nonce: ", i.Nonce)
		fmt.Println("Dif: ", i.Dif)
	}
}

func PrintCoinPool() {
	for _, i := range CoinPool {
		fmt.Println("Coin's Num: ", i.Num)
		fmt.Println("Coin's Time: ", i.Time)
		fmt.Println("Coin's Owner: ", i.Address)
	}
}

// main函数测试，虚拟两个地址，当然真实区块链的地址不可能这么简单，需要使用公钥生成地址算法。
// 给addr1先记账，获取一定币，等待可以计算币龄时再次记账，再给新地址addr2记账，看看难度对比
func main() {
	addr1 := "192.168.1.1"
	addr2 := "192.168.1.2"
	bc := GenesisBlock("reigns", addr1)
	GenerateBlock(bc, "send 1$ to alice", addr1)
	GenerateBlock(bc, "send 1$ to bob", addr1)
	GenerateBlock(bc, "send 2$ to alice", addr1)
	time.Sleep(11 * time.Second)
	GenerateBlock(bc, "send 3$ to alice", addr1)
	GenerateBlock(bc, "send 4$ to alice", addr2)
	Print(bc)
	PrintCoinPool()
}
