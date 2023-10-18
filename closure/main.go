package main

import "fmt"

func main() {
	/*
		Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel. Dengan menerapkan konsep tersebut, kita bisa membuat
		fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.
		Closure merupakan anonymous function atau fungsi tanpa nama. Biasa dimanfaatkan untuk membungkus suatu proses yang hanya
		dipakai sekali atau dipakai pada blok tertentu saja.
	*/
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	numbers := []int{2, 4, 12, 7, 15, 30, 21, 12, 14}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}
