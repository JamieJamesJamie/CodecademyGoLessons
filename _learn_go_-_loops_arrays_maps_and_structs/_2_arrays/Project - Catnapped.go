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
	catObjects := [3]string{"A Toy Chest", "A Crate", "A Box"}

	fmt.Println("Guests:", guests)
	fmt.Println("Cat's objects:", catObjects)

	rand.Seed(time.Now().UnixNano())
}
