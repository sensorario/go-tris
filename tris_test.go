package tris

import (
	"testing"
)

func TestGameStartsInIdleStatus(t *testing.T) {
	var g Game
	if "Idle" != g.Status() {
		t.Error("Status must be idle")
	}
}

func TestGameAcceptPlayers(t *testing.T) {
	var g Game
	var p Player = Player{name: "Simone"}
	g.AddPlayer(p)
}

func TestGameStartsWhenHasTwoPlayers(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	if "Started" != g.Status() {
		t.Error("Status must be Started")
	}
}

func TestFirstPlayerShouldStart(t *testing.T) {
	var g Game
	firstPlayerName := "Foo"
	g.AddPlayer(Player{firstPlayerName})
	g.AddPlayer(Player{"Demo"})
	if firstPlayerName != g.shouldPlay() {
		t.Error(
			"Luca should start the game but",
			g.currentPlayer,
		)
	}
}
