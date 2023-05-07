//package main
//
//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"strconv"
//	"time"
//)
//
//func HandleConn(conn net.Conn)  {
//	defer conn.Close()
//
//	go func() {
//		for  {
//			//These announcements will be who the winning validator is when one is chosen
//			//被选中的winner进行传入到 announcements中
//			msg := <-announcements
//			io.WriteString(conn, msg)
//		}
//	}()
//	// validator address
//	var address string
//	// 允许验证者输入他想要加入的tokens的数量
//	// 拥有足够多的tokens，就更有机会获得新的块
//	io.WriteString(conn, "Enter token balance:")// 使用natcat进行输入值
//	scanBalance := bufio.NewScanner(conn)
//	for scanBalance.Scan() {
//		balance, err := strconv.Atoi(scanBalance.Text())
//		if err != nil {
//			log.Panicf("%v not a number: %v",scanBalance.Text(),err)
//			return
//		}
//
//		t:= time.Now()
//		address = CalculateHash(t.String())
//		validators[address] = balance
//		fmt.Println(validators)
//		break
//	}
//
//	io.WriteString(conn, "\nEnter a new BPM:")
//
//	scanBPM := bufio.NewScanner(conn)
//
//	go func() {
//		for {
//			//进行输入的BPM验证
//			for scanBPM.Scan() {
//				bpm, err := strconv.Atoi(scanBPM.Text())
//				//如果恶意方试图用错误的输入来改变链,则将此map删除
//				//在这使用了一个简单的逻辑，就是判断输入的BMP是否为一个整数格式
//				if err != nil {
//					log.Printf("%v not a number: %v", scanBPM.Text(), err)
//					delete(validators, address)
//					conn.Close()
//				}
//
//				mutex.Lock()
//				oldLastIndex := Blockchain[len(Blockchain)-1]
//				mutex.Unlock()
//
//				//创建新块block，并考虑起是否伪造
//				newBlock, err := GenerateBlock(oldLastIndex, bpm, address)
//				if err != nil {
//					log.Println(err)
//					continue //输出所有log err
//				}
//				if IsBlockValid(newBlock, oldLastIndex) {
//					candidateBlocks <- newBlock
//				}
//				io.WriteString(conn, "\nEnter a new BPM:")
//			}
//		}
//	}()
//
//	//模拟接收广播
//	for {
//		time.Sleep(time.Minute)
//		mutex.Lock()
//		//用一个规整的json格式输出区块
//		output, err := json.MarshalIndent(Blockchain, "", "\t")
//		mutex.Unlock()
//		if err != nil {
//			log.Fatal(err)
//		}
//		io.WriteString(conn, string(output)+"\n")
//	}
//
//}