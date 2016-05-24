package tris

import (
	"testing"
)

func TestInvalidPositionReturnNegativeUnit(t *testing.T) {
	if g := game(); -1 != g.Play(42) {
		t.Error(
			"42 should not be valid as position",
		)
	}
}

func TestValidGameReturnsZero(t *testing.T) {
	g := game()
	for i := 1; i < 10; i++ {
		if playResult := g.Play(i); 0 != playResult {
			t.Error(
				"g.Play("+string(i)+") should return zero instead of",
				string(playResult),
			)
		}
	}
}

func TestTileCannotBeSelectedTwice(t *testing.T) {
	g := game()
	g.Play(3)
	if -1 != g.Play(3) {
		t.Error("g.pLay(position int) should not accept same position twice")
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
		g := game()
		for _, move := range test.turns {
			g.Play(move)
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
		g := game()
		for _, move := range test.turns {
			g.Play(move)
		}
		if g.TrisIsDone() != true {
			t.Errorf("TrisIsDone is not working with %d,%d,%d. ", test.set[0], test.set[1], test.set[2])
		}
	}
}

func TestCurrentAndNextPlayerAreDifferentAndChangeInEachTurn(t *testing.T) {
	var g Game
	players := []Player{
		Player{"Simone"},
		Player{"Demo"},
	}
	g.AddPlayer(players[0])
	g.AddPlayer(players[1])
	for i := 1; i < 10; i++ {
		if g.NextPlayer() != players[i%2] {
			t.Error("Current player should be ", players[i%2])
		}
		g.Play(i)
		if g.CurrentPlayer() != players[i%2] {
			t.Error("Current player should be ", players[i%2])
		}
	}
}

func game() (g Game) {
	g.AddPlayer(Player{"Simone"})
	g.AddPlayer(Player{"Demo"})
	return g
}
