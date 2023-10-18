package main

import (
	"errors"
	"fmt"
)

func main() {
	// // soal pertama
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println(total)

	// soal kedua
	result, err := calculate(10, 2, "=")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}

	return result, errorResult
}

func sum(numbers []int) (total int) {
	for _, number := range numbers {
		total += number
	}

	return
}
