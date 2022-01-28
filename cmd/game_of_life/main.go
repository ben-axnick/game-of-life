package main

import (
	"fmt"
	"time"
)

func Board(dy, dx int) [][]uint8 {
	board := make([][]uint8, dx)
	cells := make([]uint8, dx*dy)

	for y := range board {
		board[y], cells = cells[:dy], cells[dy:]
	}

	return board
}

func SeedBoard(board [][]uint8) {
	board[1][1] = 1
	board[1][3] = 1
	board[3][1] = 1
	board[3][2] = 1
	board[3][3] = 1
}

func DrawBoard(board [][]uint8) {
	for y := range board {
		row := make([]rune, len(board[y]))

		for x := range board[y] {
			switch board[y][x] {
			case 1:
				row[x] = '😺'
			case 0:
				row[x] = '🖤'
			}
		}

		fmt.Println(string(row))
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func CountLiveNeighbours(board [][]uint8, y, x int) int {
	x1 := Max(x-1, 0)
	x2 := Min(x+1, len(board[y])-1)
	y1 := Max(y-1, 0)
	y2 := Min(y+1, len(board)-1)

	count := 0

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if board[y][x] == 1 {
				count++
			}
		}
	}

	if board[y][x] == 1 {
		count--
	}

	return count
}

func NextCellState(board [][]uint8, y, x int) uint8 {
	isLive := board[y][x] == 1
	neighbours := CountLiveNeighbours(board, y, x)

	if isLive && neighbours == 2 || neighbours == 3 {
		return 1
	} else if !isLive && neighbours == 3 {
		return 1
	}

	return 0
}

func UpdateBoard(board [][]uint8) {
	next_board := Board(len(board), len(board[0]))

	for y := range board {
		for x := range board[y] {
			next_board[y][x] = NextCellState(board, y, x)
		}
	}

	copy(board, next_board)
}

func main() {
	// Initialise a zeroed game board
	board := Board(5, 5)
	tick := 0

	// Activate some pixels, probably start with something fixed
	SeedBoard(board)

	// Game loop to calculate next game state, display the board

	for {
		tick++
		fmt.Printf("Tick %d\n", tick)
		DrawBoard(board)
		UpdateBoard(board)
		time.Sleep(1 * time.Second)
	}

	// That's the basics out of the way. Other thoughts:

	// Unit test the calculation of next game state if we haven't already
	// Use Ebiten to draw the game board
	// Add play/pause toggle
	// Add clear/reset toggle
	// Add a speed slider
	// Toggle cells with mouse click
	// Implement some additional seeds, have a picker to load the chosen seed
	// Add a way to save a seed from the current game state
	// Deploy game to a website using WASM
}
