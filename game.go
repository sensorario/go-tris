package main

import (
	"./src/tris"
	"./src/utils"
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	p1 := utils.GetUser("First player: ")
	p2 := utils.GetUser("Second player: ")

	var g tris.Game

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.AddPlayer(tris.Player{p1, "x"})
	g.AddPlayer(tris.Player{p2, "o"})

	fmt.Println("Simulate a real match")

	for 0 < g.AvailableTile() && false == g.TrisIsDone() {
		cell := randInt(1, 10)

		if true == g.IsAvailable(cell) {
			g.Play(cell)
			fmt.Println(g.OutputBoard())
		}
	}

	if true == g.TrisIsDone() {
		fmt.Printf("%s wins!!", g.NextPlayer().Name)
	} else {
		fmt.Println("Nobody wins")
	}
}
