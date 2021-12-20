/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/data"
	"github.com/nemo984/chess-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var (
	engine chess.Engine
	name string
	color string

	engineCmd = &cobra.Command{
		Use:   "engine",
		Short: "Play against an engine",
		Run: func(cmd *cobra.Command, args []string) {
			data.CreateTable()
			if game == "" {
				game = utils.RandStringRunes(5)
			}
			chess.NewGame(engine,game)
		},
	}
)

func init() {
	playCmd.AddCommand(engineCmd)

	engineCmd.Flags().StringVarP(&engine.Path, "path", "p","", "Set the path where engine is stored")
	engineCmd.MarkFlagRequired("path")
	engineCmd.Flags().IntVarP(&engine.Depth,"depth","d",21,"Set the engine depth/to search x piles only")
	engineCmd.Flags().IntVarP(&engine.Nodes, "nodes", "n",862438, "Set the engine to search x nodes only")

	engineCmd.Flags().StringVar(&name, "name", "", "Set the name of the game (default random)")
	engineCmd.Flags().StringVar(&color,"color","","choose color (default random)")

	viper.BindPFlag("path", engineCmd.Flags().Lookup("path"))
	viper.BindPFlag("depth", engineCmd.Flags().Lookup("depth"))
	viper.BindPFlag("nodes", engineCmd.Flags().Lookup("nodes"))


}
