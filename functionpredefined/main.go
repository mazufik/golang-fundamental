package main

import "fmt"

func main() {
	luas, keliling := calculate(10, 15)
	fmt.Println(luas)
	fmt.Println(keliling)
}

// Basic function dengan predefined return value
func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}
