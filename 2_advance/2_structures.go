package main

import "fmt"

type Book struct {
	id          int
	name        string
	author_name string
	price       int
	ratings     float32
}

func printBook(book Book) {
	fmt.Println("id: ", book.id)
	fmt.Println("name: ", book.name)
	fmt.Println("author: ", book.author_name)
	fmt.Println("price: ", book.price)
	fmt.Println("ratings: ", book.ratings)
}

func printBookUsingPointer(book *Book) {
	fmt.Println("id: ", book.id)
	fmt.Println("name: ", book.name)
	fmt.Println("author: ", book.author_name)
	fmt.Println("price: ", book.price)
	fmt.Println("ratings: ", book.ratings)
}

func main() {
	// creating a struct
	book1 := Book{
		id:          1,
		name:        "C++",
		author_name: "Tonny Gaddis",
		price:       500,
		ratings:     4.3,
	}
	fmt.Println(book1)

	var book2 Book
	book2.id = 2
	book2.name = "Calculus"
	book2.author_name = "Howard Anton"
	book2.price = 780
	book2.ratings = 3.9
	fmt.Println(book2)

	// passing struct as function argument
	book3 := Book{
		id:          3,
		name:        "Go Programming",
		author_name: "Mahesh Kumar",
		price:       590,
		ratings:     3.3,
	}
	printBook(book3)

	// Pointer to struct
	book4 := Book{
		id:          4,
		name:        "Telecom Billing",
		author_name: "Zara Ali",
		price:       1200,
		ratings:     2.8,
	}
	var bPtr *Book
	bPtr = &book4
	fmt.Println(*bPtr)
	printBookUsingPointer(bPtr)
}
