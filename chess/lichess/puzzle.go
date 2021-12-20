package lichess

type game struct {
	PGN string `json:pgn`
}

type puzzle struct {
	Rating   int      `json:rating`
	Solution []string `json:solution`
}

type DailyPuzzle struct {
	Game   game   `json:game`
	Puzzle puzzle `json:puzzle`
}