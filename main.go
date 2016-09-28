// +build darwin linux

package main

import (
	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene"
	"github.com/telecoda/gomo-simra/simra"
)

var game *domain.Game

var titleScene simra.Driver
var levelScene simra.Driver
var suspendScene simra.Driver

func main() {
	engine := simra.GetInstance()

	game = domain.NewGame()
	initScenes()

	onStart := make(chan bool)
	onStop := make(chan bool)
	go eventHandle(onStart, onStop)
	engine.Start(onStart, onStop)
}

func initScenes() {
	if titleScene == nil {
		titleScene = &scene.TitleScene{Game: game}
		levelScene = &scene.LevelScene{Game: game}
		suspendScene = &scene.SuspendScene{}
	}
}

func eventHandle(onStart, onStop chan bool) {
	for {
		select {
		case <-onStart:
			initScenes()
			engine := simra.GetInstance()
			setScene(engine)
			game.ResumeGame()
		case <-onStop:
			// deallocate scenes
			//titleScene = nil
			//levelScene = nil
			//suspendScene = nil
			// stop the music!
			game.SuspendGame()
		}
	}
}

func setScene(engine *simra.Simra) {
	scene.ReportMemoryUsage()
	switch game.GetState() {
	case domain.Menu:
		engine.SetScene(titleScene)
	case domain.Playing:
		engine.SetScene(levelScene)
	case domain.Suspended:
		// use previous scene again
		setPreviousScene(engine)
	case domain.GameOver:
		engine.SetScene(levelScene)
	}
}

func setPreviousScene(engine *simra.Simra) {
	switch game.GetPreviousState() {
	case domain.Menu:
		engine.SetScene(titleScene)
	case domain.Playing:
		engine.SetScene(levelScene)
	case domain.Suspended:
		// do nothing..
		break
	case domain.GameOver:
		engine.SetScene(levelScene)
	}
}
