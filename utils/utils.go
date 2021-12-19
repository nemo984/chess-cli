package utils

import "github.com/notnil/chess"

func ColorStr(color chess.Color) string {
	if color == chess.White {
		return "White"
	}
	return "Black"
}

func StrColor(color string) chess.Color {
	if color == "White" {
		return chess.White
	}
	return chess.Black
}
