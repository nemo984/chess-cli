/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/nemo984/chess-cli/cmd"
	"github.com/nemo984/chess-cli/data"
)

func main() {
	data.OpenDatabase();
	cmd.Execute()
}
