package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB;

func OpenDatabase() error {
	var err error;

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS games (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"gameName" TEXT,
		"color" TEXT,
		"computerColor"	TEXT,
		"colorTurn" TEXT,
		"engine" TEXT,
		"engineDepth" INT,
		"engineNodes" INT,
		"outcome" TEXT,
		"pgn" TEXT,
		"created" TEXT DEFAULT CURRENT_TIMESTAMP, 
		"updated" TEXT
	);`

	statement, err :=  db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Games table created")
}

func SaveGame(game Game) {
	stmt, err := db.Prepare(`INSERT INTO games(gameName, color, computerColor,colorTurn,engine,engineDepth,engineNodes,outcome,pgn,updated) 
							values(?,?,?,?,?,?,?,?,?,datetime('now'))`)
	if err != nil {
		log.Fatal(err.Error()) 
	}
	_, err = stmt.Exec(game.GameName, game.Color, game.ComputerColor, game.ColorTurn, game.Engine, game.EngineDepth, game.EngineDepth,game.Outcome,game.Pgn)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func DisplayGames() {
	rows, err := db.Query(`SELECT * from games`)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(rows)
	rows.Close()
}