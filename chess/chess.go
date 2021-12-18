package chess

import (
	"fmt"
	"log"
	"os"

	"github.com/nemo984/chess-cli/data"
	"github.com/notnil/chess"
)

var Game *chess.Game
var gameName string

type playee interface{
	getMoveAndMove()
}

func NewGame(engine Engine) {
	Game = chess.NewGame()
	engine.setUp()
	engine.setColor(chess.Black)
	//TODO: func to choose color
	player := Player{chess.White}
	playees := []playee{
		player,
		engine,
	}
	startGame(playees,player,engine)
}

func ContinueGame(name string) {
	game,ok := data.GetGame(name)
	if !ok {
		fmt.Println("Game doesn't exist")
		os.Exit(0)
	}
	fen,err := chess.FEN(game.FEN)
	if err != nil {
		log.Fatal(err.Error())
	}
	Game = chess.NewGame(fen)
	gameName = game.GameName

	player := Player{
		Color: strColor(game.Color),
	}
	engine := Engine{
		Path: game.Engine,
		Depth: game.EngineDepth,
		Nodes: game.EngineNodes,
		Color: strColor(game.EngineColor),
	}
	engine.setUp()

	playees := []playee{
		player,
		engine,
	}
	//black to move - engine goes first
	if strColor(game.ColorTurn) == chess.Black {
		playees[0], playees[1] = playees[1], playees[0]
	}
	startGame(playees,player,engine)

}


func startGame(playees []playee,player Player,engine Engine) {
	i := 0
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			fmt.Println(Game.Position().Board().Draw())
			playee.getMoveAndMove()
			fmt.Println(Game.Position().Board().Draw())
			i++
			if i > 5 {
				_,exists := data.GetGame(gameName)
				saveGame(player,engine,exists)
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


func saveGame(player Player, engine Engine,update bool) {
	game := data.Game{
		Color: colorStr(player.Color),
		GameName: "anotherone",
		EngineColor: colorStr(engine.Color),
		ColorTurn: colorStr(Game.Position().Turn()),
		Engine: engine.Path,
		EngineDepth: engine.Depth,
		EngineNodes: engine.Nodes,
		Outcome: Game.Outcome().String(),
		FEN: Game.FEN(),
	}
	if update {
		data.UpdateGame(game)
	} else {
		data.CreateGame(game)
	}
}

func colorStr(color chess.Color) string {
	if color == chess.White {
		return "White"
	}
	return "Black"
}

func strColor(color string) chess.Color {
	if color == "White" {
		return chess.White
	}
	return chess.Black
}