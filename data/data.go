package data

import (
	"database/sql"
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
		"color" TEXT,
		"engine" TEXT,
		"computerColor"	TEXT,
		"level" INTEGER,
		"isWon" BOOLEAN,
		"isCheckmate" BOOLEAN,
		"isStalemate" BOOLEAN,
		"pgn" TEXT
	);`

	statement, err :=  db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Games table created")
}
