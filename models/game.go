package models

import (
	"log"

	"github.com/nemo984/chess-cli/utils"
	"github.com/notnil/chess"
)

type Game struct {
	Id          int    `db:"id"`
	GameName    string `db:"gameName"`
	Color       string `db:"color"`
	EngineColor string `db:"engineColor"`
	ColorTurn   string `db:"colorTurn"`
	Engine      string `db:"engine"`
	EngineDepth int    `db:"engineDepth"`
	EngineNodes int    `db:"engineNodes"`
	Outcome     string `db:"outcome"`
	Method      string `db:"method"`
	FEN         string `db:"fen"`
	PGN         string `db:"pgn"`
	Created     string `db:"created"`
	Updated     string `db:"updated"`
}

func (g *Game) Resign() {
	fen, err := chess.FEN(g.FEN)
	if err != nil {
		log.Fatal(err.Error())
	}
	game := chess.NewGame(fen)
	game.Resign(utils.StrColor(g.Color))
	g.Outcome = game.Outcome().String()
	g.Method = game.Method().String()
	g.FEN = game.FEN()
	g.PGN = game.String()
}
