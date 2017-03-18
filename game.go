package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"syscall"
)

const (
	O_RDONLY int = syscall.O_RDONLY
	O_CREATE int = syscall.O_CREAT
	O_RDWR   int = syscall.O_RDWR
	O_APPEND int = syscall.O_APPEND
)

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) Players() []Player {
	return g.players
}

func (g *Game) IsAvailable(position int) bool {
	for _, m := range g.moves {
		if m.position == position {
			return false
		}
	}

	return true
}

func (g *Game) Play(position int) int {
	if position < 0 || position > 9 {
		return -1
	}

	for _, m := range g.moves {
		if m.position == position {
			return -1
		}
	}

	currentPlayer := g.shouldPlay()

	g.moves = append(
		g.moves,
		move{currentPlayer, position, g.CurrentPlayer().Symbol},
	)

	for _, set := range winSets {
		if g.PlayerHasMovedInSet(currentPlayer, set.winSet) {
			g.trisIsDone = true
		}
	}

	return 0
}

func (g *Game) TrisIsDone() bool {
	return g.trisIsDone
}

func (g *Game) PlayerHasMovedInSet(p Player, positions [3]int) bool {
	setItemFound := 0

	for _, pos := range positions {
		if g.playerHasMovedIn(p, pos) {
			setItemFound++
		}
	}

	return setItemFound == 3
}

func (g *Game) PlayerHaveTwo(p Player, positions [3]int) bool {
	setItemFound := 0

	for _, pos := range positions {
		if g.playerHasMovedIn(p, pos) {
			setItemFound++
		}
	}

	return setItemFound == 2
}

func (g *Game) CurrentPlayer() Player {
	return g.players[len(g.moves)%2]
}

func (g *Game) NextPlayer() Player {
	return g.players[(len(g.moves)+1)%2]
}

func (g *Game) shouldPlay() Player {
	return g.players[len(g.moves)%2]
}

func (g *Game) AvailableTile() int {
	return 9 - len(g.moves)
}

var keys = map[int]string{
	1: "aa", 2: "ab", 3: "ac",
	4: "ba", 5: "bb", 6: "bc",
	7: "ca", 8: "cb", 9: "cc",
}

func (g *Game) render(board map[string]string) string {
	for _, m := range g.moves {
		board[keys[m.position]] = m.player.Symbol
	}

	return " " + board[keys[1]] + " | " + board[keys[2]] + " | " + board[keys[3]] + " \n" +
		"---|---|---\n" +
		" " + board[keys[4]] + " | " + board[keys[5]] + " | " + board[keys[6]] + " \n" +
		"---|---|---\n" +
		" " + board[keys[7]] + " | " + board[keys[8]] + " | " + board[keys[9]] + " \n"
}

func (g *Game) OutputHumanBoard() string {
	return g.render(map[string]string{
		"aa": "1", "ab": "2", "ac": "3",
		"ba": "4", "bb": "5", "bc": "6",
		"ca": "7", "cb": "8", "cc": "9",
	})
}

func (g *Game) OutputBoard() string {
	return g.render(map[string]string{
		"aa": " ", "ab": " ", "ac": " ",
		"ba": " ", "bb": " ", "bc": " ",
		"ca": " ", "cb": " ", "cc": " ",
	})
}

func (g *Game) playerHasMovedIn(p Player, position int) bool {
	for _, m := range g.moves {
		if m.player == p && m.position == position {
			return true
		}
	}

	return false
}

func (g *Game) GetRandomCell(min int, max int) int {
	message := []string{
		"Current user ",
		g.CurrentPlayer().Name,
	}
	g.logMessage(strings.Join(message, ""))

	for _, set := range winSets {
		if g.PlayerHaveTwo(g.CurrentPlayer(), set.winSet) {
			g.logMessage("buona scelta")
			for _, s := range set.winSet {
				if true == g.IsAvailable(s) {
					return s
				}
			}
		}
	}

	return min + rand.Intn(max-min)
}

func (g *Game) logMessage(message string) {
	file, err := os.OpenFile("go-tris.log", O_CREATE|O_RDWR|O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(message)
}
