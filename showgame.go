package main

import (
	"bufio"
	"fmt"
	"github.com/sensorario/bashutil"
	"os"
	"strconv"
	"strings"
)

func showGame(theGame []string) {
	var g Game
	var pp map[int]string
	pp = make(map[int]string)
	for i := 1; i <= 9; i++ {
		pp[i] = " "
	}
	for key, position := range theGame {
		if key > 2 {
			conv, err := strconv.Atoi(position)
			check(err)
			symbol := "x"
			if key%2 == 0 {
				symbol = "x"
			} else {
				symbol = "o"
			}
			pp[conv] = symbol
		}
	}

	fmt.Println(g.render(map[string]string{
		"aa": pp[1],
		"ab": pp[2],
		"ac": pp[3],
		"ba": pp[4],
		"bb": pp[5],
		"bc": pp[6],
		"ca": pp[7],
		"cb": pp[8],
		"cc": pp[9],
	}))
}

func showOldMatches(selection int) {
	f, err := os.Open("games")
	if err != nil {
		fmt.Println("No file")
		return
	}
	defer f.Close()

	bashutil.Clear()
	scanner := bufio.NewScanner(f)
	matchNumber := 1
	var matches map[int]string
	matches = make(map[int]string)
	completeMathces := make(map[string]string)
	for scanner.Scan() {
		match := scanner.Text()[0:32]
		completeMathces[match] = scanner.Text()
		message := "match " + Int(matchNumber).String() + ") " + match
		matches[matchNumber] = match
		fmt.Println(message)
		matchNumber++
	}
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	selection = int(n)

	fmt.Println("Selection : " + Int(selection).String())
	if hashGame, ok := matches[selection]; ok {
		completeMathce := completeMathces[hashGame]
		fmt.Println("Game selected : " + hashGame)
		fmt.Println("Complete Game : " + completeMathce)
		showGame(strings.Split(completeMathce, ","))
	} else {
		message := "Game " + Int(selection).String() + " not found"
		fmt.Println(message)
	}
}
