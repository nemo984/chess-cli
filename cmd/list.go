/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	c "github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/utils"
	"github.com/notnil/chess"

	"github.com/spf13/cobra"
)

var (
	all bool
	e   bool

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all your games (default will only show on-going games)",
		Run: func(cmd *cobra.Command, args []string) {
			if e {
				displayGamesEngine(all)
			} else {
				displayGames(all)
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "show all games with board")
	listCmd.Flags().BoolVarP(&e, "engine", "e", false, "show games' engine configuration")

}

func displayGamesEngine(listAll bool) {
	games, err := gameDAO.GetAll()
	if err != nil {
		fmt.Println("No games found")
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Game Name", "Engine Path", "Depth", "Nodes"})
	for _, game := range games {
		method := game.Method
		if c := strings.Compare(method, "NoMethod"); c != 0 && !listAll {
			continue
		}
		t.AppendRow(table.Row{game.GameName, game.Engine, game.EngineDepth, game.EngineNodes})
		t.AppendSeparator()
	}
	t.Render()
}

func displayGames(listAll bool) {
	games, err := gameDAO.GetAll()
	if err != nil {
		fmt.Println("No games found")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Last Played", "Name", "Color", "Turn", "Result", "Method", "Board"})
	for _, game := range games {
		method := game.Method
		if c := strings.Compare(method, "NoMethod"); c == 0 {
			method = "Undecided"
		} else {
			if !listAll {
				continue
			}
		}

		fen, err := chess.FEN(game.FEN)
		if err != nil {
			log.Fatal(err.Error())
		}
		g := chess.NewGame(fen)
		board := c.Board{g.Position().Board()}
		lastPlayed := utils.GetLastPlayed(game.Updated)
		//, filepath.Base(game.Engine)
		t.AppendRow(table.Row{lastPlayed, game.GameName, game.Color, game.ColorTurn, game.Outcome, method, board.DrawP(utils.StrColor(game.Color))})
		t.AppendSeparator()
	}

	t.Render()
}
