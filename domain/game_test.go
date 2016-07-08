package domain

import "testing"

func TestNewGame(t *testing.T) {

	game := NewGame()

	if game == nil {
		t.Fatal("Game is nil")
	}

	blocks := game.GetBlocks()

	if blocks == nil {
		t.Fatal("Blocks are nil")
	}

	actualWidth := len(*blocks)
	if actualWidth != BoardWidth {
		t.Errorf("Expected width: %d received: %d", BoardWidth, actualWidth)
	}

}
