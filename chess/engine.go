package chess

import (
	"fmt"

	"github.com/notnil/chess/uci"
)

type Engine struct {
	level int
	path  string
	eng   *uci.Engine
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

func (e Engine) getMoveAndMove() {
	//TODO: generate move base on specified level
	move := e.eng.SearchResults().BestMove
	if err := Game.Move(move); err != nil {
		panic(err)
	}
	fmt.Println("Engine Move ", move)
}
