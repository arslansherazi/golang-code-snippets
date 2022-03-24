package main

import "fmt"

func main() {
	var number int = 100
	var name string = "Arslan Haider Sherazi"
	const MAX int = 3
	var numbers = []int{20, 30, 40}

	// Defining a pointer
	/*
		The actual data type of the value of all pointers, whether integer, float, or otherwise, is the same,
		a long hexadecimal number that represents a memory address. The only difference between pointers of different
		data types is the data type of the variable or constant that the pointer points to.
	*/
	var ptr1 *int
	ptr1 = &number
	fmt.Println("Address of number:: ", ptr1)
	fmt.Println("Value of number using pointer:: ", *ptr1)

	fmt.Println()

	var ptr2 *string
	ptr2 = &name
	fmt.Println("Address of name:: ", ptr2)
	fmt.Println("Value of name using pointer:: ", *ptr2)

	fmt.Println()

	// Pointer to pointer
	var ptr3 *int
	ptr3 = &number
	var ptr4 **int
	ptr4 = &ptr3
	fmt.Println("Address of pointer:: ", ptr4)
	fmt.Println("Value of number using pointer to pointer:: ", **ptr4)

	fmt.Println()

	// Array of pointers
	var ptrs [MAX]*int
	for i := 0; i < MAX; i++ {
		ptrs[i] = &numbers[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("Address => %x \t Value => %d \n", ptrs[i], *ptrs[i])
	}
}
