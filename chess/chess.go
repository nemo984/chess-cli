package chess

import (
	"fmt"
	"os"

	"github.com/nemo984/chess-cli/data"
	"github.com/notnil/chess"
)

var Game *chess.Game


type playee interface{
	getMoveAndMove()
}

func StartGame(engine Engine) {
	Game = chess.NewGame()
	engine.setUp()
	engine.setColor(chess.Black)
	//TODO: func to choose color
	player := Player{chess.White}
	playees := []playee{
		player,
		engine,
	}

	fmt.Println("Game started")
	i := 0
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			fmt.Println(Game.Position().Board().Draw())
			playee.getMoveAndMove()
			i++
			if i == 3 {
				saveGame(player,engine)
				os.Exit(0)
			}
			if Game.Outcome() != chess.NoOutcome {
				fmt.Println("Game Over",Game.Outcome(),Game.Method())
				fmt.Println(Game)
				os.Exit(0)
			}
		}

    }
}

func saveGame(player Player, engine Engine) {
	pgn := Game.String()
	game := data.Game{
		Color: player.Color.String(),
		ComputerColor: engine.Color.String(),
		ColorTurn: Game.Position().Turn().String(),
		Engine: engine.Path,
		EngineDepth: engine.Depth,
		EngineNodes: engine.Nodes,
		Outcome: Game.Outcome().String(),
		Pgn: pgn,
	}

	data.SaveGame(game)
}