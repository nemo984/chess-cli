package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB;

func OpenDatabase() error {
	var err error;

	db, err = sqlx.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

var schema = `CREATE TABLE IF NOT EXISTS games (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	gameName TEXT,
	color TEXT,
	engineColor	TEXT,
	colorTurn TEXT,
	engine TEXT,
	engineDepth INT,
	engineNodes INT,
	outcome TEXT,
	pgn TEXT,
	created TEXT DEFAULT CURRENT_TIMESTAMP, 
	updated TEXT
);`

func CreateTable() {
	statement, err :=  db.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Games table created")
}


func SaveGame(game Game) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(game)
	stmt, err := db.Prepare(`INSERT INTO games(gameName, color, engineColor,colorTurn,engine,engineDepth,engineNodes,outcome,pgn,updated) 
							values(?,?,?,?,?,?,?,?,?,datetime('now'))`)
	if err != nil {
		log.Fatal(err.Error()) 
	}
	_, err = stmt.Exec(game.GameName, game.Color, game.EngineColor, game.ColorTurn, game.Engine, game.EngineDepth, game.EngineDepth,game.Outcome,game.Pgn)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func GetGames() ([]Game,error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	games := []Game{}
	err := db.Select(&games, "SELECT * FROM games")
	if err != nil {
		return nil, err
	}
	return games,nil
}

// func GameExists(gameName string) bool {
// 	stmt,err := db.Prepare("EXISTS(SELECT 1 FROM games WHERE gameName = ?)");
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	res,err := stmt.Exec(gameName)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(res)
// }