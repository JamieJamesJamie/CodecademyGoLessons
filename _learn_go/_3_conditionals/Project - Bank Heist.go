package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened.")
	}

	leftSafety := rand.Intn(5)

	if isHeistOn {
		switch leftSafety {
		case 0:
			isHeistOn = false
			fmt.Println("The heist failed.")
		case 1:
			isHeistOn = false
			fmt.Println("Fell into a cell...")
		case 2:
			isHeistOn = false
			fmt.Println("Hyperinflation caused the money to be worthless...")
		case 3:
			isHeistOn = false
			fmt.Println("Walked into a black hole...")
		default:
			fmt.Println("Start the getaway car!")
		}

		if isHeistOn {
			amtStolen := 10000 + rand.Intn(1000000)

			fmt.Println("Successfully stolen:", amtStolen)
		}
	}

	fmt.Println("isHeistOn:", isHeistOn)
}
