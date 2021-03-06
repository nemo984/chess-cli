package models

var GameSchema = `CREATE TABLE IF NOT EXISTS games (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	gameName TEXT,
	color TEXT,
	engineColor	TEXT,
	colorTurn TEXT,
	engine TEXT,
	engineDepth INT,
	outcome TEXT,
	method TEXT,
	fen TEXT,
	pgn TEXT,
	created TEXT DEFAULT CURRENT_TIMESTAMP, 
	updated TEXT
);`
