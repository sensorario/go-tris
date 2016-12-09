package main

import (
	"./tris"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var g tris.Game

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.AddPlayer(tris.Player{"Simone", "x"})
	g.AddPlayer(tris.Player{"Demo", "o"})

	fmt.Println("Simulate a real match")

	for 0 < g.AvailableTile() && false == g.TrisIsDone() {
		cell := randInt(1, 10)
		if true == g.IsAvailable(cell) {
			fmt.Println("Available tiles: " + strconv.Itoa(g.AvailableTile()))
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
