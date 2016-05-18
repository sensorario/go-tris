package main

import (
	"./tris"
	"fmt"
)

func main() {
	fmt.Println("Tris!")

	fmt.Println("Add players!")
	var g tris.Game
	g.AddPlayer(tris.Player{"Simone"})
	g.AddPlayer(tris.Player{"Demo"})

	fmt.Println("Simulate a real match")
	sequence := []int{1, 4, 2, 5, 3}
	for _, play := range sequence {
		if false == g.TrisIsDone() {
			g.Play(play)
		}
	}

	if true == g.TrisIsDone() {
		fmt.Printf("%s wins!!", g.NextPlayer().Name)
	} else {
		fmt.Println("Nobody wins")
	}
}
