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
	maxRound = 10
)

var (
	hint  string
	n     int
	quess string
)

func AskForColors(n int) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Chose colors: ")
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

func printAfterEnd() {
	fmt.Println("Right colors: ", game.GetAnswer())
	fmt.Println("Your colors: ", board.GetBoard())
}

func newGame() bool {
	game.StartGame()
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
				fmt.Println("Hint: ", hint)
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
