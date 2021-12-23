package cmd

import (
	"errors"
	"fmt"

	"github.com/nemo984/chess-cli/chess/lichess"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [game-names...]",
	Short: "get lichess analyze urls",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one game name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			game, ok := gameDAO.GetByName(name)
			if !ok {
				fmt.Printf("Game \"%v\" doesn't exist.\n", name)
				continue
			}
			url, err := lichess.AnalysisURL(game.PGN)
			fmt.Printf("Analyze Game \"%v\" on lichess: ", name)
			if err != nil {
				fmt.Println("Can't get link,", err.Error())
			} else {
				fmt.Println(url)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
