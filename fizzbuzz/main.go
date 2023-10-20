package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
