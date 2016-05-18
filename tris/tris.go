package tris

type Player struct {
	Name string
}

type move struct {
	player   Player
	position int
}

type Game struct {
	currentPlayer string
	players       []Player
	board         board
	moves         []move
	trisIsDone    bool
}

type board struct {
	tiles [9]tile
}

type tile struct {
	played bool
}

func (g *Game) status() (s string) {
	if len(g.turns()) == 9 {
		return "End"
	}
	if playerCount := len(g.players); playerCount == 2 {
		return "Started"
	}
	return "Idle"
}

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) shouldPlay() (p Player) {
	p = g.players[len(g.turns())%2]
	return
}

func (g *Game) play(position int) int {
	if position < 0 || position > 9 {
		return -1
	}
	for _, m := range g.turns() {
		if m.position == position {
			return -1
		}
	}
	currentPlayer := g.shouldPlay()
	g.moves = append(g.moves, move{currentPlayer, position})
	if g.PlayerHasMovedInSet(currentPlayer, [3]int{1, 4, 7}) {
		g.trisIsDone = true
	}
	return 0
}

func (t *tile) isFree() bool {
	return !t.played
}

func (t *tile) play() {
	t.played = true
}

func (b *board) cells() (cc [9]tile) {
	return b.tiles
}

func (g *Game) whoHasSymbol(s string) (p Player) {
	if p = g.players[1]; s == "X" {
		p = g.players[0]
	}
	return
}

func (g *Game) turns() []move {
	return g.moves
}

func (g *Game) TrisIsDone() bool {
	return g.trisIsDone
}

func (g *Game) playerHasMovedIn(p Player, position int) bool {
	for _, m := range g.turns() {
		if m.player == p && m.position == position {
			return true
		}
	}
	return false
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
