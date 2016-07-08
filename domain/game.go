package domain

type AudioManager struct {
}

type HighScores struct {
}

type Game struct {
	state  GameState
	board  Board
	Player Player
}

type Teletris struct {
	audioMgr   AudioManager
	highScores HighScores
	game       Game
}

func NewGame() *Game {
	g := new(Game)
	g.board = NewBoard()
	g.state = Menu
	return g
}

func (g *Game) StartGame() {
	// Start a new game
	// init player state
	g.Player.Init()
	g.board.reset()

	g.state = Playing
	g.Player.setNextRandomShape()
	g.Player.setNextRandomShape()
	//TODO start music

	// TODO start block down timer
}

func (g *Game) PauseGame() {
	// TODO
}

func (g *Game) ResumeGame() {
	// TODO
}

func (g *Game) GameOver() {
	// TODO
	g.state = GameOver

	// TODO update high scores
}

func (g *Game) run() {

}

func (g *Game) LoadHighScores() {
	// TODO
}

func (g *Game) SaveHighScores() {
	// TODO
}

func (g *Game) GetBlocks() *[][]Block {
	return &g.board.cells
}

func (g *Game) newShape() {
	g.Player.setNextRandomShape()
	if g.board.canPlayerFitAt(&g.Player, g.Player.X+1, g.Player.Y) {
		g.GameOver()
	}
}

func (g *Game) Rotate() {
	// rotate shape
	g.Player.Rotate()

	// test if player's block fits
	if !g.board.canPlayerFitAt(&g.Player, g.Player.X, g.Player.Y) {
		// rotate it back
		g.Player.RotateBack()
	}
}

func (g *Game) MoveDown() {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X, g.Player.Y-1) {
		g.Player.MoveDown()
	} else {
		g.board.addShapeToBoard(&g.Player)
		g.newShape()
		g.board.checkCompleteRows()
	}
}

func (g *Game) MoveLeft() {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X-1, g.Player.Y) {
		g.Player.MoveLeft()
	}
}

func (g *Game) MoveRight() {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X+1, g.Player.Y) {
		g.Player.MoveRight()
	}
}
