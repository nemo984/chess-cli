/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/nemo984/chess-cli/data"
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
		_,ok := data.GetGame(args[0])
		if !ok {
			return errors.New("game with this name doesn't exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resign called")
		//delete game
	},
}

func init() {
	rootCmd.AddCommand(resignCmd)

}
