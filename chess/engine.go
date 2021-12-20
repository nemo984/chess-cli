package chess

import (
	"fmt"
	"log"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

type Engine struct {
	Path  string
	eng   *uci.Engine
	Nodes int
	Depth int
	Color chess.Color
}

func (e *Engine) setUp() {
	eng, err := uci.New(e.Path)
	if err != nil {
		log.Fatal(err)
	}
	e.eng = eng
	// initialize uci with new game
	if err := e.eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) setColor(color chess.Color) {
	e.Color = color
}

func (e Engine) getMoveAndMove() (exit bool, save bool) {
	cmdPos := uci.CmdPosition{Position: Game.Position()}
	cmdGo := uci.CmdGo{Depth: e.Depth, Nodes: e.Nodes}

	if err := e.eng.Run(cmdPos, cmdGo); err != nil {
		log.Fatal(err)
	}

	move := e.eng.SearchResults().BestMove
	if err := Game.Move(move); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Engine Move:", move)
	return false, true
}
