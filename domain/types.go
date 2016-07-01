package domain

type AudioManager struct {
}

type HighScores struct {
}

type Player struct {
	score int
	level int
	state PlayerState
}

type Block struct {
	x, y   int
	colour BlockColour
}

type Board struct {
	cells [][]Block
}

type Game struct {
	state  GameState
	board  Board
	player Player
}

type Teletris struct {
	audioMgr   AudioManager
	highScores HighScores
	game       Game
}

func NewGame() *Game {
	g := new(Game)
	return g
}

func (g *Game) StartGame() {
	//TODO
}

func NewBoard() *Board {

	b := new(Board)
	b.resetBoard()
	return b
}

func (b *Board) resetBoard() {
	// fill with blank cells
	b.cells = make([][]Block, BoardWidth)

	for x := 0; x < BoardWidth; x++ {
		b.cells[x] = make([]Block, BoardHeight)
		for y := 0; y < BoardHeight; y++ {
			b.cells[x][y] = Block{x: x, y: y, colour: Empty}
		}
	}

	// add grey surrounding blocks
	y := 0
	for x := 0; x < BoardHeight; x++ {
		b.cells[x][y].colour = Grey
	}
	y = BoardHeight - 1
	for x := 0; x < BoardHeight; x++ {
		b.cells[x][y].colour = Grey
	}
	x := 0
	for y := 0; y < BoardHeight; y++ {
		b.cells[x][y].colour = Grey
	}

	x = BoardWidth - 1
	for y := 0; y < BoardHeight; y++ {
		b.cells[x][y].colour = Grey
	}

}
