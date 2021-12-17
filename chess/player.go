package chess

import (
	"fmt"

	"github.com/notnil/chess"
)

type Player struct {
	Color chess.Color
}

func (p Player) getMoveAndMove() {
	var input string
	for {
		fmt.Printf("Enter Your Move (%v): ", p.Color)
		fmt.Scanln(&input)
		if err := Game.MoveStr(input); err != nil {
			fmt.Println("Invalid Move, Try Again")
		} else {
			break
		}
	}
}

func (p Player) resign() {
	Game.Resign(p.Color)
}

