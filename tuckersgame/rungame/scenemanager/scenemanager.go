package scenemanager

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Update(*ebiten.Image)
	StartUp()
}

type scenemanager struct {
	currentScene Scene
}

var manager *scenemanager

//init main()함수 처럼 자동 호출 약속된 함수
func init() {
	manager = &scenemanager{}
}

func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		manager.currentScene.Update(screen)
	}
	return nil
}

func SetScene(scene Scene) {
	manager.currentScene = scene
	scene.StartUp()
}
