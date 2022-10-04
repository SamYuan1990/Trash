package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var address string
var dataC chan string

func main() {
	address = os.Args[1]
	fmt.Println("run at ", address)
	dataC = make(chan string, 10)

	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/consensus", consensusHandler)

	go http.ListenAndServe(address, nil)
	for {
		d := <-dataC
		fmt.Println(d)
	}
}

func dataHandler(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	data := vars["data"][0]
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`data been received ` + data))
	if err != nil {
		fmt.Errorf("%s", fmt.Sprintf("failed to write response: %v", err))
	}
	dataC <- data
}

type ConsensusData struct {
	Data string `json: "Data"`
}

func consensusHandler(w http.ResponseWriter, req *http.Request) {
	var d ConsensusData
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	dataC <- d.Data
}
