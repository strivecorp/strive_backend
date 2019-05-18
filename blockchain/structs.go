package blockchain

import (
	"time"
	"os"
)

type Block struct {
	Index int
	Timestamp time.Time
	Event Event           // ko bo api končan, daj sem not dejanski event iz API dela
	Hash string
	PreviousHash string
}

type Event struct {
	Title string
	Description string
	Deadline time.Time
	TransactionType int
}

type ChainInfo struct {
	BlockChain []Block
	FileVar *os.File
}
