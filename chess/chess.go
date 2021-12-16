package chess

import (
	"fmt"

	"github.com/notnil/chess"
)

func StartGame(level int) {
	game := chess.NewGame()
	fmt.Println("Game started, Difficulty Level:",level)
	fmt.Println(game.Position().Board().Draw())
}