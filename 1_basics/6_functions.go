package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(num int) bool {
	if num == 1 || num == 0 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%2 == 0 {
			return false
		}
	}
	return true
}

// pass by reference
func swap(num1 *int, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}

// returning multiple values
func swapValues(val1, val2 string) (string, string) {
	return val2, val1
}

func main() {
	var num int
	fmt.Print("Enter Number::")
	fmt.Scanln(&num)
	fmt.Println(isPrimeNumber(num))
	if isPrimeNumber(num) {
		fmt.Printf("%d is prime number \n", num)
	} else {
		fmt.Printf("%d is not prime number \n", num)
	}

	// pass by reference
	var num1, num2 int = 10, 20
	fmt.Println("values before calling swap function")
	fmt.Println(num1, num2)
	swap(&num1, &num2)
	fmt.Println("values after calling swap function")
	fmt.Println(num1, num2)

	// function as value (function variable)
	getSquareRoot := func(num float64) float64 {
		return math.Sqrt(num)
	}
	fmt.Println(getSquareRoot(16))

	// receiving muliple values
	val1, val2 := "PUCIT", "PU"
	fmt.Println(val1, val2)
	val1, val2 = swapValues(val1, val2)
	fmt.Println(val1, val2)
}
