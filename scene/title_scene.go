package scene

import (
	"fmt"
	"image"
	"runtime"
	"sync"

	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene/config"
	"github.com/telecoda/gomo-simra/simra"
)

// TitleScene represents a scene object for TitleScene
type TitleScene struct {
	sync.Mutex
	Game       *domain.Game
	background *simra.Sprite
}

// Initialize initializes TitleScene scene
func (t *TitleScene) Initialize() {
	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	fmt.Printf("TEMP: before title init\n")
	ReportMemoryUsage()
	// initialize sprites
	t.initialize()
	fmt.Printf("TEMP: after title init\n")
	ReportMemoryUsage()
}

func (t *TitleScene) initialize() {
	// add background sprite
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.initBackground()
	t.background.AddTouchListener(t)
}

func (t *TitleScene) Destroy() {
	fmt.Printf("TEMP: before title destroy\n")
	ReportMemoryUsage()
	go t.destroy()
}

func (t *TitleScene) destroy() {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.background = nil
	runtime.GC()
	fmt.Printf("TEMP: after title destroy\n")
	ReportMemoryUsage()
}

func (t *TitleScene) initBackground() {
	// add background sprite
	t.background = &simra.Sprite{}
	t.background.W = float32(config.ScreenWidth)
	t.background.H = float32(config.ScreenHeight)

	// put center of screen
	t.background.X = config.ScreenWidth / 2
	t.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("title.png",
		image.Rect(0, 0, int(t.background.W), int(t.background.H)),
		t.background)
}

func (t *TitleScene) Drive() {
}

func (t *TitleScene) OnTouchBegin(x, y float32) {
}

func (t *TitleScene) OnTouchMove(x, y float32) {
}

func (t *TitleScene) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	simra.GetInstance().SetScene(&IntroScene{Game: t.Game})
}
