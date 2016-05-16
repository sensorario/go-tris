package tris

type player struct {
	name string
}

type move struct {
	player   player
	position int
}

type game struct {
	currentPlayer string
	players       []player
	board         board
	moves         []move
}

type board struct {
	tiles [9]tile
}

type tile struct {
	played bool
}

func (g *game) status() (s string) {
	if len(g.turns()) == 9 {
		return "End"
	}
	if playerCount := len(g.players); playerCount == 2 {
		return "Started"
	}
	return "Idle"
}

func (g *game) addPlayer(p player) {
	g.players = append(g.players, p)
}

func (g *game) shouldPlay() (p player) {
	p = g.players[len(g.turns())%2]
	return
}

func (g *game) play(position int) int {
	if position < 0 || position > 9 {
		return -1
	}
	for _, m := range g.turns() {
		if m.position == position {
			return -1
		}
	}
	g.moves = append(g.moves, move{g.shouldPlay(), position})
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

func (g *game) whoHasSymbol(s string) (p player) {
	if p = g.players[1]; s == "X" {
		p = g.players[0]
	}
	return
}

func (g *game) turns() []move {
	return g.moves
}
