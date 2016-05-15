package tris

import (
	"testing"
)

func TestGameStartsInIdleStatus(t *testing.T) {
	var g game
	if "Idle" != g.status() {
		t.Error(
			"g.status() must be Idle instead of",
			g.status(),
		)
	}
}

func TestGameAcceptPlayers(t *testing.T) {
	var g game
	var p player = player{name: "Simone"}
	g.addPlayer(p)
}

func TestGameStartsWhenHasTwoPlayers(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	if "Started" != g.status() {
		t.Error("status must be Started")
	}
}

func TestFirstPlayerShouldStart(t *testing.T) {
	var g game
	firstPlayerName := "Foo"
	g.addPlayer(player{firstPlayerName})
	g.addPlayer(player{"Demo"})
	if firstPlayerName != g.shouldPlay() {
		t.Error(
			firstPlayerName,
			"should start the game but",
			g.shouldPlay(),
		)
	}
}

func TestCurentPlayerChangeAfterTurn(t *testing.T) {
	var g game
	players := [2]string{
		"Simone",
		"Demo",
	}
	g.addPlayer(player{players[0]})
	g.addPlayer(player{players[1]})
	turnToPlay := 9
	i := 0
	for _ = range players {
		if players[i] != g.shouldPlay() {
			t.Error(
				players[0],
				"should play the game but",
				g.players[turnToPlay%2],
			)
		}
		g.play(0)
		i++
	}
}
