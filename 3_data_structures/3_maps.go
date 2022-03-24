package main

import "fmt"

func main() {
	// creating map
	var capitals = map[string]string{
		"Pakistan":     "Islamabad",
		"Turkey":       "Istanbul",
		"Iran":         "tehran",
		"Saudi Arabia": "Riaz",
	}

	// iterating through map
	for key, value := range capitals {
		fmt.Println(key, " => ", value)
	}

	// accessing map using key
	capital_of_pakistan := capitals["Pakistan"]
	fmt.Println(capital_of_pakistan)

	// Delete key from map
	delete(capitals, "Saudi Arabia")
	fmt.Println(capitals)

	// Error handling
	var country string
	fmt.Print("Enter Country::")
	fmt.Scanln(&country)
	capital, is_capital_exists := capitals[country]
	if is_capital_exists {
		fmt.Printf("Capital of %s:: %s \n", country, capital)
	} else {
		fmt.Printf("Capital of %s does not exist in database \n", country)
	}
}
