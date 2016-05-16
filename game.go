package main

import (
	"./tris"
	"fmt"
)

func main() {
	fmt.Println("Tris!")

	fmt.Println("Add players!")
	var g tris.Game
	g.AddPlayer(tris.Player{"Simone"})
	g.AddPlayer(tris.Player{"Demo"})

	fmt.Println("Simulate a real match")

	fmt.Println("Declare game as ended")

	fmt.Println("Declare winner if there is one")
}
