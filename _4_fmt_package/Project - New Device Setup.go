package main

import "fmt"

func main() {
	var name string

	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello", name, "!")

	var age int

	fmt.Scan(&age)
	fmt.Printf("Hello %s! You are %d years old.\n", name, age)

	newMessage := fmt.Sprintf("Wow %s, you're %d years old!", name, age)
	fmt.Println(newMessage)
}
