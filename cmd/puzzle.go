/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/nemo984/chess-cli/chess"
	"github.com/spf13/cobra"
)

// puzzleCmd represents the puzzle command
var puzzleCmd = &cobra.Command{
	Use:   "puzzle",
	Short: "Play a daily lichess puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("puzzle called")
		chess.StartPuzzle()
	},
}

func init() {
	rootCmd.AddCommand(puzzleCmd)
}
