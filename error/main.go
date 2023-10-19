package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		error merupakan sebuah tipe. Error memiliki 1 buah property berupa method Error() , method ini mengembalikan detail pesan
		error dalam string. Error termasuk tipe yang isinya bisa nil .
	*/

	var input string
	fmt.Print("Type same number : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
