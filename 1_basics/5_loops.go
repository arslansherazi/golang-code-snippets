package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Print(a, " ")
	}

	fmt.Println()

	var b int = 10

	for b > 0 {
		fmt.Print(b, " ")
		b--
	}

	fmt.Println()
}
