package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nemo984/chess-cli/chess"
	"github.com/nemo984/chess-cli/data"
	"github.com/nemo984/chess-cli/utils"
	"github.com/spf13/cobra"
)

var (
	engine chess.Engine
	name   string
	color  string

	engineCmd = &cobra.Command{
		Use:   "engine",
		Short: "Start a game against an engine",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := data.CreateTable(); err != nil {
				return err
			}
			if name != "" {
				_, ok := gameDAO.GetByName(name)
				if ok {
					return fmt.Errorf("game \"%v\" already exists", name)
				}
			} else {
				name = utils.RandStringRunes(5)
			}

			if color == "" {
				rand.Seed(time.Now().UnixNano())
				color = []string{"white", "black"}[rand.Intn(2)]
			}
			if err := chess.NewGame(engine, name, color); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	playCmd.AddCommand(engineCmd)

	engineCmd.Flags().StringVarP(&engine.Path, "path", "p", "", "Set the UCI chess engine path (required)")
	engineCmd.MarkFlagRequired("path")
	engineCmd.Flags().IntVarP(&engine.Depth, "depth", "d", 3, "Set the engine depth to search x piles only")

	engineCmd.Flags().StringVar(&name, "name", "", "Set the name of the game (default random)")
	engineCmd.Flags().StringVar(&color, "color", "", "choose your color: white/black (default random)")

}
