package main

import (
	"log"
	"stacew/gostudy/tuckersgame/12jangi/global"
	"stacew/gostudy/tuckersgame/12jangi/scenemanager"
	"stacew/gostudy/tuckersgame/12jangi/scenes"

	"github.com/hajimehoshi/ebiten"
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
