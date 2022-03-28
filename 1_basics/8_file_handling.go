package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Creating file
	file, error := os.Create("./students_details.txt")
	if error != nil {
		panic(error) // terminates the execution of program
	}

	// Writing to the file
	student_1 := "Arslan Haider Sherazi\nPUCIT\nSoftware Engineering\n3.42"
	student_2 := "Danish Ali\nPUCIT\nSoftware Engineering\n2.52"
	io.WriteString(file, student_1)
	io.WriteString(file, "\n\n")
	io.WriteString(file, student_2)

	defer file.Close() // executes at the last of program (LIFO)

	// Reading from file
	data_bytes, error := ioutil.ReadFile("./students_details.txt")
	if error != nil {
		panic(error)
	}
	fmt.Println(string(data_bytes))
}
