package tris

import (
	"reflect"
	"testing"
)

func TestGameStartsInIdleStatus(t *testing.T) {
	var g Game
	if "Idle" != g.status() {
		t.Error(
			"g.status() must be Idle instead of",
			g.status(),
		)
	}
}

func TestGameMovesIsAnEmptyArray(t *testing.T) {
	var g Game
	if 0 != len(g.turns()) {
		t.Error(
			"Game should not have",
			len(g.turns()),
			"items!!!",
		)
	}
}

func TestGameAcceptPlayers(t *testing.T) {
	var g Game
	var p Player = Player{Name: "Simone"}
	g.AddPlayer(p)
}

func TestGameStartsWhenHasTwoPlayers(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	if "Started" != g.status() {
		t.Error("status must be Started")
	}
}

func TestGameMovesCountEachTurnPlayed(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
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
	var g Game
	players := [2]string{
		"Simone",
		"Demo",
	}
	g.AddPlayer(Player{players[0]})
	g.AddPlayer(Player{players[1]})
	turnToPlay := 9
	i := 0
	for _ = range players {
		if players[i] != g.shouldPlay().Name {
			t.Error(
				players[0],
				"should play the Game but",
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
	var g Game
	if reflect.TypeOf(g.board).Name() != "board" {
		t.Error("A Game must have its own board")
	}
}

func TestNumberOfFreeTiles(t *testing.T) {
	var g Game
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
	var tests = []struct {
		playerName string
		symbol     string
	}{
		{"Foo", "X"},
		{"Bar", "O"},
	}

	var g Game
	g.AddPlayer(Player{"Foo"})
	g.AddPlayer(Player{"Bar"})

	for _, test := range tests {
		if test.playerName != g.whoHasSymbol(test.symbol).Name {
			t.Error("Second player myst have `O` as symbol")
		}
	}
}

func TestInvalidPositionReturnNegativeUnit(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	if -1 != g.play(42) {
		t.Error(
			"42 should not be valid as position",
		)
	}
}

func TestValidGameReturnsZero(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	for i := 1; i < 10; i++ {
		if playResult := g.play(i); 0 != playResult {
			t.Error(
				"g.play("+string(i)+") should return zero instead of",
				string(playResult),
			)
		}
	}
}

func TestTileCannotBeSelectedTwice(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	g.play(3)
	if -1 != g.play(3) {
		t.Error("g.play(position int) should not accept same position twice")
	}
}

func TestWhenAllTilesAreOccupiedGameStatusIsEnd(t *testing.T) {
	var g Game
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	for i := 1; i <= 9; i++ {
		g.play(i)
	}
	if "End" != g.status() {
		t.Error("g.status() must be End when there are no more cells")
	}
}

func TestCannotPlayMoreThanNineTimes(t *testing.T) {
	var g Game
	var result int
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	for i := 1; i <= 10; i++ {
		result = g.play(i)
	}
	if result != -1 {
		t.Error("FAIL")
	}
}

func TestSetPresenceInTurns(t *testing.T) {
	var tests = []struct {
		turns      []int
		set        [3]int
		trisIsDone bool
	}{
		{[]int{1, 2, 4, 3, 7}, [3]int{1, 4, 7}, true},
		{[]int{1, 2, 4, 3, 7}, [3]int{1, 3, 7}, false},
		{[]int{2, 1, 5, 3, 8}, [3]int{2, 5, 8}, true},
		{[]int{1, 2, 5, 3, 9}, [3]int{1, 5, 9}, true},
		{[]int{3, 2, 6, 4, 9}, [3]int{3, 6, 9}, true},
	}
	for _, test := range tests {
		var g Game
		g.AddPlayer(Player{"Simone"})
		g.AddPlayer(Player{"Demo"})
		for _, move := range test.turns {
			g.play(move)
		}
		result := g.PlayerHasMovedInSet(Player{"Simone"}, test.set)
		if test.trisIsDone != result {
			t.Errorf("Set %d, %d, %d = %v", test.set[0], test.set[1], test.set[2], test.trisIsDone)
		}
	}
}

func TestTrisIsDone(t *testing.T) {
	var tests = []struct {
		turns []int
		set   [3]int
	}{
		{[]int{1, 5, 2, 4, 3}, [3]int{1, 2, 3}},
		{[]int{4, 2, 5, 1, 6}, [3]int{4, 5, 6}},
		{[]int{7, 1, 8, 2, 9}, [3]int{7, 8, 9}},
		{[]int{1, 2, 4, 3, 7}, [3]int{1, 4, 7}},
		{[]int{2, 1, 5, 3, 8}, [3]int{2, 5, 8}},
		{[]int{1, 2, 5, 3, 9}, [3]int{1, 5, 9}},
		{[]int{3, 2, 6, 4, 9}, [3]int{3, 6, 9}},
		{[]int{3, 2, 5, 4, 7}, [3]int{3, 5, 7}},
	}
	for _, test := range tests {
		var g Game
		g.AddPlayer(Player{"Simone"})
		g.AddPlayer(Player{"Demo"})
		for _, move := range test.turns {
			g.play(move)
		}
		if g.TrisIsDone() != true {
			t.Errorf("TrisIsDone is not working with %d,%d,%d. ", test.set[0], test.set[1], test.set[2])
		}
	}
}
