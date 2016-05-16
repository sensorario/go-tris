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

func TestGameMovesIsAnEmptyArray(t *testing.T) {
	var g game
	if 0 != len(g.turns()) {
		t.Error(
			"Game should not have",
			len(g.turns()),
			"items!!!",
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

func TestGameMovesCountEachTurnPlayed(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	g.play(7)
	if 1 != len(g.turns()) {
		t.Error(
			"Game should not have",
			len(g.turns()),
			"items!!!",
			"But one!!",
		)
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
		g.play(7)
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

func TestGameHasItsOwnBoard(t *testing.T) {
	var g game
	if reflect.TypeOf(g.board).Name() != "board" {
		t.Error("A game must have its own board")
	}
}

func TestNumberOfFreeTiles(t *testing.T) {
	var g game
	freeTiles := 0
	for _, tt := range g.board.cells() {
		if tt.isFree() == true {
			freeTiles++
		}
	}
	if freeTiles != 9 {
		t.Error("There are not enough free tiles")
	}
}

func TestSecondPlayersSymbol(t *testing.T) {
	var g game
	g.addPlayer(player{"Foo"})
	g.addPlayer(player{"Bar"})
	if "Bar" != g.whoHasSymbol("O").name {
		t.Error("Second player myst have `O` as symbol")
	}
}

func TestInvalidPositionReturnNegativeUnit(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	if -1 != g.play(42) {
		t.Error(
			"42 should not be valid as position",
		)
	}
}

func TestValidPositionReturnZero(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	if 0 != g.play(3) {
		// todo: improve error message
		t.Error("g.play(position int) should return zero")
	}
}

func TestTileCannotBeSelectedTwice(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	g.play(3)
	if -1 != g.play(3) {
		t.Error("g.play(position int) should not accept same position twice")
	}
}

func TestWhenAllTilesAreOccupiedGameStatusIsEnd(t *testing.T) {
	var g game
	g.addPlayer(player{"Simone"})
	g.addPlayer(player{"Demo"})
	for i := 1; i <= 9; i++ {
		g.play(i)
	}
	if "End" != g.status() {
		t.Error("g.status() must be End when there are no more cells")
	}
}

func TestGameEndsWhenTrisIsDone(t *testing.T) {}

func TestPositionHasValues(t *testing.T) {}

func TestBotChooseBetterChoice(t *testing.T) {}
