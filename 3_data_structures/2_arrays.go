package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func main() {
	// 1D arrays
	var numbers = []int{20, 30, 40, 50, 60} // size of array = no of initialized array elements
	printArray(numbers)

	var nums [10]int
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	for i := 0; i < len(nums); i++ {
		random_number := rand.Intn(max-min+1) + min
		nums[i] = random_number
	}
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	// 2D Arrays
	var two_dim_numbers = [2][4]int{
		{20, 30, 40, 23},
		{10, 20, 60, 70},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(two_dim_numbers[i][j], " ")
		}
		fmt.Println()
	}

	// iterate through array
	var names = []string{"Arslan", "Waqar", "Fasi", "Haider"}
	for index := range names {
		fmt.Println(names[index])
	}
}
