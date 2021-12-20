package utils

import (
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}