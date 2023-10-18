package main

import "fmt"

func main() {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)

	// result := add(12, 23)
	// fmt.Println(result)

	luas, keliling := calculate(10, 15)
	fmt.Println(luas)
	fmt.Println(keliling)
}

// dua balikan
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// // satu balikan
// func add(number, numberTwo int) int {
// 	return number + numberTwo
// }

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar Golang"
// 	return newSentence
// }
