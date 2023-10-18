package main

import (
	"fmt"               // standart library
	"kedua/calculation" // package berbeda
)

func main() {
	fmt.Println("Halo, belajar golang")

	// package library
	result := calculation.Add(15, 12)

	fmt.Println(result)
}
