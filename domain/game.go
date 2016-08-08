package domain

import (
	"log"
	"time"

	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/exp/audio"
)

type AudioManager struct {
}

type HighScores struct {
}

type Game struct {
	state       GameState
	board       Board
	Player      Player
	audioPlayer *audio.Player
	dirty       bool
}

type Teletris struct {
	audioMgr   AudioManager
	highScores HighScores
	game       Game
}

func NewGame() *Game {
	g := new(Game)
	g.state = Menu
	return g
}

func (g *Game) initAudio() {
	rc, err := asset.Open("game_music_16.wav")
	if err != nil {
		log.Fatal(err)
	}
	g.audioPlayer, err = audio.NewPlayer(rc, audio.Mono16, 16000)
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) StartGame() {
	// Start a new game
	g.board = NewBoard()
	g.initAudio()
	// init player state
	g.Player.Init()
	g.board.reset()
	g.audioPlayer.Seek(0)
	g.audioPlayer.SetVolume(1.0)
	g.audioPlayer.Play()

	g.state = Playing
	g.Player.setNextRandomShape()
	g.Player.setNextRandomShape()

	go g.run()

}

func (g *Game) SetBoardDirty() {
	g.dirty = true
}

func (g *Game) IsBoardDirty() bool {
	return g.dirty
}

func (g *Game) CleanBoard() {
	g.dirty = false
}

func (g *Game) PauseGame() {
	// TODO
}

func (g *Game) ResumeGame() {
	// TODO
}

func (g *Game) GameOver() {
	log.Println("Game Over!")
	g.state = GameOver
	g.audioPlayer.Stop()

	// TODO update high scores
}

func (g *Game) run() {

	for g.state == Playing {
		// drop blocks exery x milliseconds

		// calc delay speed
		delaySpeed := BlockStartSpeed - (g.Player.level - 1*LevelSpeedIncrease)

		time.Sleep(time.Duration(delaySpeed) * time.Millisecond)
		g.MoveDown()
	}

}

func (g *Game) LoadHighScores() {
	// TODO
}

func (g *Game) SaveHighScores() {
	// TODO
}

func (g *Game) GetBlocks() *[][]*Block {
	return &g.board.cells
}

func (g *Game) newShape() {
	g.Player.setNextRandomShape()
	if !g.board.canPlayerFitAt(&g.Player, g.Player.X+1, g.Player.Y) {
		g.GameOver()
	}
}

func (g *Game) GetState() GameState {
	return g.state
}

func (g *Game) Rotate() bool {
	// rotate shape
	g.Player.Rotate()

	// test if player's block fits
	if !g.board.canPlayerFitAt(&g.Player, g.Player.X, g.Player.Y) {
		// rotate it back
		g.Player.RotateBack()
		return false
	}
	return true
}

func (g *Game) MoveDown() bool {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X, g.Player.Y-1) {
		g.Player.MoveDown()
		return true
	} else {
		g.board.addShapeToBoard(&g.Player)
		g.newShape()
		g.board.checkCompleteRows()
		g.SetBoardDirty()
		return false
	}
}

func (g *Game) MoveLeft() bool {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X-1, g.Player.Y) {
		g.Player.MoveLeft()
		return true
	}
	return false
}

func (g *Game) MoveRight() bool {
	// test if player's block fits
	if g.board.canPlayerFitAt(&g.Player, g.Player.X+1, g.Player.Y) {
		g.Player.MoveRight()
		return true
	}
	return false
}
