package chess

import (
	"fmt"
	"math/rand"

	"github.com/notnil/chess"
)

type Player struct {
	Color chess.Color
}

func (p Player) getMoveAndMove(options string) (exit bool, save bool) {
	var input string
	for {
		fmt.Print("Your Move, Enter (?) for options: ")
		fmt.Scanln(&input)
		switch input {
		case "?":
			fmt.Println(options)
		
		case "v":
			fmt.Println("Valid Moves:",Game.ValidMoves())

		case "r":
			moves := Game.ValidMoves()
			move := rand.Intn(len(moves))
			Game.Move(moves[move])
			fmt.Println("Random move!:", moves[move])
			return false, true

		case "q":
			return true, true

		case "q!":
			return true, false

		case "resign":
			Game.Resign(p.Color)
			return true, true

		default:
			if err := Game.MoveStr(input); err != nil {
				fmt.Println("Invalid Move, Try Again")
			} else {
				fmt.Println("You played: ", input)
				return false, true
			}
		}

	}
}

func (p Player) getMove() string {
	var input string
	fmt.Print("Your Move, Enter (?) for options: ")
	fmt.Scanln(&input)
	return input
}

var (
	EngineGameOptions = `To make a move, Enter an Algebratic Notation, Examples: e2, e5, O-O (short castling), e8=Q (promotion)
	or Long Algebraic Notation, Examples: Rd1xd8+, Ng8f6.
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
