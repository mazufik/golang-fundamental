package management

import "fmt"

// struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// method struct User
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

// method struct Group
func (group Group) DisplayGroup() {
	fmt.Println("Group name : ", group.Name)
	fmt.Println("Member count : ", len(group.Users))
	fmt.Println("Member name : ")
	for index, user := range group.Users {
		fmt.Println(index+1, ". ", user.FirstName)
	}
}
