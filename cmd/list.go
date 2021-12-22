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

var (
	all bool

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all your games (default will only show on-going games)",
		Run: func(cmd *cobra.Command, args []string) {
			displayGames(all)
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "list all games")

}


func displayGames(listAll bool) {
	games,err := gameDAO.GetAll()
	if err != nil {
		fmt.Println("No games found")
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
    t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Last Played","Name", "Color", "Turn", "Engine","Result","Method","Board"})

	for _,game := range games {
		method := game.Method
		if c := strings.Compare(method,"NoMethod"); c == 0 {
			method = "Undecided"
		} else {
			if !listAll {
				continue
			}
		}

		fen,err := chess.FEN(game.FEN)
		if err != nil {
			log.Fatal(err.Error())
		}
		g := chess.NewGame(fen)
		board := c.Board{g.Position().Board()}
		lastPlayed := utils.GetLastPlayed(game.Updated)

		t.AppendRow(table.Row{lastPlayed, game.GameName, game.Color, game.ColorTurn, filepath.Base(game.Engine), game.Outcome, method, board.DrawP(utils.StrColor(game.Color))})	
		t.AppendSeparator()
	}
	t.Render()
}