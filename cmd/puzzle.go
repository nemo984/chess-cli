package cmd

import (
	"fmt"

	"github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/data"
	"github.com/spf13/cobra"
)

// puzzleCmd represents the puzzle command
var puzzleCmd = &cobra.Command{
	Use:   "puzzle",
	Short: "Play a daily lichess puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		game := chess.NewChessGame(data.Game{}, "Puzzle")
		if err := game.StartPuzzle(); err != nil {
			fmt.Println("Error:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(puzzleCmd)
}
