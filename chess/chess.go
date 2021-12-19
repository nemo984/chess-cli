package chess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nemo984/chess-cli/data"
	"github.com/nemo984/chess-cli/models"
	"github.com/notnil/chess"
)


var Game *chess.Game
var gameDAO data.Game
var _gameName string

type playee interface{
	getMoveAndMove() (exit bool)
}

func NewGame(engine Engine) {
	Game = chess.NewGame()
	_gameName = "newName" //get game name from a flag
	engine.setUp()
	engine.setColor(chess.Black)
	//TODO: func to choose color
	player := Player{chess.White}
	playees := []playee{
		player,
		engine,
	}
	startGame(playees,player,engine)
}

func ContinueGame(name string) {
	game,ok := gameDAO.GetByName(name)
	if !ok {
		fmt.Println("Game doesn't exist")
		os.Exit(0)
	}
	if c := gameContinuable(game); !c {
		os.Exit(0)
	}

	log.Println("Continue Game",game)
	fen,err := chess.FEN(game.FEN)
	if err != nil {
		log.Fatal(err.Error())
	}
	Game = chess.NewGame(fen)
	_gameName = game.GameName

	player := Player{
		Color: strColor(game.Color),
	}
	engine := Engine{
		Path: game.Engine,
		Depth: game.EngineDepth,
		Nodes: game.EngineNodes,
		Color: strColor(game.EngineColor),
	}
	engine.setUp()

	playees := []playee{
		player,
		engine,
	}
	//black to move - engine goes first
	if strColor(game.ColorTurn) == chess.Black {
		playees[0], playees[1] = playees[1], playees[0]
	}
	startGame(playees,player,engine)

}

func startGame(playees []playee,player Player,engine Engine) {
	fmt.Println(Game.Position().Board().Draw())
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			exit := playee.getMoveAndMove()
			fmt.Println(Game.Position().Board().Draw())

			if exit || Game.Outcome() != chess.NoOutcome  {
				fmt.Println("Game Status: ",Game.Outcome(),Game.Method())
				_,exists := gameDAO.GetByName(_gameName)
				saveGame(player,engine,exists)
				fmt.Println("Game",_gameName,"Saved")
				os.Exit(0)
			}
		}

    }
}

func saveGame(player Player, engine Engine,update bool) {
	game := models.Game{
		Color: colorStr(player.Color),
		GameName: _gameName,
		EngineColor: colorStr(engine.Color),
		ColorTurn: colorStr(Game.Position().Turn()),
		Engine: engine.Path,
		EngineDepth: engine.Depth,
		EngineNodes: engine.Nodes,
		Outcome: Game.Outcome().String(),
		Method: Game.Method().String(),
		FEN: Game.FEN(),
		PGN: Game.String(),
	}
	if update {
		gameDAO.Update(game)
	} else {
		gameDAO.Insert(game)
	}
}

func colorStr(color chess.Color) string {
	if color == chess.White {
		return "White"
	}
	return "Black"
}

func strColor(color string) chess.Color {
	if color == "White" {
		return chess.White
	}
	return chess.Black
}

func gameContinuable(game models.Game) bool {
	if game.Outcome != chess.NoOutcome.String() {
		fmt.Printf("Game \"%v\" isn't continuable, Status: %v %v\n",game.GameName,game.Outcome,game.Method) 
		lichessURL,err := lichessAnalysisURL(game.PGN)
		if err != nil { //can't get lichess url
			os.Exit(0) 
		}
		fmt.Printf("Analyze on lichess: %v\n", lichessURL)
		return false
	}
	return true
}

func lichessAnalysisURL(pgn string) (string, error) {
	url := "https://lichess.org/api/import"
	values := map[string]string{"pgn": strings.TrimSpace(pgn)}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Non-OK HTTP status: %v",resp.StatusCode)
	}

	j := struct {
		ID string `json:"id"`
		URL string `json:"url"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return "",err
	}
	return j.URL, nil
}