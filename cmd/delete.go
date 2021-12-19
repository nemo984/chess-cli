/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "chess-cli delete [game-name]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a game name argument")
		}
		_,ok := gameDAO.GetByName(args[0])
		if !ok {
			return fmt.Errorf("game \"%v\" doesn't exist",args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := gameDAO.DeleteByName(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		fmt.Printf("Game \"%v\" is deleted",args[0])
	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)
}
