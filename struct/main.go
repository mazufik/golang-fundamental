package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	// // cara pertama
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Zaydan"
	// user.LastName = "Rahman"
	// user.Email = "zaydan@gmail.com"
	// user.IsActive = true
	//
	// fmt.Println(user)
	// fmt.Println(user.LastName)

	// // cara kedua
	// user := User{
	// 	ID:        1,
	// 	FirstName: "Zaydan",
	// 	LastName:  "Rahman",
	// 	Email:     "zaydan@gmail.com",
	// 	IsActive:  true,
	// }
	//
	// fmt.Println(user)
	// fmt.Println(user.Email)

	// cara ketiga ini harus urut dan benar-benar tahu key nya
	user := User{1, "Zaydan", "Rahman", "zaydan@gmail.com", true}

	fmt.Println(user)
}
