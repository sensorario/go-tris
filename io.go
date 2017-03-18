package main

import (
	"bufio"
	"os"
)

func getUser() string {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	return scan.Text()
}
