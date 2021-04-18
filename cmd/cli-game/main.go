package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/niakr1s/chess/chess"
)

func main() {
	g := chess.NewGame()
	fmt.Println("Moves should be entered in format like e2-e4.")

	scanner := bufio.NewScanner(os.Stdin)
loop:
	for {
		fmt.Println(g)
		fmt.Print(enterMoveStr())
		for scanner.Scan() {
			cmd := scanner.Text()

			cmd = strings.TrimSpace(cmd)
			splitted := strings.SplitN(cmd, "-", 2)
			if len(splitted) < 2 {
				fmt.Print(enterMoveStr())
			}

			err := g.MoveStr(splitted[0], splitted[1])
			if err != nil {
				fmt.Print(enterMoveStr())
				continue
			}
			fmt.Println()
			continue loop
		}
	}
}
func enterMoveStr() string {
	return "Enter valid move: "
}
