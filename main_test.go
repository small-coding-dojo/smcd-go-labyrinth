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
