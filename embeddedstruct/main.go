package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{1, "Zaydan", "Athaya", "zaydan@ninetendo.com", true}
	user2 := User{2, "Taufik", "Rahman", "taufik@ninetendo.com", true}

	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}

	displayGroup(group)
}

func displayGroup(group Group) {
	fmt.Println("Name : ", group.Name)
	fmt.Println("Member count : ", len(group.Users))
	fmt.Println("Member name : ")
	for index, user := range group.Users {
		fmt.Println(index+1, ". ", user.FirstName)
	}
}
