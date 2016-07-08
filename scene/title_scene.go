package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene/config"
)

// TitleScene represents a scene object for TitleScene
type TitleScene struct {
	Game       *domain.Game
	background simra.Sprite
}

// Initialize initializes TitleScene scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (t *TitleScene) Initialize() {
	simra.LogDebug("[IN]")
	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	t.initialize()

	simra.LogDebug("[OUT]")
}

func (t *TitleScene) initialize() {
	// add background sprite
	t.initBackground()
	t.background.AddTouchListener(t)
}

func (t *TitleScene) initBackground() {
	// add background sprite
	t.background.W = float32(config.ScreenWidth)
	t.background.H = float32(config.ScreenHeight)

	// put center of screen
	t.background.X = config.ScreenWidth / 2
	t.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("space_title.png",
		image.Rect(0, 0, int(t.background.W), int(t.background.H)),
		&t.background)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (t *TitleScene) Drive() {
}

// OnTouchBegin is called when TitleScene scene is Touched.
// It is caused by calling AddtouchListener for TitleScene.background sprite.
func (t *TitleScene) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when TitleScene scene is Touched and moved.
// It is caused by calling AddtouchListener for TitleScene.background sprite.
func (t *TitleScene) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when TitleScene scene is Touched and it is released.
// It is caused by calling AddtouchListener for TitleScene.background sprite.
func (t *TitleScene) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	t.Game.StartGame()
	simra.GetInstance().SetScene(&LevelScene{game: t.Game})
}
