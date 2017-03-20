package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sensorario/bashutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var cell int
	var g Game

	read := flag.String("level", "hard", "the foo value")
	flag.Parse()
	g.logMessage("Init computer level")
	if *read == "easy" {
		g.logMessage("Computer level: easy")
		g.playEasy()
	} else if *read == "hard" {
		g.logMessage("Computer level: hard")
		g.playHard()
	} else {
		g.playHard()
	}

	g.logMessage(" --- Che il gioco abbia inizio --- ")

	bashutil.Clear()

	bashutil.Center("Your name: ")
	pHuman := getUser()
	pComputer := "Computer"

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	g.logMessage("Select first player")
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
		var level string
		if g.isHard == true {
			level = "hard"
		} else {
			level = "easy"
		}
		computerLevelMessage := []string{"Computer level : ", level}
		fmt.Println(strings.Join(computerLevelMessage, ""))

		turnNumber++

		turnNumberMessage := []string{
			"Turn ",
			strconv.Itoa(10 - turnNumber),
			" of ",
			strconv.Itoa(turnNumber),
		}
		fmt.Println(strings.Join(turnNumberMessage, ""))

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

		rawMessage := []string{
			g.CurrentPlayer().Name,
			"'s turn",
		}
		message := strings.Join(rawMessage, "")
		g.logMessage(message)

		for {
			if g.CurrentPlayer().Name == "Computer" {
				cell = g.GetRandomCell(1, 10)
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

	for number, m := range g.moves {
		message := []string{
			"Move ",
			strconv.Itoa(number + 1),
			" ( ",
			m.player.Name,
			" ) ",
			" : ",
			strconv.Itoa(m.position),
		}
		fmt.Println(strings.Join(message, ""))
	}

	if true == g.TrisIsDone() {
		raw := []string{
			g.NextPlayer().Name,
			" wins!!",
		}
		message := strings.Join(raw, "")
		g.logMessage(message)
		fmt.Printf(message)
	} else {
		fmt.Println("Nobody wins")
	}

	g.logMessage(" --- il gioco e' terminato --- ")
	g.logMessage(" --- moves --- ")
	for number, m := range g.moves {
		message := []string{
			"Move ",
			strconv.Itoa(number + 1),
			" ( ",
			m.player.Name,
			" ) ",
			" : ",
			strconv.Itoa(m.position),
		}
		g.logMessage(strings.Join(message, ""))
	}
	g.logMessage(" --- moves --- \n")
}
