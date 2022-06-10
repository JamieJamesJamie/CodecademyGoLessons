package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func main() {
	guests := [4]string{"Person1", "Person2", "Person3", "Person4"}
	catObjects := [3]string{"Toy Chest", "Crate", "Box"}

	fmt.Println("Guests:", guests)
	fmt.Println("Cat's objects:", catObjects)

	rand.Seed(time.Now().UnixNano())

	randomGuest := getRandomElement(guests[:])
	randomCatObject := getRandomElement(catObjects[:])

	fmt.Println(randomGuest, "hid the cat by putting it in the", randomCatObject)
}
