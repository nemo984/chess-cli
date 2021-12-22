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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "chess-cli delete [game-names...]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one game name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			_,ok := gameDAO.GetByName(name)
			if !ok {
				fmt.Printf("Game \"%v\" doesn't exist.\n",name)
				continue
			}
			err := gameDAO.DeleteByName(name)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			fmt.Printf("Game \"%v\" is deleted.\n",name)
		}
	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)
}
