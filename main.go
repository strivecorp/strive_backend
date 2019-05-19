package main

import (
	"net/http"
	"strive_backend/blockchain"
	"time"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"strive_backend/API"
)

type Init interface {
	makeBlock(w http.ResponseWriter, r *http.Request)
	readBlockChain(w http.ResponseWriter, r *http.Request)
}

type i struct {
	bc *blockchain.ChainInfo
}

func main() {
	API.APIHandler()
	var info i
	info.bc = blockchain.InitiateBlockChain("./data.json")
	http.HandleFunc("/makeBlock", info.makeBlock)
	http.HandleFunc("/readBlockchain", info.readBlockchain)

	http.ListenAndServe(":8080", nil)
}

func startBC(w http.ResponseWriter, r *http.Request) {

}

func (i i) makeBlock(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query()["title"][0]
	description := r.URL.Query()["description"][0]
	eventType, _ := strconv.Atoi(r.URL.Query()["event"][0])
	i.bc.BlockChain = append(
		i.bc.BlockChain,
		blockchain.GenerateBlock(
			i.bc.BlockChain[len(i.bc.BlockChain)-1],
			blockchain.Event{title, description, time.Now(), eventType},
		),
	)


	js, _ := json.Marshal(i.bc.BlockChain)
	_ = ioutil.WriteFile("./data.json", js, 0644)
}

func (i i) readBlockchain(w http.ResponseWriter, r *http.Request) {
	js, _ := json.Marshal(i.bc.BlockChain)
	w.Write(js)
}
