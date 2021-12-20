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
	"github.com/nemo984/chess-cli/utils"
	"github.com/notnil/chess"
)

var (
	Game *chess.Game
	gameDAO data.Game
	_gameName string
)

type playee interface{
	getMoveAndMove() (exit bool, save bool)
}

func NewGame(engine Engine, name string) {
	Game = chess.NewGame()
	_gameName = name
	engine.setUp()
	engine.setColor(chess.Black)
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
		fmt.Printf("Game \"%v\" doesn't exist", name)
		os.Exit(0)
	}
	if c := GameContinuable(game); !c {
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
		Color: utils.StrColor(game.Color),
	}
	engine := Engine{
		Path: game.Engine,
		Depth: game.EngineDepth,
		Nodes: game.EngineNodes,
		Color: utils.StrColor(game.EngineColor),
	}
	engine.setUp()

	playees := []playee{
		player,
		engine,
	}
	//black to move - engine goes first
	if utils.StrColor(game.ColorTurn) == chess.Black {
		playees[0], playees[1] = playees[1], playees[0]
	}
	startGame(playees,player,engine)

}

func startGame(playees []playee,player Player,engine Engine) {
	fmt.Println(Game.Position().Board().Draw())
	
    for Game.Outcome() == chess.NoOutcome {
		for _,playee := range playees {
			exit,save := playee.getMoveAndMove()

			if !exit {
				fmt.Println(Game.Position().Board().Draw())
			}

			if exit || Game.Outcome() != chess.NoOutcome  {
				fmt.Println("Game Status: ",Game.Outcome(),Game.Method())
				if save {
					_,exists := gameDAO.GetByName(_gameName)
					err := saveGame(player,engine,exists)
					if err != nil {
						fmt.Println("Error at saving game:",err.Error())
						os.Exit(1)
					}
					fmt.Printf("Game %v Saved", _gameName)
				}
				os.Exit(0)
			}
		}

    }
}

func saveGame(player Player, engine Engine,update bool) error {
	game := models.Game{
		Color: utils.ColorStr(player.Color),
		GameName: _gameName,
		EngineColor: utils.ColorStr(engine.Color),
		ColorTurn: utils.ColorStr(Game.Position().Turn()),
		Engine: engine.Path,
		EngineDepth: engine.Depth,
		EngineNodes: engine.Nodes,
		Outcome: Game.Outcome().String(),
		Method: Game.Method().String(),
		FEN: Game.FEN(),
		PGN: Game.String(),
	}
	var err error
	if update {
		err = gameDAO.Update(game)
	} else {
		err = gameDAO.Insert(game)
	}
	return err
}

func GameContinuable(game models.Game) bool {
	if game.Outcome != chess.NoOutcome.String() {
		fmt.Printf("Game \"%v\" isn't continuable, Status: %v %v\n",game.GameName,game.Outcome,game.Method) 
		URL,err := lichessAnalysisURL(game.PGN)
		fmt.Print("Analyze on lichess: ")
		if err != nil { 
			fmt.Println("Can't get link,",err.Error())
		} else {
			fmt.Println(URL)
		}
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