/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/nemo984/chess-cli/chess"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a chess game",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		startNewGame()
		fmt.Println(Level);
		chess.StartGame(Level)
	},

}

type promptContent struct {
	errorMsg string
	label string
}



var Level int;

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().IntVar(&Level, "level", 1, "Computer difficulty level")
	viper.BindPFlag("level", rootCmd.Flags().Lookup("level"))
}

// func promptGetInput(pc promptContent) string {

// }

// func promptGetSelect(pc promptContent) string {

// }

func startNewGame() {
	
}