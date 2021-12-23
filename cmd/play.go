package cmd

import (
	"github.com/nemo984/chess-cli/chess"
	"github.com/spf13/cobra"
)

var (
	game string

	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Play/Continue a chess game",
		Run: func(cmd *cobra.Command, args []string) {
			chess.ContinueGame(game)
		},
	}
)

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVar(&game, "game", "", "continue an existing game with x name")
	playCmd.MarkFlagRequired("game")
}
