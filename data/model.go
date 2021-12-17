package data

type Game struct {
	GameName      string
	Color         string
	ComputerColor string
	ColorTurn     string
	Engine        string
	EngineDepth   int
	EngineNodes   int
	Outcome       string
	Pgn           string
}