package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/stacew/gostudy/tuckersgame/12jangi/global"
	"github.com/stacew/gostudy/tuckersgame/12jangi/scenemanager"
	"github.com/stacew/gostudy/tuckersgame/12jangi/scenes"
)

func main() {

	//Set StartScene
	scenemanager.SetScene(&scenes.StartScene{})
	//run
	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
