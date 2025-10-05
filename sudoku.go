package main

import (
	"fmt"
	"os"
)

func main() {
	const size = 9

	if len(os.Args) != 10 {
		fmt.Println("error")
		return
	}

	board := [9][9]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			ch := os.Args[i+1][j]
			if ch == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = int(ch - '0')
			}
		}
	}

	var solve func() bool
	solve = func() bool {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] == 0 {
					for num := 1; num <= 9; num++ {
						valid := true
						for k := 0; k < size; k++ {
							if board[i][k] == num || board[k][j] == num {
								valid = false
								break
							}
						}
						
						startRow := (i / 3) * 3
						startCol := (j / 3) * 3
						for r := 0; r < 3; r++ {
							for c := 0; c < 3; c++ {
								if board[startRow+r][startCol+c] == num {
									valid = false
								}
							}
						}

						if valid {
							board[i][j] = num
							if solve() {
								return true
							}
							board[i][j] = 0
						}
					}
					return false
				}
			}
		}
		return true
	}

	if solve() {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if j == size-1 {
					fmt.Printf("%d", board[i][j])
				} else {
					fmt.Printf("%d ", board[i][j])
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("error")
	}
}