package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func onetime() int {
	return rand.Intn(2)
}

//GOOS=windows GOARCH=amd64 go build .
func main() {
	var turns int
	fmt.Println("please input times for experiment")
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence[0 : len(sentence)-1]))
	}
	turns, err = strconv.Atoi(string(sentence[0 : len(sentence)-1]))
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().UnixNano())
	ones := 0
	for i := 0; i < turns; i++ {
		if onetime() == 1 {
			ones++
		}
	}
	fmt.Printf("for %d times, 1 appears %d times, 0 as %d times \n", turns, ones, (turns - ones))
}
