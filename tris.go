package tris

type Game struct {
	numPlayers int
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
	g.numPlayers++
}
