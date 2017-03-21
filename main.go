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

	level := "easy"
	if g.isHard == true {
		level = "hard"
	}
	fmt.Println("Level : " + level + "\n")
	for number, m := range g.moves {
		message := "" +
			"Move " + strconv.Itoa(number+1) +
			" ( " + m.player.Name + " ) " +
			" : " + strconv.Itoa(m.position)
		fmt.Println(message)
	}

	message := g.lastMessage()
	g.logMessage(message)
	fmt.Printf("\n" + message + "\n")

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

	saveMoves := g.Players()[0].Name + "," + g.Players()[1].Name
	for _, m := range g.moves {
		saveMoves += "," + strconv.Itoa(m.position)
	}
	f, err := os.OpenFile("games", O_CREATE|O_RDWR|O_APPEND, 0777)
	check(err)
	w := bufio.NewWriter(f)
	f.WriteString(saveMoves + "\n")
	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (g *Game) lastMessage() string {
	if true == g.TrisIsDone() {
		return g.NextPlayer().Name + " wins!!"
	}

	return "Nobody wins."
}
