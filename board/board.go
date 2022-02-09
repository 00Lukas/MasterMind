package board

const (
	colors     = 4
	col        = 10
	EmptyField = "0"
)

var (
	board [col]string
	//Available colors for the game
	availableColors = []string{"w", "b", "r", "g", "p", "y"}
)

func colorsDuplicate(s string) bool {
	x := make(map[string]int, 0)
	for i := 0; i < colors; i++ {
		x[string(s[i])]++
	}
	for i := range x {
		if x[i] > 1 {
			return true
		}
	}
	return false
}

func checkColorsExist(s string) bool {
	var count int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(availableColors); j++ {
			if string(s[i]) == availableColors[j] {
				count++
			}
		}
	}
	if count == colors {
		return true
	}
	return false
}

func ClearBoard() {
	for i := 0; i < col; i++ {
		for j := 0; j < colors; j++ {
			board[i] = EmptyField
		}
	}
}

func AddColorsToBoard(n int, s string) bool {
	if len(s) != colors {
		return false
	}
	if !checkColorsExist(s) || colorsDuplicate(s) {
		return false
	}
	for i := 0; i < colors; i++ {
		board[n] = s
	}
	return true
}

func GetBoard() [col]string {
	return board
}

func GetAvailableColors() []string {
	return availableColors
}

func GetColors() int {
	return colors
}
