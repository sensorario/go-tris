package tris

import (
	"reflect"
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

func TestCurrentPlayerChangeAfterTurn(t *testing.T) {
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
		if players[i] != g.shouldPlay().name {
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

func TestBoardMustContainNineTiles(t *testing.T) {
	var b board
	if len(b.tiles) != 9 {
		t.Error(
			"Tiles must be 9",
		)
	}
}

func TestTileShouldBeFreeOrOccupied(t *testing.T) {
	var tile tile
	if tile.isFree() != true {
		t.Error("Just created tile must be free")
	}
}

func TestWhenPlayedTileIsNoMoreFree(t *testing.T) {
	var tile tile
	tile.play()
	if tile.isFree() != false {
		t.Error("Tile must be occupied when played")
	}
}

func TestBoardIsComposedByTiles(t *testing.T) {
	var b board
	for _, tt := range b.cells() {
		if "tile" != reflect.TypeOf(tt).Name() {
			t.Error(
				"Oops! Tile is of type",
				reflect.TypeOf(tt).Name(),
			)
		}
	}
}
