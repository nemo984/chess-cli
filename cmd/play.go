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
		RunE: func(cmd *cobra.Command, args []string) error {
			err := chess.ContinueGame(game)
			return err
		},
	}
)

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVar(&game, "game", "", "continue an existing game with x name")
	playCmd.MarkFlagRequired("game")
}
