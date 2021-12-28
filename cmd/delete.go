package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [game-names...]",
	Short: "Delete your chess games from database",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one game name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			_, ok := gameDAO.GetByName(name)
			if !ok {
				fmt.Printf("Game \"%v\" doesn't exist.\n", name)
				continue
			}
			if err := gameDAO.DeleteByName(name); err != nil {
				fmt.Printf("Error at deleting game \"%v\": %s\n", name, err.Error())
				continue
			}
			fmt.Printf("Game \"%v\" is deleted.\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
