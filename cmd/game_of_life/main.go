package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// Initialise a zeroed game board
	// Activate some pixels, probably start with something fixed
	// Display the board, probably just terminal based for starters
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
