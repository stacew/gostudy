package main

import (
	"log"
	"math/rand"
	"stacew/gostudy/tuckersgame/rungame/global"
	"stacew/gostudy/tuckersgame/rungame/scenemanager"
	"stacew/gostudy/tuckersgame/rungame/scenes"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	//random seed set
	rand.Seed(time.Now().UnixNano())
	//Set StartScene
	scenemanager.SetScene(&scenes.StartScene{})
	//run
	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 2.0, "Run Game")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
