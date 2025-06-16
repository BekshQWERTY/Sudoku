package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 9 || !IsValidInput(args) {
		PrintError()
		return
	}

	var board [9][9]rune
	for i, row := range args {
		for j, c := range row {
			board[i][j] = c
		}
	}

	if !IsInitialBoardValid(board) {
		PrintError()
		return
	}

	if Solve(&board) {
		PrintBoard(board)
	} else {
		PrintError()
	}
}

func IsValid(board [9][9]rune, row int, col int, num rune) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	stertCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][stertCol+j] == num {
				return false
			}
		}
	}
	return true
}

func Solve(board *[9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := '1'; num <= '9'; num++ {
					if IsValid(*board, row, col, num) {
						board[row][col] = num
						if Solve(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func IsValidInput(args []string) bool {
	if len(args) != 9 {
		return false
	}

	for _, v := range args {
		if len(v) != 9 {
			return false
		}
		for _, f := range v {
			switch {
			case f == '.':
				continue
			case f >= '1' && f <= '9':
				continue
			default:
				return false
			}
		}
	}
	return true
}

func IsInitialBoardValid(board [9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				board[i][j] = '.'
				if !IsValid(board, i, j, num) {
					return false
				}
				board[i][j] = num
			}
		}
	}
	return true
}

func PrintBoard(board [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(board[i][j])

			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func PrintError() {
	args := os.Args[1:]
	if len(args) != 9 {
		message := "Error\n"
		for _, ch := range message {
			z01.PrintRune(ch)
		}
		return
	}

	message := "Error\n"
	for _, ch := range message {
		z01.PrintRune(ch)
	}
}
