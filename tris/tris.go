package tris

type Player struct {
	Name   string
	Symbol string
}

type Game struct {
	currentPlayer string
	players       []Player
	board         board
	moves         []move
	trisIsDone    bool
}

type move struct {
	player   Player
	position int
	symbol   string
}

type board struct {
	tiles [9]tile
}

type tile struct {
	played bool
}

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) IsAvailable(position int) bool {
	for _, m := range g.turns() {
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
	for _, m := range g.turns() {
		if m.position == position {
			return -1
		}
	}
	currentPlayer := g.shouldPlay()
	g.moves = append(
		g.moves,
		move{currentPlayer, position, g.CurrentPlayer().Symbol},
	)
	var winSets = []struct {
		winSet [3]int
	}{
		{[3]int{1, 2, 3}},
		{[3]int{4, 5, 6}},
		{[3]int{7, 8, 9}},
		{[3]int{1, 4, 7}},
		{[3]int{1, 5, 9}},
		{[3]int{2, 5, 8}},
		{[3]int{3, 6, 9}},
		{[3]int{3, 5, 7}},
	}
	for _, set := range winSets {
		if g.PlayerHasMovedInSet(
			currentPlayer,
			set.winSet,
		) {
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

func (g *Game) CurrentPlayer() Player {
	return g.players[len(g.turns())%2]
}

func (g *Game) NextPlayer() Player {
	return g.players[(len(g.turns())+1)%2]
}

func (g *Game) shouldPlay() (p Player) {
	p = g.players[len(g.turns())%2]
	return
}

func (g *Game) AvailableTile() int {
	return 9 - len(g.moves)
}

func (g *Game) OutputBoard() string {
	aa := " "
	ab := " "
	ac := " "

	ba := " "
	bb := " "
	bc := " "

	ca := " "
	cb := " "
	cc := " "

	for _, m := range g.turns() {
		if m.position == 1 {
			aa = m.symbol
		}
		if m.position == 2 {
			ab = m.symbol
		}
		if m.position == 3 {
			ac = m.symbol
		}
		if m.position == 4 {
			ba = m.symbol
		}
		if m.position == 5 {
			bb = m.symbol
		}
		if m.position == 6 {
			bc = m.symbol
		}
		if m.position == 7 {
			ca = m.symbol
		}
		if m.position == 8 {
			cb = m.symbol
		}
		if m.position == 9 {
			cc = m.symbol
		}
	}

	return " " + aa + " | " + ab + " | " + ac + " \n" +
		"---|---|---\n" +
		" " + ba + " | " + bb + " | " + bc + " \n" +
		"---|---|---\n" +
		" " + ca + " | " + cb + " | " + cc + " \n"
}

func (g *Game) turns() []move {
	return g.moves
}

func (g *Game) playerHasMovedIn(p Player, position int) bool {
	for _, m := range g.turns() {
		if m.player == p && m.position == position {
			return true
		}
	}
	return false
}
