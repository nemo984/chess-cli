package data

import (
	"database/sql"
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
	fen TEXT,
	created TEXT DEFAULT CURRENT_TIMESTAMP, 
	updated TEXT
);`

func CreateTable() {
	statement, err :=  db.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}


func CreateGame(game Game) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(game)
	query := `INSERT INTO games(gameName, color, engineColor,colorTurn,engine,engineDepth,engineNodes,outcome,fen,updated) 
			values(:gameName,:color,:engineColor,:colorTurn,:engine,:engineDepth,:engineNodes,:outcome,:fen,datetime('now'))`
	_, err := db.NamedExec(query, game)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func UpdateGame(game Game) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(game)
	db.NamedExec(`UPDATE games SET gameName=:gameName, color=:color,engineColor=:engineColor,
				colorTurn=:colorTurn, engine=:engine, engineDepth=:engineDepth, engineNodes=:engineNodes,
				outcome=:outcome, fen=:fen, updated=datetime('now') WHERE id =:id`, game)

} 

func GetGame(name string) (Game,bool) {
	var game Game
	err := db.Get(&game,"SELECT * FROM games WHERE gameName = ? LIMIT 1",name)
	if err != nil || err == sql.ErrNoRows {
		return game,false
	}
	return game,true
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