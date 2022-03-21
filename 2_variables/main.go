package main

import "fmt"

func main() {
	// Static type
	var num1, num2 int = 10, 20
	var message string
	message = "Enter your message here"
	fmt.Println(message)

	// Dynamic type
	name := "Arslan Haider Sherazi"
	fmt.Println(name, num1, num2)
}
