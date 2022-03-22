package main

import "fmt"

func main() {
	// Static type
	var num1, num2 int = 10, 20
	fmt.Println(num1, num2)

	var message string
	message = "Enter your message here"
	fmt.Println(message)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var cgpa float32 = 3.42
	fmt.Println(cgpa)

	// Dynamic type
	name := "Arslan Haider Sherazi"
	fmt.Println(name)

	age := 26
	fmt.Println(age)
}
