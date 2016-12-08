package main

import (
	"./tris"
	"fmt"
	"strconv"
)

func main() {
	// Add players
	var g tris.Game
	g.AddPlayer(tris.Player{"Simone"})
	g.AddPlayer(tris.Player{"Demo"})

	// Play the game
	/*
		the 3x3 tris board

		| 1 | 2 | 3 |
		| 4 | 5 | 6 |
		| 7 | 8 | 9 |
	*/
	fmt.Println("Simulate a real match")
	cellSelectedInEachTurn := []int{
		1, // Simone
		4, // Demo
		2, // Simone
		5, // ..
		3,
	}
	for _, selectedCell := range cellSelectedInEachTurn {
		fmt.Println("Available tiles: " + strconv.Itoa(g.AvailableTile()))
		g.Play(selectedCell)
	}

	// Send result's feedback
	if true == g.TrisIsDone() {
		fmt.Printf("%s wins!!", g.NextPlayer().Name)
	} else {
		fmt.Println("Nobody wins")
	}
}
