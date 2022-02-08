package board

const (
	row        = 4
	col        = 10
	emptyField = "0"
)

var (
	board [col][row]string
	//Available colors for the game
	availableColors = []string{"w", "b", "r", "g", "p", "y"}
)

func colorsDuplicate(s string) bool {
	x := make(map[string]int, 0)
	for i := 0; i < row; i++ {
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
	if count == row {
		return true
	}
	return false
}

func ClearBoard() {
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			board[i][j] = emptyField
		}
	}
}

func AddColorsToBoard(n int, s string) bool {
	if len(s) != row {
		return false
	}
	if !checkColorsExist(s) || colorsDuplicate(s) {
		return false
	}
	for i := 0; i < row; i++ {
		board[n][i] = string(s[i])
	}
	return true
}

func GetBoard() [col][row]string {
	return board
}

func GetAvailableColors() []string {
	return availableColors
}

func GetColors() int {
	return row
}
