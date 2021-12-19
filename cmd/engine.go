/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// engineCmd represents the engine command
var engineCmd = &cobra.Command{
	Use:   "engine",
	Short: "Play against an engine",
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
		if game == "" {
			chess.NewGame(engine)
		} else {
			chess.ContinueGame(game)
		}
	},
}

var (
	engine chess.Engine
	game string
)

func init() {

	playCmd.AddCommand(engineCmd)
	engineCmd.Flags().StringVarP(&engine.Path, "path", "p","", "Set the path where engine is stored")
	viper.BindPFlag("path", engineCmd.Flags().Lookup("path"))
	// engineCmd.MarkFlagRequired("path")

	engineCmd.Flags().IntVarP(&engine.Depth,"depth","d",21,"Set the engine depth/to search x piles only")
	viper.BindPFlag("depth", engineCmd.Flags().Lookup("depth"))
	
	engineCmd.Flags().IntVarP(&engine.Nodes, "nodes", "n",862438, "Set the engine to search x nodes only")
	viper.BindPFlag("nodes", engineCmd.Flags().Lookup("nodes"))
	
	engineCmd.Flags().StringVar(&game, "game","", "continue an existing game with x name")
	viper.BindPFlag("game", engineCmd.Flags().Lookup("game"))
}
