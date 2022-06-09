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
func contains(menu []string, order string) bool {
	for _, menuItem := range menu {
		if order == menuItem {
			return true
		}
	}
	return false
}

func fastFoodMenuContains(fastFoodMenu []string, order string) int {
	if contains(fastFoodMenu, order) {
		return 4
	} else if order == "quit" {
		fmt.Println("Quitting...")
	} else {
		fmt.Println(order, "is not on the menu.")
	}
	return 0
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string
	// WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW

	for order != "quit" {
		order = askOrder()
		total += fastFoodMenuContains(fastfoodMenu, order)
	}

	var numBillsToPay uint

	for amount := total; amount > 0; amount -= 2 {
		numBillsToPay++
	}

	fmt.Println("Number of £2 notes to pay:", numBillsToPay)

	// WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("Amount in GBP(£) to be paid to cashier:", amount)
		fmt.Println("Giving £2 to cashier...")
	}

	fmt.Println("The total for the order is", total)
}
