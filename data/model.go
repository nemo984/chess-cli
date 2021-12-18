package data

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
	FEN         string `db:"fen"`
	Created     string `db:"created"`
	Updated     string `db:"updated"`
}