//package main
//
//import (
//	"crypto/sha256"
//	"encoding/hex"
//	"strconv"
//	"time"
//)
//
//
//
////哈希算法方法定义
//func CalculateHash(s string) string {
//	h := sha256.New()
//	h.Write([]byte(s))
//	hashed := h.Sum(nil)
//	return hex.EncodeToString(hashed)
//}
//
////计算块哈希
//func CalculateBlockHash(block Block) string{
//	record := strconv.Itoa(block.Index)  +
//		block.Timestamp +
//		strconv.Itoa(block.BPM) +
//		block.PrevHash
//	return CalculateHash(record)
//}
//
////创建新块block
//func GenerateBlock(oldBlock Block, BPM int, address string) (Block, error) {
//	var newBlock Block
//	t := time.Now()
//	newBlock.Index = oldBlock.Index + 1
//	newBlock.Timestamp = t.String()
//	newBlock.BPM =BPM
//	newBlock.PrevHash = oldBlock.Hash
//	newBlock.Hash = CalculateBlockHash(newBlock)
//	newBlock.Validator = address
//	return newBlock, nil
//}
//
////块校验
//func IsBlockValid(newBlock Block, oldBlock Block) bool {
//	if oldBlock.Index +1 != newBlock.Index{
//		return false
//	}
//	if oldBlock.Hash != newBlock.PrevHash {
//		return false
//	}
//	if newBlock.Hash != CalculateBlockHash(newBlock) {
//		return false
//	}
//	return true
//}