/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/nemo984/chess-cli/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your games",
	Run: func(cmd *cobra.Command, args []string) {
		displayGames()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

var gameDAO data.Game

func displayGames() {
	games,err := gameDAO.GetAll()
	if err != nil {
		fmt.Println("No games found")
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Created\tUpdated\tName\tColor\tTurn\tStatus\tEngine\tDepth\tNodes")
	for _,game := range games {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",game.Created,game.Updated,game.GameName,game.Color,game.ColorTurn,game.Outcome,game.Engine,game.EngineDepth,game.EngineNodes)
	}
	w.Flush()
}