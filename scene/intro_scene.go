package scene

import (
	"fmt"
	"image"
	"time"

	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene/config"
	"github.com/telecoda/gomo-simra/simra"
)

// IntroScene represents a scene object for IntroScene
type IntroScene struct {
	Game         *domain.Game
	introSprites [5]simra.Sprite
	currentPage  int
}

// Initialize initializes IntroScene
func (i *IntroScene) Initialize() {
	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	i.initialize()
}

func (i *IntroScene) initialize() {
	// add background sprite
	i.initBackground()
	i.introSprites[4].AddTouchListener(i)
}

func (i *IntroScene) initBackground() {
	// add sprites

	i.currentPage = len(i.introSprites) - 1
	for n := i.currentPage; n >= 0; n-- {

		i.introSprites[n].W = float32(config.ScreenWidth)
		i.introSprites[n].H = float32(config.ScreenHeight)

		// put center of screen
		i.introSprites[n].X = config.ScreenWidth / 2
		i.introSprites[n].Y = config.ScreenHeight / 2

		simra.GetInstance().AddSprite(fmt.Sprintf("intro-%d.png", n),
			image.Rect(0, 0, int(i.introSprites[n].W), int(i.introSprites[n].H)),
			&i.introSprites[n])
	}

	i.currentPage = 0

}

func (i *IntroScene) hideSprite(idx int) {
	// slide sprite down off screen
	for n := 0; n < config.ScreenHeight; n++ {
		switch idx {
		case 0: // left
			i.introSprites[idx].X--
		case 1: // right
			i.introSprites[idx].X++
		case 2: // down
			i.introSprites[idx].Y--
		case 3: // rotate
			i.introSprites[idx].R += 0.01
			i.introSprites[idx].W--
			i.introSprites[idx].H--
		}
		time.Sleep(2 * time.Millisecond)
	}

}

func (i *IntroScene) Drive() {
	//   i.introSprites[]
}

func (i *IntroScene) OnTouchBegin(x, y float32) {
}

func (i *IntroScene) OnTouchMove(x, y float32) {
}

func (i *IntroScene) OnTouchEnd(x, y float32) {
	// on end tap decrease current page counter & hide sprite
	go i.hideSprite(i.currentPage)
	i.currentPage++

	if i.currentPage > len(i.introSprites)-1 {
		// scene end. go to next scene
		i.Game.StartGame()
		simra.GetInstance().SetScene(&LevelScene{Game: i.Game})
	}
}
