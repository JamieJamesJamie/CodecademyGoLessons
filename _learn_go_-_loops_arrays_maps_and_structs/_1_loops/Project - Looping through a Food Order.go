package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

// WRITE CONTAINS FUNCTION BELOW

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string
	// WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW

	// WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW

	fmt.Println("The total for the order is", total)
}
