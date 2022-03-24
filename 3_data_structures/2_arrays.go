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
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30

	// 1D arrays
	var numbers = []int{20, 30, 40, 50, 60} // size of array = no of initialized array elements
	printArray(numbers)

	var nums [10]int
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
	for index, element := range names {
		fmt.Printf("index=%d, element=%s \n", index, element)
	}

	// Dynamic size array
	var random_numbers []int
	var count int
	fmt.Print("Enter count::")
	fmt.Scanln(&count)
	for i := 0; i <= count; i++ {
		random_number := rand.Intn(max-min+1) + min
		random_numbers = append(random_numbers, random_number)
	}
	fmt.Println(random_numbers)

	// slicing array
	var numbers2 = []int{11, 23, 45, 67, 24, 78, 94, 20, 12, 34, 90, 45}
	array_slice_1 := numbers2[2:5]
	array_slice_2 := numbers2[5:]
	array_slice_3 := numbers2[:5]
	fmt.Println(array_slice_1)
	fmt.Println(array_slice_2)
	fmt.Println(array_slice_3)
}
