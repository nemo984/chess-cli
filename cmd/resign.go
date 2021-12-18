/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// resignCmd represents the resign command
var resignCmd = &cobra.Command{
	Use:   "resign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a color argument")
		}
		// if exist := data.GameExists(args[0]); !exist {
		// 	return errors.New("Game doesn't exist")
		// }
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resign called")
		//delete game
	},
}

func init() {
	rootCmd.AddCommand(resignCmd)

}
