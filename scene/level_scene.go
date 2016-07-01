package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/simra"
	"github.com/telecoda/go-teletris/scene/config"
)

// LevelScene represents a scene object for LevelScene
type LevelScene struct {
	ball       simra.Sprite
	background simra.Sprite
	ctrlup     simra.Sprite
	ctrldown   simra.Sprite
	ctrlleft   simra.Sprite
	ctrlright  simra.Sprite
	// buttonState represents which ctrl is pressed (or no ctrl pressed)
	buttonState int

	buttonReplaced bool
}

const (
	ctrlNop = iota
	ctrlUp
	ctrlDown
	ctrlLeft
	ctrlRight
)

// Initialize initializes LevelScene scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (l *LevelScene) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(l)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(LevelScene)

	// initialize sprites
	l.initSprites()
	l.buttonReplaced = false

	simra.LogDebug("[OUT]")
}

// OnTouchBegin is called when LevelScene scene is Touched.
func (l *LevelScene) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when LevelScene scene is Touched and moved.
func (l *LevelScene) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when LevelScene scene is Touched and it is released.
func (l *LevelScene) OnTouchEnd(x, y float32) {
	l.buttonState = ctrlNop
}

func (l *LevelScene) initSprites() {
	l.initBackground()
	l.initctrlDown()
	l.initctrlUp()
	l.initctrlLeft()
	l.initctrlRight()
	l.initBall()

}

func (l *LevelScene) initBall() {
	// set size of ball
	l.ball.W = float32(48)
	l.ball.H = float32(48)

	// put center of screen at start
	l.ball.X = config.ScreenWidth / 2
	l.ball.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("ball.png",
		image.Rect(0, 0, int(l.ball.W), int(l.ball.H)),
		&l.ball)
}

func (l *LevelScene) initBackground() {
	// add background sprite
	l.background.W = float32(config.ScreenWidth)
	l.background.H = float32(config.ScreenHeight)

	// put center of screen
	l.background.X = config.ScreenWidth / 2
	l.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("space_background.png",
		image.Rect(0, 0, int(l.background.W), int(l.background.H)),
		&l.background)

}

const (
	ctrlMarginLeft      = 10
	ctrlMarginBottom    = 10
	ctrlMarginBetween   = 10
	buttonMarginRight   = 20
	buttonMarginBottom  = 20
	buttonMarginBetween = 10
)

// ctrlUp
type ctrlUpTouchListener struct {
	parent *LevelScene
}

func (l *ctrlUpTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (l *ctrlUpTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (l *ctrlUpTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlUp End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlUp() {
	// set size of ctrlUp
	l.ctrlup.W = float32(120)
	l.ctrlup.H = float32(120)

	// put ctrlUp on middle bottom of screen
	l.ctrlup.X = (config.ScreenWidth / 2)
	l.ctrlup.Y = ctrlMarginBottom + l.ctrldown.H + ctrlMarginBetween + (l.ctrlup.H / 2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlup.W), int(l.ctrlup.H)),
		&l.ctrlup)

	// add touch listener for sprite
	ctrlup := &ctrlUpTouchListener{}
	l.ctrlup.AddTouchListener(ctrlup)
	ctrlup.parent = l
}

// ctrlDown
type ctrlDownTouchListener struct {
	parent *LevelScene
}

func (l *ctrlDownTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (l *ctrlDownTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (l *ctrlDownTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlDown End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlDown() {
	// set size of ctrlDown
	l.ctrldown.W = float32(120)
	l.ctrldown.H = float32(120)

	// put ctrlDown on middle bottom of screen
	l.ctrldown.X = (config.ScreenWidth / 2)
	l.ctrldown.Y = ctrlMarginBottom + (l.ctrldown.H / 2)

	// rotate arrow to indicate down control
	l.ctrldown.R = math.Pi

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrldown.W), int(l.ctrldown.H)),
		&l.ctrldown)

	// add touch listener for sprite
	ctrldown := &ctrlDownTouchListener{}
	l.ctrldown.AddTouchListener(ctrldown)
	ctrldown.parent = l
}

// ctrlLeft
type ctrlLeftTouchListener struct {
	parent *LevelScene
}

func (LevelScene *ctrlLeftTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft Begin!")

	ctrl := LevelScene.parent
	ctrl.buttonState = ctrlLeft

	simra.LogDebug("[OUT]")
}

func (l *ctrlLeftTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlLeft

	simra.LogDebug("[OUT]")
}

func (l *ctrlLeftTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlLeft End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlLeft() {
	// set size of ctrlLeft
	l.ctrlleft.W = float32(120)
	l.ctrlleft.H = float32(120)

	// put ctrlLeft on left bottom
	l.ctrlleft.X = (config.ScreenWidth / 2) - l.ctrlleft.W - ctrlMarginBetween
	l.ctrlleft.Y = ctrlMarginBottom + l.ctrlleft.H

	// rotate arrow to indicate left control
	l.ctrlleft.R = math.Pi / 2

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlleft.W), int(l.ctrlleft.H)),
		&l.ctrlleft)

	// add touch listener for sprite
	ctrlleft := &ctrlLeftTouchListener{}
	l.ctrlleft.AddTouchListener(ctrlleft)
	ctrlleft.parent = l
}

// ctrlRight
type ctrlRightTouchListener struct {
	parent *LevelScene
}

func (l *ctrlRightTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlRight Begin!")

	ctrl := l.parent
	ctrl.buttonState = ctrlRight

	simra.LogDebug("[OUT]")
}

func (l *ctrlRightTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlRight Move!")

	ctrl := l.parent
	ctrl.buttonState = ctrlRight

	simra.LogDebug("[OUT]")
}

func (l *ctrlRightTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlRight End")

	ctrl := l.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (l *LevelScene) initctrlRight() {
	// set size of ctrlRight
	l.ctrlright.W = float32(120)
	l.ctrlright.H = float32(120)

	// put ctrlRight on left bottom
	l.ctrlright.X = (config.ScreenWidth / 2) + l.ctrlright.W + ctrlMarginBetween
	l.ctrlright.Y = ctrlMarginBottom + l.ctrlright.H

	// rotate arrow to indicate right control
	l.ctrlright.R = -math.Pi / 2

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(l.ctrlright.W), int(l.ctrlright.H)),
		&l.ctrlright)

	// add touch listener for sprite
	ctrlright := &ctrlLeftTouchListener{}
	l.ctrlright.AddTouchListener(ctrlright)
	ctrlright.parent = l
}

// func (LevelScene *LevelScene) replaceButtonColor() {
// 	simra.LogDebug("IN")
// 	// red changes to blue
// 	LevelScene.buttonRed.ReplaceTexture("blue_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonBlue.W), int(LevelScene.buttonBlue.H)))
// 	// blue changes to red
// 	LevelScene.buttonBlue.ReplaceTexture("red_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonRed.W), int(LevelScene.buttonRed.H)))

// 	LevelScene.buttonReplaced = true
// 	simra.LogDebug("OUT")
// }

// func (LevelScene *LevelScene) originalButtonColor() {
// 	simra.LogDebug("IN")
// 	// set red button to buttonRed
// 	LevelScene.buttonRed.ReplaceTexture("red_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonBlue.W), int(LevelScene.buttonBlue.H)))
// 	// set blue button to buttonBlue
// 	LevelScene.buttonBlue.ReplaceTexture("blue_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonRed.W), int(LevelScene.buttonRed.H)))

// 	LevelScene.buttonReplaced = false
// 	simra.LogDebug("OUT")
// }

// // ButtonBlueTouchListener represents a listener object
// // to notify touch event of Blue Button
// type ButtonBlueTouchListener struct {
// 	parent *LevelScene
// }

// // OnTouchBegin is called when Blue Button is Touched.
// func (LevelScene *ButtonBlueTouchListener) OnTouchBegin(x, y float32) {
// 	simra.LogDebug("IN")
// 	if LevelScene.parent.buttonReplaced {
// 		LevelScene.parent.originalButtonColor()
// 	} else {
// 		LevelScene.parent.replaceButtonColor()
// 	}

// 	simra.GetInstance().RemoveSprite(&LevelScene.parent.ball)
// 	simra.LogDebug("OUT")
// }

// // OnTouchMove is called when Blue Button is Touched and moved.
// func (LevelScene *ButtonBlueTouchListener) OnTouchMove(x, y float32) {
// 	// nop
// }

// // OnTouchEnd is called when Blue Button is Touched and it is released.
// func (LevelScene *ButtonBlueTouchListener) OnTouchEnd(x, y float32) {
// 	// nop
// }

// func (LevelScene *LevelScene) initButtonBlue() {
// 	simra.LogDebug("IN")
// 	// set size of button blue
// 	LevelScene.buttonBlue.W = float32(80)
// 	LevelScene.buttonBlue.H = float32(80)

// 	// put button red on right bottom
// 	LevelScene.buttonBlue.X = config.ScreenWidth - buttonMarginRight - LevelScene.buttonBlue.W/2
// 	LevelScene.buttonBlue.Y = buttonMarginBottom + (80) + buttonMarginBetween + LevelScene.buttonBlue.W/2

// 	// add sprite to glpeer
// 	simra.GetInstance().AddSprite("blue_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonBlue.W), int(LevelScene.buttonBlue.H)),
// 		&LevelScene.buttonBlue)

// 	// add touch listener for sprite
// 	listener := &ButtonBlueTouchListener{}
// 	LevelScene.buttonBlue.AddTouchListener(listener)
// 	listener.parent = LevelScene
// 	simra.LogDebug("OUT")
// }

// // ButtonRedTouchListener represents a listener object
// // to notify touch event of Red Button
// type ButtonRedTouchListener struct {
// 	parent *LevelScene
// }

// // OnTouchBegin is called when Red Button is Touched.
// func (LevelScene *ButtonRedTouchListener) OnTouchBegin(x, y float32) {
// 	simra.LogDebug("IN")
// 	if LevelScene.parent.buttonReplaced {
// 		LevelScene.parent.originalButtonColor()
// 	} else {
// 		LevelScene.parent.replaceButtonColor()
// 	}
// 	simra.GetInstance().AddSprite("ball.png",
// 		image.Rect(0, 0, int(LevelScene.parent.ball.W), int(LevelScene.parent.ball.H)),
// 		&LevelScene.parent.ball)
// 	simra.LogDebug("OUT")
// }

// // OnTouchMove is called when Red Button is Touched and moved.
// func (LevelScene *ButtonRedTouchListener) OnTouchMove(x, y float32) {
// 	// nop
// }

// // OnTouchEnd is called when Red Button is Touched and it is released.
// func (LevelScene *ButtonRedTouchListener) OnTouchEnd(x, y float32) {
// 	// nop
// }

// func (LevelScene *LevelScene) initButtonRed() {
// 	// set size of button red
// 	LevelScene.buttonRed.W = float32(80)
// 	LevelScene.buttonRed.H = float32(80)

// 	// put button red on right bottom
// 	LevelScene.buttonRed.X = config.ScreenWidth - buttonMarginRight - LevelScene.buttonBlue.W -
// 		buttonMarginBetween - LevelScene.buttonRed.W/2
// 	LevelScene.buttonRed.Y = buttonMarginBottom + (LevelScene.buttonRed.H / 2)

// 	// add sprite to glpeer
// 	simra.GetInstance().AddSprite("red_circle.png",
// 		image.Rect(0, 0, int(LevelScene.buttonRed.W), int(LevelScene.buttonRed.H)),
// 		&LevelScene.buttonRed)

// 	// add touch listener for sprite
// 	listener := &ButtonRedTouchListener{}
// 	LevelScene.buttonRed.AddTouchListener(listener)
// 	listener.parent = LevelScene
// }

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (l *LevelScene) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}

	switch l.buttonState {
	case ctrlUp:
		l.ball.Y++
	case ctrlDown:
		l.ball.Y--
	case ctrlLeft:
		l.ball.X--
	case ctrlRight:
		l.ball.X++
	}

	l.ball.R = float32(degree) * math.Pi / 180
}
