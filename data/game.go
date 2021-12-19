package data

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nemo984/chess-cli/models"
)

var _db *sqlx.DB;

func OpenDatabase() error {
	var err error;

	_db, err = sqlx.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return _db.Ping()
}

//Create a Table with GameSchema
func CreateTable() {
	statement, err :=  _db.Prepare(models.GameSchema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

//Game manages Game CRUD
type Game struct {

}

//Inserts a new Game into database
func (g *Game) Insert(game models.Game) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Create Game",game)
	query := `INSERT INTO games(gameName, color, engineColor,colorTurn,engine,engineDepth,engineNodes,outcome,fen,pgn,updated) 
			values(:gameName,:color,:engineColor,:colorTurn,:engine,:engineDepth,:engineNodes,:outcome,:fen,:pgn,datetime('now'))`
	_, err := _db.NamedExec(query, game)
	return err
}

//Update updates an existing Game
func (g *Game) Update(game models.Game) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Update game",game)
	_,err := _db.NamedExec(`UPDATE games SET gameName=:gameName, color=:color,engineColor=:engineColor,
				colorTurn=:colorTurn, engine=:engine, engineDepth=:engineDepth, engineNodes=:engineNodes,
				outcome=:outcome, fen=:fen, pgn=:pgn, updated=datetime('now') WHERE gameName =:gameName`, game)
	return err
} 

//GetByName find a Game by name - bool true if game exists
func (g *Game) GetByName(name string) (models.Game,bool) {
	var game models.Game
	err := _db.Get(&game,"SELECT * FROM games WHERE gameName = ? LIMIT 1",name)
	if err != nil || err == sql.ErrNoRows {
		return game,false
	}
	return game,true
}

//GetAll gets the list of Game
func (g *Game) GetAll() ([]models.Game,error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	games := []models.Game{}
	err := _db.Select(&games, "SELECT * FROM games")
	if err != nil {
		return nil, err
	}
	return games,nil
}