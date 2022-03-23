package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Arslan Haider Sherazi"

	// length of string
	name_length := len(name)
	fmt.Println("Length of name = ", name_length)

	// hex bytes of string characters
	for i := 0; i < len(name); i++ {
		fmt.Print(name[i], " ")
	}
	fmt.Println()

	// Concate strings
	var first_name, last_name = "Arslan", "Haider"
	fmt.Println(first_name + " " + last_name)

	// Concatenating list of strings
	var student_details = []string{"Arslan Haider Sherazi", "PUCIT", "2014-2018", "3.42"}
	student_details_message := strings.Join(student_details, ", ")
	fmt.Println(student_details_message)
}
