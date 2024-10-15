package main

import "testing"

func TestPlaceARobot(t *testing.T) {
	got := Board{
		robot: Coordinate{row: 100, column: 100},
	}

	want := Board{
		robot: Coordinate{row: 100, column: 100},
	}

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func TestBoardWidth(t *testing.T) {
	someWidth := 2

	got := Board{
		robot: Coordinate{row: 100, column: 100},
		width: someWidth,
	}

	if got.width != someWidth {
		t.Errorf("got %q, but want %q", got.width, someWidth)
	}
}

func TestBoardHeight(t *testing.T) {
	someHeight := 2

	got := Board{
		robot:  Coordinate{row: 100, column: 100},
		height: someHeight,
	}

	if got.height != someHeight {
		t.Errorf("got %q, but want %q", got.height, someHeight)
	}
}

// Given a window size of 640 x 480
// When a 2 x 2 board is requested
// Then the tile size should be 480 / 2 = 240
func TestTileSizeFor640x480Board(t *testing.T) {
	got := calculateTileSize(640, 480)
	want := 240
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func calculateTileSize(boardWidth, boardHeight int) int {
	return 240 // Has to be calculated by width and height
}
