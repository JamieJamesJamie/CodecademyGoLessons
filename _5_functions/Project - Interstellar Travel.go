package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("You have this much fuel left:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Printf("Welcome to the planet %s!", planet)
	fmt.Println()
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	// Test your functions!
	fmt.Println("Test 'fuelGauge()':")
	fuelGauge(10)
	fuelGauge(0)
	fuelGauge(-1)
	fmt.Println()

	fmt.Println("Test 'calculateFuel()':")
	fmt.Println("Venus:", calculateFuel("Venus"))
	fmt.Println("Mercury:", calculateFuel("Mercury"))
	fmt.Println("Mars:", calculateFuel("Mars"))
	fmt.Println("Jupiter:", calculateFuel("Jupiter"))
	fmt.Println()

	fmt.Println("Test 'greetPlanet()':")
	greetPlanet("Mercury")
	greetPlanet("")
	fmt.Println()

	fmt.Println("Test 'cantFly()':")
	cantFly()
	fmt.Println()

	// Create `planetChoice` and `fuel`
	var fuel int = 1000000
	planetChoice := "Venus"

	// And then liftoff!
	fuelGauge(fuel)
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
