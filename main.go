package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/sensorario/bashutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Int int

func (integer Int) String() string {
	return strconv.Itoa(int(integer))
}

func main() {

	stay := true
	selection := 0
	for stay {
		bashutil.Clear()
		bashutil.Centerln("----------------------")
		bashutil.Centerln("| Menu                ")
		bashutil.Centerln("----------------------")
		bashutil.Centerln(" 1) play              ")
		bashutil.Centerln("----------------------")
		bashutil.Centerln(" 2) Replay past game  ")
		bashutil.Centerln("----------------------")
		bashutil.Centerln(" 3) Exit              ")
		bashutil.Centerln("----------------------")
		bashutil.Center(": ")

		if selection == menu["replay"] {
			showOldMatches(selection)
		}

		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		n, _ := strconv.ParseInt(scan.Text(), 10, 32)
		selection = int(n)

		if selection == menu["exit"] {
			os.Exit(0)
		}

		if selection == menu["play"] {
			stay = false
		}
	}

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
	for turn, m := range g.moves {
		fmt.Println(turnMessage(m, turn))
	}

	message := g.lastMessage()
	g.logMessage(message)
	fmt.Printf("\n" + message + "\n")

	g.logMessage(" --- il gioco e' terminato --- ")
	g.logMessage(" --- moves --- ")
	for turn, m := range g.moves {
		g.logMessage(turnMessage(m, turn))
	}
	g.logMessage(" --- moves --- \n")

	hasher := md5.New()
	hasher.Write([]byte(time.Now().String()))
	md5String := hex.EncodeToString(hasher.Sum(nil))
	saveMoves := md5String +
		"," + g.Players()[0].Name +
		"," + g.Players()[1].Name

	for _, m := range g.moves {
		saveMoves += "," + strconv.Itoa(m.position)
	}
	f, err := os.OpenFile("games", O_CREATE|O_RDWR|O_APPEND, 0777)
	check(err)
	w := bufio.NewWriter(f)
	f.WriteString(saveMoves + "\n")
	w.Flush()
}

func turnMessage(m move, turn int) string {
	return strconv.Itoa(turn+1) + ") " +
		m.player.Name + "'s turn : " +
		strconv.Itoa(m.position)
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
