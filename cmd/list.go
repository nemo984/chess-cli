/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/nemo984/chess-cli/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your games",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayGames()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func displayGames() {
	games,err := data.GetGames()
	if err != nil {
		fmt.Println("No games found")
	}
	for _,game := range games {
		fmt.Println(game)
	}
}