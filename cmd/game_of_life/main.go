package main

import "fmt"

func Board(dx, dy int) [][]uint8 {
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

func main() {
	// Initialise a zeroed game board
	board := Board(5, 5)
	tick := 1

	// Activate some pixels, probably start with something fixed
	SeedBoard(board)

	// Display the board, probably just terminal based for starters
	fmt.Printf("Tick %d\n", tick)
	DrawBoard(board)

	// Function to calculate next board state from current board state
	// Game loop to calculate next game state, display the board

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
