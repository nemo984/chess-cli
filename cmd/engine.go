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

// engineCmd represents the engine command
var engineCmd = &cobra.Command{
	Use:   "engine",
	Short: "Play against an engine",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(engine)
		chess.StartGame(engine)
	},
}

var engine chess.Engine

func init() {
	playCmd.AddCommand(engineCmd)
	engineCmd.Flags().StringVarP(&engine.Path, "path", "p","", "Set the path where engine is stored")
	viper.BindPFlag("path", rootCmd.Flags().Lookup("path"))
	engineCmd.MarkFlagRequired("path")

	engineCmd.Flags().IntVarP(&engine.Depth,"depth","d",21,"Set the engine depth/to search x piles only")
	viper.BindPFlag("depth", rootCmd.Flags().Lookup("depth"))
	
	engineCmd.Flags().IntVarP(&engine.Nodes, "nodes", "n",862438, "Set the engine to search x nodes only")
	viper.BindPFlag("nodes", rootCmd.Flags().Lookup("nodes"))
	
}
