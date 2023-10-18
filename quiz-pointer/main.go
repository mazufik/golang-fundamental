package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// method AddGame
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zaydan"}

	fmt.Println("Gamer Name : ", gamer.Name)

	fmt.Println("Game list : ")
	gamer.AddGame("Mario")
	gamer.AddGame("Borderlands")
	gamer.AddGame("Bioshock")
	gamer.AddGame("GTA 5")

	for index, game := range gamer.Games {
		fmt.Println(index+1, ".", game)
	}
}
