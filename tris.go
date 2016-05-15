package tris

type Game struct {
}

type Player struct {
	name string
}

func (g *Game) Status() (s string) {
	s = "Idle"
	return
}

func (g *Game) AddPlayer(p Player) {
}
