package main

import (
	"./tris"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getUser(player string) string {
	fmt.Print(player)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}

func main() {
	p1 := getUser("First player: ")
	p2 := getUser("Second player: ")

	var g tris.Game

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.AddPlayer(tris.Player{p1, "x"})
	g.AddPlayer(tris.Player{p2, "o"})

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
