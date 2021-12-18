package chess

import (
	"fmt"

	"github.com/notnil/chess"
)

type Player struct {
	Color chess.Color
}

func (p Player) getMoveAndMove() (exit bool) {
	var input string
	for {
		fmt.Printf("Enter Your Move (%v): ", p.Color)
		fmt.Scanln(&input)
		if err := Game.MoveStr(input); err != nil {
			if input == "q" {
				return true
			}
			fmt.Println("Invalid Move, Try Again")
		} else {
			return false
		}
	}
}

func (p Player) resign() {
	Game.Resign(p.Color)
}

