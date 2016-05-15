package tris

type player struct {
	name string
}

type game struct {
	currentPlayer string
	currentTurn   int
	players       []player
}

type board struct {
	tiles [9]tile
}

type tile struct {
	played bool
}

func (g *game) status() (s string) {
	if playerCount := len(g.players); playerCount == 2 {
		return "Started"
	}
	return "Idle"
}

func (g *game) addPlayer(p player) {
	g.players = append(g.players, p)
}

func (g *game) shouldPlay() (p player) {
	p = g.players[g.currentTurn%2]
	return
}

func (g *game) play(position int) {
	g.currentTurn++
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
