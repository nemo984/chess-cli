/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a chess game",
	Run: func(cmd *cobra.Command, args []string) {
	},

}

func init() {
	rootCmd.AddCommand(playCmd)

}
