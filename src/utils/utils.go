package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetUser(player string) string {
	fmt.Print(player)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}
