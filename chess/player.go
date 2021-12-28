package chess

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/notnil/chess"
)

type Player struct {
	Color       chess.Color
	MoveOptions string
	Out         io.Writer
}

var (
	EngineGameOptions = `To make a move, Enter an Algebratic Notation, Examples: e2, e5, O-O (short castling), e8=Q (promotion)
	To see valid moves, Enter (v)
	To resign, Enter (resign)
	To make a random move, Enter (r)
	To quit and save the game, Enter (q)
	To quit without saving, Enter (q!)
	`
	PuzzleGameOptions = `To make a move, Enter an Algebratic Notation, Examples: e2, e5, O-O (short castling), e8=Q (promotion)
	To see a hint, Enter (h)
	To see the solution, Enter (s)
	To quit, Enter (q)
	`
)

func (p Player) getMoveAndMove(game *chess.Game) (exit bool, save bool, err error) {
	for {
		move, err := p.getMove()
		if err != nil {
			return false, false, err
		}
		switch move {
		case "v":
			fmt.Fprintln(p.Out, "Valid Moves:", game.ValidMoves())

		case "r":
			moves := game.ValidMoves()
			move := rand.Intn(len(moves))
			if err := game.Move(moves[move]); err != nil {
				return false, false, err
			}
			fmt.Fprintln(p.Out, "Random move!:", moves[move])
			return false, true, nil

		case "q":
			return true, true, nil

		case "q!":
			return true, false, nil

		case "resign":
			game.Resign(p.Color)
			return true, true, nil

		default:
			if err := game.MoveStr(move); err != nil {
				fmt.Fprintln(p.Out, "Invalid Move, Try Again")
			} else {
				fmt.Fprintln(p.Out, "You played: ", move)
				return false, true, nil
			}
		}

	}
}

func (p Player) getMove() (move string, err error) {
	var input string
	for {
		fmt.Print("Your Move, Enter (?) for options: ")
		if _, err := fmt.Scanln(&input); err != nil {
			return "", err
		}
		switch input {
		case "?":
			fmt.Fprintln(p.Out, p.MoveOptions)
			continue
		}
		break
	}
	return input, nil

}
