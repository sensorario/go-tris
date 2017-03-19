package main

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
	isHard        bool
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
