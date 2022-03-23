package root

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")
}
