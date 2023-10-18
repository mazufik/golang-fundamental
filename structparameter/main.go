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
	user1 := User{1, "Zaydan", "Rahman", "zaydan@ninetendo.com", true}
	user2 := User{2, "Taufik", "Rahman", "taufik@ninetendo.com", true}

	displayuser1 := displayUser(user1)
	displayuser2 := displayUser(user2)

	fmt.Println(displayuser1)
	fmt.Println(displayuser2)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}
