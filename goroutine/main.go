package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	/*
		Goroutine mirip dengan thread, tapi sebenarnya bukan. Sebuah native thread bisa berisikan sangat banyak goroutine. Mungkin lebih pas
		kalau goroutine disebut sebagai mini thread. Goroutine sangat ringan, hanya dibutuhkan sekitar 2kB memori saja untuk satu buah
		goroutine. Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
	*/

	runtime.GOMAXPROCS(2)

	go print(5, "hello")
	print(5, "How are you")

	var input string
	fmt.Scanln(&input)
}
