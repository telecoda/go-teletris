package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene/config"
)

// GameOverScene represents a scene object for GameOver
type GameOverScene struct {
	Game       *domain.Game
	background simra.Sprite
}

func (g *GameOverScene) Initialize() {
	simra.LogDebug("[IN]")
	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	g.initialize()

	simra.LogDebug("[OUT]")
}

func (g *GameOverScene) initialize() {
	// add background sprite
	g.initBackground()
	g.background.AddTouchListener(g)
}

func (g *GameOverScene) initBackground() {
	// add background sprite
	g.background.W = float32(config.ScreenWidth)
	g.background.H = float32(config.ScreenHeight)

	// put center of screen
	g.background.X = config.ScreenWidth / 2
	g.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("game_over.png",
		image.Rect(0, 0, int(g.background.W), int(g.background.H)),
		&g.background)
}

func (g *GameOverScene) Drive() {
}

func (g *GameOverScene) OnTouchBegin(x, y float32) {
}

func (g *GameOverScene) OnTouchMove(x, y float32) {
}

func (g *GameOverScene) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	simra.GetInstance().SetScene(&TitleScene{Game: g.Game})
}
