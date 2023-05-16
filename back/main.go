package main

import (
	simpleBlockChain "awesomeProject/simple-block-chain"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io"
	"log"
	"net/http"
)

var bc *simpleBlockChain.Blockchain

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.WriteString(w, string(bytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	bc.AddBlock(data)
	handleGetBlockchain(w, r)
}

func main() {
	bc = simpleBlockChain.NewBlockchain()
	mux := http.NewServeMux()
	mux.HandleFunc("/get", handleGetBlockchain)
	mux.HandleFunc("/write", handleWriteBlock)
	fmt.Println("Listening on http://localhost:8080")
	c := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", c)
	if err != nil {
		log.Fatal(err)
	}
}
