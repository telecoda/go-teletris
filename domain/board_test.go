package domain

import "testing"

func TestNewBoard(t *testing.T) {

	board := NewBoard()

	if board.cells == nil {
		t.Fatal("Board cells are nil")
	}

	actualWidth := len(board.cells)
	if actualWidth != BoardWidth {
		t.Errorf("Expected width: %d received: %d", BoardWidth, actualWidth)
	}

	actualHeight := len(board.cells[0])
	if actualHeight != BoardHeight {
		t.Errorf("Expected height: %d received: %d", BoardHeight, actualHeight)
	}

}
