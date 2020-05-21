package main

import (
	"log"
	"math/rand"
	"stacew/gostudy/tuckersgame/picturepuzzle/global"
	"stacew/gostudy/tuckersgame/picturepuzzle/scenemanager"
	"stacew/gostudy/tuckersgame/picturepuzzle/scenes"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	//random seed set
	rand.Seed(time.Now().UnixNano())
	//Set StartScene
	scenemanager.SetScene(&scenes.StartScene{})
	//run
	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
