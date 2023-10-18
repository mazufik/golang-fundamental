package main

import "fmt"

func main() {
	// cara pertama membuat slice
	// var gamingConsoles []string
	// gamingConsoles = append(gamingConsoles, "Playstation 5")
	// gamingConsoles = append(gamingConsoles, "Ninetendo Switch")
	//
	// fmt.Println(gamingConsoles)

	// cara kedua membuat slice
	gamingConsoles := []string{
		"Playstation 5",
		"Ninetendo Switch",
		"Xbox One",
	}

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
