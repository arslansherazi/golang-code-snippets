package main

import "fmt"

func main() {
	var marks int
	fmt.Print("Enter Marks::")
	fmt.Scanln(&marks)

	if marks >= 85 {
		fmt.Println("Your GPA: 4.0")
	} else if marks >= 80 {
		fmt.Println("Your GPA: 3.7")
	} else if marks >= 75 {
		fmt.Println("Your GPA: 3.3")
	} else if marks >= 70 {
		fmt.Println("Your GPA: 3.0")
	} else if marks >= 65 {
		fmt.Println("Your GPA: 2.7")
	} else if marks >= 60 {
		fmt.Println("Your GPA: 2.3")
	} else if marks >= 50 {
		fmt.Println("Your GPA: 2.0")
	} else {
		fmt.Println("You failed this course")
	}
}
