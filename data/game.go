package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nemo984/chess-cli/models"
)

var _db *sqlx.DB

func OpenDatabase() error {
	var err error

	_db, err = sqlx.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return _db.Ping()
}

//Create a Table with GameSchema
func CreateTable() error {
	statement, err := _db.Prepare(models.GameSchema)
	if err != nil {
		return err
	}
	if _, err := statement.Exec(); err != nil {
		return err
	}
	return nil
}

//Game manages Game CRUD
type Game struct {
}

//Inserts a new Game into database
func (g *Game) Insert(game *models.Game) error {
	game.Created = time.Now().Format(time.RFC3339)
	query := `INSERT INTO games(gameName, color, engineColor,colorTurn,engine,engineDepth,outcome,method,fen,pgn,created,updated) 
			values(:gameName,:color,:engineColor,:colorTurn,:engine,:engineDepth,:outcome,:method,:fen,:pgn,:created,:updated)`
	_, err := _db.NamedExec(query, game)
	return err
}

//Update updates an existing Game
func (g *Game) Update(game models.Game) error {
	_, err := _db.NamedExec(`UPDATE games SET gameName=:gameName, color=:color,engineColor=:engineColor,
				colorTurn=:colorTurn, engine=:engine, engineDepth=:engineDepth,
				outcome=:outcome, method=:method, fen=:fen, pgn=:pgn, updated=:updated WHERE gameName =:gameName`, game)
	return err
}

//GetByName find a Game by name - bool true if game exists
func (g *Game) GetByName(name string) (models.Game, bool) {
	var game models.Game
	err := _db.Get(&game, "SELECT * FROM games WHERE gameName = ? LIMIT 1", name)
	if err != nil || err == sql.ErrNoRows {
		return game, false
	}
	return game, true
}

//GetAll gets the list of Game
func (g *Game) GetAll() ([]models.Game, error) {
	games := []models.Game{}
	err := _db.Select(&games, "SELECT * FROM games ORDER BY updated DESC")
	if err != nil {
		return nil, err
	}
	return games, nil
}

//Delete deletes a Game by Name
func (g *Game) DeleteByName(name string) error {
	res, err := _db.Exec("DELETE FROM games WHERE gameName=$1", name)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Game \"%v\" doesn't exist", name)
	}
	return nil
}
