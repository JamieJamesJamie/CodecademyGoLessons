package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

// Task implementation goes here

type Stock struct {
	name  string
	price float32
}

func (stock *Stock) updateStock() {
	change := randomNumberGen(-10000, 10000)
	stock.price += change
}

func displayMarket(market []Stock) {
	for _, stock := range market {
		fmt.Println(stock)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Function calls go here
}
