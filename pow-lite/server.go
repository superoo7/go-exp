package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Message basic data being stored in blockchain
type Message struct {
	Msg string
}

// RunServer run server
func RunServer() error {
	serverHandler := http.HandlerFunc(serverLogic)
	http.Handle("/blockchain", serverHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}

func serverLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get BC
		bytes, err := json.MarshalIndent(Blockchain, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(bytes))
	} else if r.Method == "POST" {
		// Post BC
		if r.Header.Get("Content-type") != "application/json" {
			respondWithJSON(w, r, http.StatusUnsupportedMediaType, "415 - Unsupported Media Type. Please send JSON")
			return
		}
		var m Message
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()
		log.Println(m.Msg)

		//ensure atomicity when creating new block
		lock.Lock()
		newBlock := GenerateBlock(Blockchain[len(Blockchain)-1], m.Msg)
		lock.Unlock()

		if IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			Blockchain = append(Blockchain, newBlock)
		}

		respondWithJSON(w, r, http.StatusCreated, newBlock)
		return
	} else {
		http.Error(w, "Invalid method, use GET or POST", http.StatusInternalServerError)
	}
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
