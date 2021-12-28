package main

import (
	"log"

	"github.com/nemo984/chess-cli/cmd"
	"github.com/nemo984/chess-cli/data"
)

func main() {
	if err := data.OpenDatabase(); err != nil {
		log.Fatal("Cannot open database", err.Error())
	}
	cmd.Execute()
}
