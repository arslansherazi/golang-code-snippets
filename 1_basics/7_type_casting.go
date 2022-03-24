package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = 20

	var mean float32
	// average = num1 / num2  Gives Error
	mean = float32(num1) / float32(num2)
	fmt.Println("Mean: ", mean)

	var average float32 = 20.90
	fmt.Println(int(average)) // prints integer part of average
}
