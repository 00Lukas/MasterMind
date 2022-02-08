package game

import (
	"github.com/00Lukas/MasterMind/board"
	"math/rand"
	"time"
)

var (
	//Right colors
	answer string
)

func randomColors() {
	b := board.GetAvailableColors()
	s := time.Now().UnixNano()
	rand.Seed(s)
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	for i := 0; i < 4; i++ {
		answer += b[i+1]
	}
}

func countBlack(s string) int {
	var count int
	for i := 0; i < board.GetColors(); i++ {
		if string(answer[i]) == string(s[i]) {
			count++
		}
	}
	return count
}

func countWhite(s string) int {
	var count int
	for i := 0; i < board.GetColors(); i++ {
		for j := 0; j < board.GetColors(); j++ {
			if i != j {
				if string(answer[i]) == string(s[j]) {
					count++
					break
				}
			}
		}
	}
	return count
}

func StartGame() {
	board.ClearBoard()
	randomColors()
}

func CalculateHint(s string) string {
	var temp string
	black := countBlack(s)
	white := countWhite(s)
	for i := 0; i < black; i++ {
		temp += "b"
	}
	for i := 0; i < white; i++ {
		temp += "w"
	}
	return temp
}

func CheckIfWin(s string) bool {
	if countBlack(s) == board.GetColors() {
		return true
	}
	return false
}

func GetAnswer() string {
	return answer
}
