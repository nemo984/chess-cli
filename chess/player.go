package chess

import (
	"fmt"

	"github.com/notnil/chess"
)

type Player struct {
	color chess.Color
}

func (p Player) getMoveAndMove() {
	var input string
	for {
		fmt.Printf("Enter Your Move (%v): ", p.color)
		fmt.Scanln(&input)
		if err := Game.MoveStr(input); err != nil {
			fmt.Println("Invalid Move, Try Again")
		} else {
			break
		}
	}
}
