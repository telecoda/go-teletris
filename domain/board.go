package domain

type Block struct {
	X, Y   int
	Colour BlockColour
}

type Board struct {
	cells [][]*Block
}

func NewBoard() Board {

	b := Board{}
	// fill with blank cells
	b.cells = make([][]*Block, BoardWidth)

	for x := 0; x < BoardWidth; x++ {
		b.cells[x] = make([]*Block, BoardHeight)
		for y := 0; y < BoardHeight; y++ {
			b.cells[x][y] = &Block{X: x, Y: y, Colour: Empty}
		}
	}

	return b
}

func (b *Board) reset() {
	// add grey surrounding blocks
	y := 0
	for x := 0; x < BoardWidth; x++ {
		b.cells[x][y].Colour = Grey
	}
	y = BoardHeight - 1
	for x := 0; x < BoardWidth; x++ {
		b.cells[x][y].Colour = Grey
	}
	x := 0
	for y := 0; y < BoardHeight; y++ {
		b.cells[x][y].Colour = Grey
	}

	x = BoardWidth - 1
	for y := 0; y < BoardHeight; y++ {
		b.cells[x][y].Colour = Grey
	}
}

func (b *Board) canPlayerFitAt(player *Player, x, y int) bool {
	/*
	   Check if player's shape can move down one row
	   without colliding into any other blocks
	*/

	if player.shape == nil {
		return false
	}

	blocks := player.shape.GetBlocks()

	for _, block := range blocks {
		// check board is empty for all player blocks
		blockX := x + block.X
		blockY := y + block.Y

		if b.cells[blockX][blockY].Colour != Empty {
			return false
		}
	}

	return true
}

func (b *Board) addShapeToBoard(player *Player) {
	/*
		Shape has collided so add to the permanent board
	*/

	if player == nil {
		return
	}
	blocks := player.shape.GetBlocks()
	if blocks == nil {
		return
	}

	for _, copiedBlock := range blocks {
		blockX := player.X + copiedBlock.X
		blockY := player.Y + copiedBlock.Y
		copiedBlock.X = blockX
		copiedBlock.Y = blockY
		b.cells[blockX][blockY].Colour = copiedBlock.Colour
	}
}

func (b *Board) checkCompleteRows() {
	/*
		Check if there are any complete rows
	*/

	fullRows := make(map[int]bool)

	// remember to ignore first and last grey rows
	boardWidth := len(b.cells)
	boardHeight := len(b.cells[0])

	for y := 1; y < (boardHeight - 1); y++ {

		rowFull := true
		for x := 1; x < (boardWidth - 1); x++ {
			if b.cells[x][y].Colour == Empty {
				rowFull = false
				continue
			}
		}

		if rowFull {
			// add full row to list
			fullRows[y] = true
		}
	}

	if len(fullRows) > 0 {
		b.destroyRows(fullRows)
	}
}

func (b *Board) destroyRows(rows map[int]bool) {
	/*
	   This method destroys ALL rows passed in the list
	   then readjusts the board to account for the destroyed rows

	   Given a list of rows to destroy
	   The rows must be removed and the remaining rows adjusted
	   and the the board must be refilled
	*/

	boardHeight := len(b.cells)
	for y := 1; y < boardHeight-1; y++ {
		if _, ok := rows[y]; ok {
			//this is a row to destroy
			//move all rows above it down 1 row
			//self.move_row_down(y
			b.moveRowDown(y)
		}
	}
}

func (b *Board) moveRowDown(lastRow int) {
	// we don't "really" move the row down, we just copy the colour of the blocks to the row below

	boardWidth := len(b.cells)
	boardHeight := len(b.cells[0])
	firstRow := boardHeight - 2

	// new blank row at bottom
	for x := 1; x < boardWidth-1; x++ {
		b.cells[x][lastRow].Colour = Empty
	}

	for y := lastRow; y < firstRow; y++ {
		for x := 1; x < boardWidth-1; x++ {
			// move cell from row above
			b.cells[x][y].Colour = b.cells[x][y+1].Colour
		}

	}
	// new blank row at top
	for x := 1; x < boardWidth-1; x++ {
		b.cells[x][firstRow].Colour = Empty
	}

}
