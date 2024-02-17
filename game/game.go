package game

import (
	"math/rand"
	"strings"

	"github.com/00Lukas/MasterMind/board"
)

var (
	//Right colors
	answer string
)

func randomColors() {
	b := board.GetAvailableColors()
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
			if i != j && string(answer[i]) == string(s[j]) {
				count++
				break
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
	var temp []string
	black := countBlack(s)
	white := countWhite(s)
	for i := 0; i < black; i++ {
		temp = append(temp, "b")
	}
	for i := 0; i < white; i++ {
		temp = append(temp, "w")
	}

	rand.Shuffle(len(temp), func(i, j int) {
		temp[i], temp[j] = temp[j], temp[i]
	})

	return strings.Join(temp,"")
}

func CheckIfWin(s string) bool {
	return countBlack(s) == board.GetColors() 
}

func GetAnswer() string {
	return answer
}
