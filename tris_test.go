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
