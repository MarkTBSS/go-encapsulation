package main

import (
	"fmt"

	"github.com/MarkTBSS/go-encapsulation/player"
)

func main() {
	player1 := player.NewPlayer("John", 100)
	player1.Prefix = "Mr."
	fmt.Println("Prefix: ", player1.Prefix)
	// Accessing player information using getter methods
	fmt.Println("Player Name: ", player1.GetName())
	fmt.Println("Player Health: ", player1.GetHealth())

	// Modifying player information using setter methods
	player1.SetName("Jame")
	player1.SetHealth(80)
	player1.DisplayInfo()

	// Displaying updated player information
	player1.DeleteHealth(50)
	player1.DisplayInfo()

	player1.DeleteHealth(40)
	player1.DisplayInfo()
}
