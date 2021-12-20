/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/nemo984/chess-cli/chess"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	game string

	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Start a chess game",
		Run: func(cmd *cobra.Command, args []string) {
			chess.ContinueGame(game)
		},
	
	}
)


func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVar(&game, "game","", "continue an existing game with x name")
	viper.BindPFlag("game", engineCmd.Flags().Lookup("game"))

}
