package cmd

import (
	"os"

	"github.com/nemo984/chess-cli/data"
	"github.com/spf13/cobra"
)

var gameDAO data.Game

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chess-cli",
	Short: "Play chess against computer inside your terminal",
	Long: `Chess-cli is a CLI to play chess against an engine of your choice with the ability to specify depth and nodes
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
