package main

import (
	"bufio"
	"github.com/sensorario/bashutil"
	"os"
)

func getUser() string {
	username := ""

	for username == "" {
		bashutil.Clear()
		bashutil.Center("Your name: ")

		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		username = scan.Text()
	}

	return username
}
