package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var port string
var dataC chan string
var port0 string
var port1 string
var port2 string
var port3 string
var port4 string
var groups []string
var d_commitMap, d_prepareMap map[string]int
var lock sync.Mutex

func main() {
	port = os.Args[1]
	fmt.Println("run at ", port)
	dataC = make(chan string, 10)
	port0 = "10000"
	port1 = "10001"
	port2 = "10002"
	port3 = "10003"
	port4 = "10004"
	groups = make([]string, 5)
	groups[0] = port0
	groups[1] = port1
	groups[2] = port2
	groups[3] = port3
	groups[4] = port4
	d_prepareMap = make(map[string]int)
	d_commitMap = make(map[string]int)

	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/preprepare", preprepareHandler)
	http.HandleFunc("/prepare", prepareHandler)
	http.HandleFunc("/commit", commitHandler)

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
	d_commitMap[data] = 0
	for _, v := range groups {
		if v != port {
			preprepareToOthers(data, v)
		}
	}
}

func preprepareToOthers(input, port string) {
	data := &ConsensusData{
		Data: input,
	}
	d, _ := json.Marshal(data)
	body := strings.NewReader(string(d))
	http.Post("http://127.0.0.1:"+port+"/preprepare", "application/x-www-form-urlencoded", body)
}

type ConsensusData struct {
	Data string `json: "Data"`
}

func preprepareHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("start in preprepare phase")
	var d ConsensusData
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	d_prepareMap[d.Data] = 0
	// should have a sign for d.data skip for now
	for _, v := range groups {
		if v != port {
			fmt.Println("send prepare to " + v + " with data " + d.Data)
			prepareToOthers(d.Data, v)
		}
	}
}

func prepareHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("start in prepare phase")
	var d ConsensusData
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// should have a verify for d.data skip for now
	// 2f+1 = 5 f=2 factor=3
	lock.Lock()
	d_prepareMap[d.Data] += 1
	fmt.Println("d_prepareMap", d_prepareMap[d.Data])
	lock.Unlock()
	if d_prepareMap[d.Data] == 2 {
		for _, v := range groups {
			if v != port {
				fmt.Println("send commit to " + v + " with data " + d.Data)
				commitToOthers(d.Data, v)
			}
		}
	}
}

func prepareToOthers(input, port string) {
	data := &ConsensusData{
		Data: input,
	}
	d, _ := json.Marshal(data)
	body := strings.NewReader(string(d))
	http.Post("http://127.0.0.1:"+port+"/prepare", "application/x-www-form-urlencoded", body)
}

func commitToOthers(input, port string) {
	data := &ConsensusData{
		Data: input,
	}
	d, _ := json.Marshal(data)
	body := strings.NewReader(string(d))
	http.Post("http://127.0.0.1:"+port+"/commit", "application/x-www-form-urlencoded", body)
}

func commitHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("start in commit phase")
	var d ConsensusData
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// 2f+1 = 5 f=2 factor=3
	lock.Lock()
	d_commitMap[d.Data] += 1
	fmt.Println("d_commitMap", d_commitMap[d.Data])
	lock.Unlock()
	if d_commitMap[d.Data] == 2 {
		dataC <- d.Data
	}
}
