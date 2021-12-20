/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	c "github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/utils"
	"github.com/notnil/chess"

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


func displayGames() {
	games,err := gameDAO.GetAll()
	if err != nil {
		fmt.Println("No games found")
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
    t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Name", "Color", "Turn", "Engine","Result","Method","Board"})

	for _,game := range games {
		method := game.Method
		if c := strings.Compare(method,"NoMethod"); c == 0 {
			method = "On-going"
		}

		fen,err := chess.FEN(game.FEN)
		if err != nil {
			log.Fatal(err.Error())
		}
		g := chess.NewGame(fen)
		board := g.Position().Board()

		t.AppendRow(table.Row{game.GameName, game.Color, game.ColorTurn, filepath.Base(game.Engine), game.Outcome, method, c.DrawP(board, utils.StrColor(game.Color))})	
		t.AppendSeparator()
	}
	t.Render()
}