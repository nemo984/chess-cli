package chess

import (
	"fmt"
	"os"

	"github.com/notnil/chess"
)

var Game *chess.Game

type playees []playee

type playee interface{
	getMoveAndMove()
}

func StartGame(engine Engine) {
	Game = chess.NewGame()
	engine.setUp()
	//TODO: func to set up who vs who, in this case player vs player
	playees := playees{
		Player{chess.White},
		engine,
	}

	fmt.Println("Game started")
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
