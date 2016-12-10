package main

import (
	"./src/console"
	"./src/tris"
	"./src/utils"
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

func main() {
	console.Clear()

	p1 := utils.GetUser("First player: ")
	p2 := "Computer"

	var g tris.Game

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.AddPlayer(tris.Player{p1, "x"})
	g.AddPlayer(tris.Player{p2, "o"})

	fmt.Println("Simulate a real match")

	var cell int
	for 0 < g.AvailableTile() && false == g.TrisIsDone() {
		console.Clear()

		cell = randInt(1, 10)

		if g.CurrentPlayer().Name == p1 {
			fmt.Println("Available moves:\n")
			fmt.Println(g.OutputHumanBoard())
			fmt.Println("Clean board:\n")
			fmt.Println(g.OutputBoard())
			fmt.Print("Type a number between 1 and 9 (your choice): ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			n, _ := strconv.ParseInt(scan.Text(), 10, 32)
			cell = int(n)
		}

		if true == g.IsAvailable(cell) {
			g.Play(cell)
			fmt.Println("\nFinal board:\n")
			fmt.Println(g.OutputBoard())
		}
	}

	if true == g.TrisIsDone() {
		fmt.Printf("%s wins!!\n\n", g.NextPlayer().Name)
	} else {
		fmt.Println("Nobody wins")
	}
}
