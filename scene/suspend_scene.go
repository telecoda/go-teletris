package scene

import (
	"fmt"
	"time"
)

// SuspendScene does nothing
type SuspendScene struct {
}

func (s *SuspendScene) Initialize() {
	fmt.Printf("TEMP: suspend scene initialise: %s\n", time.Now())

}

func (s *SuspendScene) Drive() {

	fmt.Printf("TEMP: SuspendScene.Drive()\n")
}

// // OnTouchBegin is called when TitleScene scene is Touched.
// // It is caused by calling AddtouchListener for TitleScene.background sprite.
// func (t *TitleScene) OnTouchBegin(x, y float32) {
// }

// // OnTouchMove is called when TitleScene scene is Touched and moved.
// // It is caused by calling AddtouchListener for TitleScene.background sprite.
// func (t *TitleScene) OnTouchMove(x, y float32) {
// }

// // OnTouchEnd is called when TitleScene scene is Touched and it is released.
// // It is caused by calling AddtouchListener for TitleScene.background sprite.
// func (t *TitleScene) OnTouchEnd(x, y float32) {
// 	// scene end. go to next scene
// 	t.Game.StartGame()
// 	simra.GetInstance().SetScene(&LevelScene{Game: t.Game})
// }
