// +build darwin linux

package main

import (
	"github.com/telecoda/go-teletris/domain"
	"github.com/telecoda/go-teletris/scene"
	"github.com/telecoda/gomo-simra/simra"
)

var game *domain.Game

func main() {
	simra.LogDebug("[IN]")
	engine := simra.GetInstance()

	game = domain.NewGame()

	onStart := make(chan bool)
	onStop := make(chan bool)
	go eventHandle(onStart, onStop)
	engine.Start(onStart, onStop)
	simra.LogDebug("[OUT]")
}

func eventHandle(onStart, onStop chan bool) {
	for {
		select {
		case <-onStart:
			simra.LogDebug("receive chan. onStart")
			engine := simra.GetInstance()
			// TODO: this will be called on rotation.
			// to keep state on rotation, SetScene must not call
			// every onStart.
			engine.SetScene(&scene.TitleScene{Game: game})
		case <-onStop:
			simra.LogDebug("receive chan. onStop")
		}
	}
}
