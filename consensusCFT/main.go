package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var port string
var dataC chan string
var LPort string
var port0 string
var port1 string
var port2 string
var port3 string
var port4 string

func main() {
	port = os.Args[1]
	fmt.Println("run at ", port)
	LPort = os.Args[2]
	dataC = make(chan string, 10)
	port0 = "1000"
	port1 = "1001"
	port2 = "1002"
	port3 = "1003"
	port4 = "1004"

	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/consensus", consensusHandler)

	go http.ListenAndServe("0.0.0.0:"+port, nil)
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
	if port == LPort {
		if port != port0 {
			sendToOthers(data, port0)
		}
		if port != port1 {
			sendToOthers(data, port1)
		}
		if port != port2 {
			sendToOthers(data, port2)
		}
		if port != port3 {
			sendToOthers(data, port3)
		}
		if port != port4 {
			sendToOthers(data, port4)
		}
		dataC <- data
	} else {
		fwToLeader(LPort, data)
	}
}

func fwToLeader(LPort, data string) {
	http.Get("http://127.0.0.1:" + LPort + "/data?data=" + data)
}

func sendToOthers(input, port string) {
	data := &ConsensusData{
		Data: input,
	}
	d, _ := json.Marshal(data)
	body := strings.NewReader(string(d))
	http.Post("http://127.0.0.1:"+port+"/consensus", "application/x-www-form-urlencoded", body)
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
