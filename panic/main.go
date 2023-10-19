package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	/*
		Panic digunakan untuk menampilkan stack trace error sekaligus menghentikan flow goroutine (karena main() juga merupakan
		goroutine, maka behaviour yang sama juga berlaku). Setelah ada panic, proses akan terhenti, apapun setelah tidak di-eksekusi kecuali
		proses yang sudah di-defer sebelumnya (akan muncul sebelum panic error).

		Recover berguna untuk meng-handle panic error. Pada saat panic error muncul, recover men-take-over goroutine yang sedang panic
		(pesan panic tidak akan muncul).
	*/
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
