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
	Short: "chess-cli resign [game-name]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a game name argument")
		}
		game,ok := gameDAO.GetByName(args[0])
		if !ok {
			return fmt.Errorf("game \"%v\" doesn't exist",args[0])
		}
		if continuable := chess.GameContinuable(game); !continuable {
			return fmt.Errorf("game \"%v\" is already over",args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		game,_ := gameDAO.GetByName(args[0])
		game.Resign()
		err := gameDAO.Update(game)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("You resigned on Game \"%v\" Status: %v, %v", game.GameName, game.Outcome, game.Method)
	},
}

func init() {
	rootCmd.AddCommand(resignCmd)

}
