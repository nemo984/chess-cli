package chess

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nemo984/chess-cli/chess/lichess"
	"github.com/nemo984/chess-cli/data"
	"github.com/nemo984/chess-cli/models"
	"github.com/nemo984/chess-cli/utils"
	"github.com/notnil/chess"
)

type ChessGame struct {
	Game *chess.Game
	DB   data.Game
	Name string
}

func NewChessGame(db data.Game, name string) *ChessGame {
	return &ChessGame{Game: chess.NewGame(), DB: db, Name: name}
}

type playee interface {
	getMoveAndMove(game *chess.Game) (exit bool, save bool, err error)
}

func (c *ChessGame) NewEngineGame(engine Engine, color string) error {
	player := Player{
		Color:       utils.StrColor(color),
		MoveOptions: EngineGameOptions,
		Out:         os.Stdout,
	}

	if err := engine.setUp(); err != nil {
		return err
	}
	engine.setColor(player.Color.Other())

	playees := setUpTurn(chess.White, player, engine)
	if err := c.startGame(playees, player, engine); err != nil {
		return err
	}

	return nil
}

func (c *ChessGame) ContinueGame() error {
	game, ok := c.DB.GetByName(c.Name)
	if !ok {
		return fmt.Errorf("game \"%v\" doesn't exist", c.Name)
	}
	if ok := GameContinuable(game); !ok {
		return fmt.Errorf("game \"%v\" isn't continuable", c.Name)
	}

	fen, err := chess.FEN(game.FEN)
	if err != nil {
		return err
	}
	c.Game = chess.NewGame(fen)
	
	player := Player{
		Color:       utils.StrColor(game.Color),
		MoveOptions: EngineGameOptions,
		Out:         os.Stdout,
	}
	engine := Engine{
		Path:  game.Engine,
		Depth: game.EngineDepth,
		Color: player.Color.Other(),
	}
	if err = engine.setUp(); err != nil {
		return err
	}

	turns := setUpTurn(utils.StrColor(game.ColorTurn), player, engine)
	if err := c.startGame(turns, player, engine); err != nil {
		return err
	}
	return nil
}

func setUpTurn(colorTurn chess.Color, player Player, engine Engine) []playee {
	playees := []playee{
		player,
		engine,
	}
	if engine.Color == colorTurn {
		playees[0], playees[1] = playees[1], playees[0]
	}
	return playees
}

func (c *ChessGame) startGame(playees []playee, player Player, engine Engine) error {
	board := Board{c.Game.Position().Board()}
	fmt.Println(board.DrawP(player.Color))

	for c.Game.Outcome() == chess.NoOutcome {
		for _, playee := range playees {
			exit, save, err := playee.getMoveAndMove(c.Game)
			if err != nil {
				return err
			}
			if !exit {
				board := Board{c.Game.Position().Board()}
				fmt.Println(board.DrawP(player.Color))
			}

			if exit || c.Game.Outcome() != chess.NoOutcome {
				var method string
				if c.Game.Method() == chess.NoMethod {
					method = "Undecided"
				} else {
					method = c.Game.Method().String()
				}
				fmt.Println("Game Status:", c.Game.Outcome(), method)

				if save {
					_, exists := c.DB.GetByName(c.Name)
					if err := c.saveGame(player, engine, exists); err != nil {
						return fmt.Errorf("error at saving game: %w", err)
					}

					fmt.Printf("Game \"%v\" Saved", c.Name)
				}
				os.Exit(0)
			}
		}

	}
	return nil
}

func (c *ChessGame) saveGame(player Player, engine Engine, update bool) error {
	game := models.Game{
		Color:       utils.ColorStr(player.Color),
		GameName:    c.Name,
		EngineColor: utils.ColorStr(engine.Color),
		ColorTurn:   utils.ColorStr(c.Game.Position().Turn()),
		Engine:      engine.Path,
		EngineDepth: engine.Depth,
		Outcome:     c.Game.Outcome().String(),
		Method:      c.Game.Method().String(),
		FEN:         c.Game.FEN(),
		PGN:         c.Game.String(),
		Updated:     time.Now().Format(time.RFC3339),
	}
	var err error
	if update {
		err = c.DB.Update(game)
	} else {
		err = c.DB.Insert(&game)
	}
	return err
}

func (c *ChessGame) StartPuzzle() error {
	puzzle, err := lichess.GetPuzzle()
	if err != nil {
		return err
	}

	new, err := chess.PGN(strings.NewReader(puzzle.Game.PGN))
	if err != nil {
		return err
	}
	solution := puzzle.Puzzle.Solution
	rating := puzzle.Puzzle.Rating
	c.Game = chess.NewGame(new)
	player := Player{
		Color:       c.Game.Position().Turn(),
		MoveOptions: PuzzleGameOptions,
		Out:         os.Stdout,
	}

	board := Board{c.Game.Position().Board()}

	g := chess.AlgebraicNotation{}
	uci := chess.UCINotation{}

	fmt.Println("Daily Puzzle started, Rating:", rating)
	if player.Color == chess.White {
		fmt.Println("White to Move")
	} else {
		fmt.Println("Black to Move")
	}
	fmt.Println(board.DrawP(player.Color))
	next := false
	for i := 0; i < len(solution); i++ {
		for !next {
			move, err := player.getMove()
			if err != nil {
				return err
			}
			switch move {
			case "h":
				fmt.Printf("Hint: %v piece\n", solution[i][:2])

			case "s":
				fmt.Println("Solution/Remaining Moves:", solution[i:])

			case "q":
				os.Exit(0)

			default:
				moveSol, err := uci.Decode(c.Game.Position(), solution[i])
				if err != nil {
					return errors.New("can't decode lichess move solution")
				}
				input, err := g.Decode(c.Game.Position(), move)
				if err != nil {
					fmt.Println("Not valid move")
				} else {
					if moveSol.S1() == input.S1() && moveSol.S2() == input.S2() {
						if err := c.Game.Move(input); err != nil {
							return err
						}
						board = Board{c.Game.Position().Board()}
						fmt.Println(board.DrawP(player.Color))
						next = true
					} else {
						fmt.Println("Incorrect Move, Try Again")
					}
				}

			}
		}

		if i+1 >= len(solution) {
			fmt.Println("Daily puzzle solved.")
			os.Exit(0)
		}

		fmt.Println("Correct Move, Continue:", solution[i+1])
		moveSol, err := uci.Decode(c.Game.Position(), solution[i+1])
		if err != nil {
			fmt.Println("Solution/Remaining Moves:", solution[i+1:])
			return fmt.Errorf("can't decode lichess next move\n%v", err)
		}
		err = c.Game.Move(moveSol)
		if err != nil {
			fmt.Println("Solution/Remaining Moves:", solution[i+1:])
			return fmt.Errorf("lichess next move is invalid \n%v", err)
		}
		board = Board{c.Game.Position().Board()}
		fmt.Println(board.DrawP(player.Color))
		i++
		next = false
	}

	return nil
}

func GameContinuable(game models.Game) bool {
	if game.Outcome != chess.NoOutcome.String() {
		fmt.Printf("Game \"%v\" isn't continuable, Status: %v %v\n", game.GameName, game.Outcome, game.Method)
		URL, err := lichess.AnalysisURL(game.PGN)
		fmt.Print("Analyze on lichess: ")
		if err != nil {
			fmt.Println("Can't get link,", err.Error())
		} else {
			fmt.Println(URL)
		}
		return false
	}
	return true
}

type Board struct {
	*chess.Board
}

//Draw Board based on color perspective
func (b *Board) DrawP(color chess.Color) string {
	s := "\n A B C D E F G H\n"
	rows := []int{7, 6, 5, 4, 3, 2, 1, 0}
	if color == chess.Black {
		b.Flip(chess.UpDown)
		rows = []int{0, 1, 2, 3, 4, 5, 6, 7}
	}
	for _, r := range rows {
		s += chess.Rank(r).String()
		for f := 0; f < 8; f++ {
			p := b.Piece(getSquare(chess.File(f), chess.Rank(r)))
			if p == chess.NoPiece {
				s += "-"
			} else {
				s += p.String()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

func getSquare(f chess.File, r chess.Rank) chess.Square {
	return chess.Square((int(r) * 8) + int(f))
}
