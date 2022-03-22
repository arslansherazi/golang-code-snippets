package main

import "fmt"

func main() {
	const URL string = "https://www.google.com/"
	// URL = "test" // Gives error
	fmt.Println(URL)

	const VERSION int = 2
	fmt.Println(VERSION)
}
