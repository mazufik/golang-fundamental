package main

import (
	"exported/management"
	"fmt"
)

func main() {
	user := management.User{1, "Zaydan", "Rahman", "zaydan@ninetendo.com", true}
	fmt.Println(user.Display())

	user2 := management.User{2, "Link", "Awakening", "link@ninetendo.com", true}
	fmt.Println(user2.Display())

	users := []management.User{user, user2}

	group := management.Group{"Gamer", user, users, true}

	group.DisplayGroup()
}
