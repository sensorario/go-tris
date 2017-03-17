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
	pHuman := getUser()
	pComputer := "Computer"

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	fmt.Println("Select first player")
	randomNumber := rand.Intn(2)

	if randomNumber == 0 {
		g.AddPlayer(Player{pHuman, "x"})
		g.AddPlayer(Player{pComputer, "o"})
	} else {
		g.AddPlayer(Player{pComputer, "x"})
		g.AddPlayer(Player{pHuman, "o"})
	}

	fmt.Println("Start match")
	turnNumber := 0

	for 0 < g.AvailableTile() && false == g.TrisIsDone() {
		bashutil.Clear()

		turnNumber++

		fmt.Println("Available moves:\n")
		fmt.Println(g.OutputHumanBoard())
		fmt.Println("Clean board:\n")
		fmt.Println(g.OutputBoard())
		fmt.Printf("Starter player is : %s", g.Players()[0].Name)
		fmt.Println("")
		fmt.Printf("Starter symbon is : %s", g.Players()[0].Symbol)
		fmt.Println("")
		fmt.Printf("Turn number : %d", turnNumber)
		fmt.Println("")
		fmt.Printf("Type a number between 1 and 9 (%s's turn): ", g.CurrentPlayer().Name)

		for {
			if g.CurrentPlayer().Name == "Computer" {
				cell = GetRandomCell(1, 10)
			} else {
				scan := bufio.NewScanner(os.Stdin)
				scan.Scan()
				n, _ := strconv.ParseInt(scan.Text(), 10, 32)
				cell = int(n)
			}
			if cell >= 1 && cell <= 9 {
				break
			}
		}

		if true == g.IsAvailable(cell) {
			g.Play(cell)
			fmt.Println("\nFinal board:\n")
			fmt.Println(g.OutputBoard())
		} else {
			turnNumber--
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
