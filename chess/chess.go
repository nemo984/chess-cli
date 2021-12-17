package chess

import (
	"fmt"
	"os"

	"github.com/notnil/chess"
)



var Game *chess.Game

func StartGame(level int) {
	Game = chess.NewGame()

	engine := Engine{level : level, path: "D:/Programming/go-workspace/stockfish.exe"}
	engine.setUp()
	//TODO: func to set up who vs who, in this case player vs player
	playees := playees{
		Player{chess.White},
		Player{chess.Black},
	}

	fmt.Println("Game started, Difficulty Level:",level)
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			fmt.Println(Game.Position().Board().Draw())
			playee.getMoveAndMove()
			if Game.Outcome() != chess.NoOutcome {
				fmt.Println("Game Over",Game.Outcome(),Game.Method())
				fmt.Println(Game)
				os.Exit(0)
			}
		}

    }
}
type playees []playee
//am i shakespare?
type playee interface{
	getMoveAndMove()
}
