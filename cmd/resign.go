/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/nemo984/chess-cli/chess"
	"github.com/spf13/cobra"
)

// resignCmd represents the resign command
var resignCmd = &cobra.Command{
	Use:   "resign",
	Short: "chess-cli resign [game-names...]",
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
			if continuable := chess.GameContinuable(game); !continuable {
				fmt.Printf("Game \"%v\" is already over.\n", name)
				continue
			}

			game.Resign()
			err := gameDAO.Update(game)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("You resigned on Game \"%v\" Status: %v, %v\n", game.GameName, game.Outcome, game.Method)
		}
	},
}

func init() {
	rootCmd.AddCommand(resignCmd)

}
