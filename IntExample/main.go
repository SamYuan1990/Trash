package main

import (
	"math/big"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func randomInt() int {
	return seededRand.Intn(100-1) + 1
}

func Int() int {
	return randomInt() + randomInt()
}

func BigInt() int {
	a := new(big.Int).SetInt64(int64(randomInt()))
	b := new(big.Int).SetInt64(int64(randomInt()))
	sum := a.Add(a, b)
	return int(sum.Int64())
}

func main() {

}
