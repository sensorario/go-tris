package tris

type game struct {
	currentPlayer string
	currentTurn   int
	players       []string
}

type player struct {
	name string
}

func (g *game) status() (s string) {
	if playerCount := len(g.players); playerCount == 2 {
		return "Started"
	}

	return "Idle"
}

func (g *game) addPlayer(p player) {
	g.players = append(g.players, p.name)
}

func (g *game) shouldPlay() (s string) {
	s = g.players[g.currentTurn%2]
	return
}

func (g *game) play(position int) {
	g.currentTurn++
}
