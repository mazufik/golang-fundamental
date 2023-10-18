package main

import "fmt"

func main() {
	// cara pertama
	// var languages [5]string
	// languages[0] = "Ruby"
	// languages[1] = "Go"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Python"
	//
	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	// cara kedua
	// city := [6]string{"Ambon", "Bandung", "Jakarta", "Surabaya", "Banajrmasin", "Ternate"}
	//
	// fmt.Println(city)
	// lengths := len(city)
	// fmt.Println(lengths)

	//  mengisi array dari looping
	// var numbers []int
	// for i := 1; i < 20; i++ {
	// 	if i%2 == 0 {
	// 		numbers = append(numbers, i)
	// 	}
	// }
	//
	// fmt.Println(numbers)

	// looping untuk menampilkan array
	languages := [...]string{
		"Go",
		"Ruby",
		"Javascript",
		"C#",
		"Python",
		"Java",
		"Kotlin",
		"F#",
		"Elixir",
	}

	for index, lang := range languages {
		fmt.Println("index : ", index, " language : ", lang)
	}

}
