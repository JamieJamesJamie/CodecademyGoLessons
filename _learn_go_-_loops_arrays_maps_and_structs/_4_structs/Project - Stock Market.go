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

	stockMarket := []Stock{
		{name: "GOOG", price: 2313.50},
		{name: "AAPL", price: 157.28},
		{name: "FB", price: 203.77},
		{name: "TWTR", price: 50.00},
	}

	fmt.Println("Current state of the stock market:")
	displayMarket(stockMarket)
	fmt.Println()

	stockMarket[1].updateStock()

	fmt.Println(fmt.Sprintf("Updated state of the stock market (the price of %s has been updated):", stockMarket[1].name))
	displayMarket(stockMarket)
}
