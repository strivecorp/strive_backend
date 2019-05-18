package blockchain

import (
	"crypto/sha256"
	"strconv"
	"encoding/hex"
	"time"
	"os"
	"io/ioutil"
	"encoding/json"
)

func InitiateBlockChain(loc string) (*ChainInfo) {
	info := &ChainInfo{}

	if _, err := os.Stat(loc); os.IsNotExist(err) {
		f, err := os.Create(loc)
		if err != nil {
			return nil
		}
		info.FileVar = f
		info.BlockChain = []Block{generateFirst()}
	} else {
		saved, _ := ioutil.ReadFile(loc)
		err := json.Unmarshal(saved, &info.BlockChain)
		if err != nil {
			return nil
		}
		f, _ := os.Open(loc)
		info.FileVar = f
	}

	return info
}

func generateFirst() Block {
	initial := Block{
		Index:     0,
		Timestamp: time.Now(),
		Event: Event{
			Title:       "Init",
			Description: "Init",
			Deadline:    time.Now().Add(1 * time.Second),
		},
		Hash:         "0",
		PreviousHash: "0",
	}

	initial.Hash = calculateHash(initial)
	return initial
}

func GenerateBlock(oldBlock Block, event Event) Block {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now(),
		Event:        event,
		Hash:         "",
		PreviousHash: oldBlock.Hash,
	}

	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func validateBlock(oldBlock, newBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PreviousHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	h := sha256.New()
	unHashed := strconv.Itoa(block.Index) + block.Timestamp.String() + goalToString(block.Event) + block.PreviousHash + strconv.Itoa(block.Event.TransactionType)
	h.Write([]byte(unHashed))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func goalToString(event Event) string {
	return event.Title + event.Description + event.Deadline.String()
}
