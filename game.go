package main

import (
	"bufio"
	"fmt"
	"github.com/sensorario/bashutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var cell int
	var g Game

	bashutil.Clear()

	bashutil.Center("Your name: ")
	p1 := getUser()
	p2 := "Computer"

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.AddPlayer(Player{p1, "x"})
	g.AddPlayer(Player{p2, "o"})

	fmt.Println("Simulate a real match")

	for 0 < g.AvailableTile() && false == g.TrisIsDone() {
		bashutil.Clear()

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
		} else {
			cell = GetRandomCell(1, 10)
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

func getUser() string {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	return scan.Text()
}
