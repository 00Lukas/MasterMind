package main

import (
	"bufio"
	"fmt"
	"github.com/00Lukas/MasterMind/board"
	"github.com/00Lukas/MasterMind/game"
	"os"
	"strings"
)

const (
	maxRound   = 10
	emptyField = board.EmptyField
)

var (
	hint     string
	n        int
	quess    string
	allHints [maxRound]string
)

func AskForColors(n int) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[%d] Chose colors: ", n+1)
	col, _ := reader.ReadString('\n')
	col = strings.TrimSpace(col)
	quess = col
	return board.AddColorsToBoard(n, col)
}

func win() {
	fmt.Println("You win! Congratz")
	printAfterEnd()
}

func loose() {
	fmt.Println("You loose, sorry :(")
	printAfterEnd()

}

func clearHints() {
	for i := 0; i < maxRound; i++ {
		allHints[i] = emptyField
	}
}

func printAfterEnd() {
	fmt.Println("Right colors: ", game.GetAnswer())
	b := board.GetBoard()
	h := allHints
	for i := 0; i < maxRound; i++ {
		if b[i] == emptyField {
			break
		}
		fmt.Printf("[%d] Colors: %s\tHint: %s\n", i+1, b[i], h[i])
	}
}

func newGame() bool {
	game.StartGame()
	clearHints()
	n = 0
	for n = 0; n < maxRound; n++ {
		hint = ""
		for {
			if AskForColors(n) {
				break
			}
			fmt.Println("Wrong colors!")
			fmt.Println(board.GetAvailableColors())
		}

		if !game.CheckIfWin(quess) {
			if n < maxRound {
				hint = game.CalculateHint(quess)
				allHints[n] = hint
				fmt.Printf("[%d] Hint: %s\n", n+1, hint)
			}
		} else {
			return true
		}
	}
	return false
}

func main() {
	nGame := newGame()
	if !nGame && n == maxRound {
		loose()
	} else if nGame {
		win()
	}
}
