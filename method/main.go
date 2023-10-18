package main

import "fmt"

// struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// method
func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

// struct ke 2 (embedded struct)
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// method Group
func (group Group) displayGroup() {
	fmt.Println("Group name : ", group.Name)
	fmt.Println("Member count : ", len(group.Users))
	fmt.Println("Member name : ")
	for index, user := range group.Users {
		fmt.Println(index+1, ". ", user.FirstName)
	}
}

func main() {
	// method adalah sebuah function yang dimiliki oleh sebuah objek
	// objek adalah hasil intensiasi dari sebuah struct
	user := User{1, "Zaydan", "Rahman", "zaydan@nintendo.com", true}
	/*
	   user = objek
	   proses mengisi objek disebut intensiasi
	*/

	result := user.display()
	fmt.Println(result)

	user2 := User{2, "Link", "Awakening", "link@ninetendo.com", true}
	fmt.Println(user2.display())

	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}
	group.displayGroup()
}
