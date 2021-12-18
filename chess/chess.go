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
	getMoveAndMove() (exit bool)
}

func NewGame(engine Engine) {
	Game = chess.NewGame()
	gameName = "newName" //get game name from a flag
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
	log.Println("Continue Game",game)
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
	log.Println(playees)
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			fmt.Println(Game.Position().Board().Draw())
			exit := playee.getMoveAndMove()
			fmt.Println(Game.Position().Board().Draw())

			if exit || Game.Outcome() != chess.NoOutcome  {
				fmt.Println("Game Status: ",Game.Outcome(),Game.Method())
				_,exists := data.GetGame(gameName)
				saveGame(player,engine,exists)
				fmt.Println("Game",gameName,"Saved")
				os.Exit(0)
			}
		}

    }
}


func saveGame(player Player, engine Engine,update bool) {
	game := data.Game{
		Color: colorStr(player.Color),
		GameName: gameName,
		EngineColor: colorStr(engine.Color),
		ColorTurn: colorStr(Game.Position().Turn()),
		Engine: engine.Path,
		EngineDepth: engine.Depth,
		EngineNodes: engine.Nodes,
		Outcome: Game.Method().String(),
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