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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// puzzleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// puzzleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
