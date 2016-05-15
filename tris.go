package tris

type Game struct {
	numPlayers    int
	currentPlayer string
}

type Player struct {
	name string
}

func (g *Game) Status() (s string) {
	if g.numPlayers == 2 {
		return "Started"
	}

	return "Idle"
}

func (g *Game) AddPlayer(p Player) {
	if g.numPlayers == 0 {
		g.currentPlayer = p.name
	}
	g.numPlayers++
}

func (g *Game) shouldPlay() (s string) {
	s = g.currentPlayer
	return
}
