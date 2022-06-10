package main

import fmt "fmt"

func main() {
	currencies := map[string]float32{"JPY": 130.2, "EUR": 0.95}

	var dollarAmount float32
	var currency string

	fmt.Println("How many dollars (USD) would you like to convert?")
	fmt.Print("$")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Invalid stick price.")
	} else {
		fmt.Println("What is the target currency?")
		fmt.Scan(&currency)

		rate, isValid := currencies[currency]

		if !isValid {
			fmt.Printf("\"%s\" is not a valid currency.", currency)
			fmt.Println()
			return
		}

		fmt.Printf("%.2f (USD) has been converted to %.2f (%s)", dollarAmount, dollarAmount*rate, currency)
		fmt.Println()
	}
}
