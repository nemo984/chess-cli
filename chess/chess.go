package chess

import (
	"fmt"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)



var game *chess.Game

func StartGame(level int) {
	game = chess.NewGame()
	engine := Engine{level : level, path: "D:/Programming/go-workspace/stockfish.exe"}
	engine.setUp()
	fmt.Println("Game started, Difficulty Level:",level)
    for game.Outcome() == chess.NoOutcome {
		fmt.Println(game.Position().Board().Draw())
		getUserMove()
		fmt.Println(game.Position().Board().Draw())
		engine.move()
    }


}

func getUserMove()  {
	var input string
	for {
		fmt.Print("Enter Your Move: ")
		fmt.Scanln(&input)
		if err := game.MoveStr(input); err != nil {
			fmt.Println("Invalid Move, Try Again")
		} else {
			break
		}
	}
}

type Engine struct {
	level int
	path string
	eng *uci.Engine
}

func (e *Engine) setUp() {
	eng, err := uci.New(e.path)
	e.eng = eng
	if err != nil {
		panic(err)
	}
	// initialize uci with new game
	if err := e.eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
}


func (e Engine) move() {
	//TODO: generate move base on specified level
	move := e.eng.SearchResults().BestMove
	if err := game.Move(move); err != nil {
		panic(err)
	}
	fmt.Println("Engine Move ",move)
}

