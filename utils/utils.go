package utils

import (
	"fmt"
	"math/rand"
	"strings"
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
	if strings.ToLower(color) == "white" {
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

func GetLastPlayed(timeStr string) string {
	updated, _ := time.Parse(time.RFC3339,timeStr)
	t := time.Since(updated)
	if t.Hours()/ (24*30) > 1 {
		return fmt.Sprintf("%v months ago", int(t.Hours() / (24*30)))
	}
	if t.Hours()/24 > 1 {
		return fmt.Sprintf("%v days ago", int(t.Hours()/24))
	}
	if t.Hours() > 1 {
		return fmt.Sprintf("%v hours ago", int(t.Hours()))
	} 
	if t.Minutes() > 1 {
		return fmt.Sprintf("%v minutes ago", int(t.Minutes()))
	}
	return fmt.Sprintf("%v seconds ago", int(t.Seconds()))
}