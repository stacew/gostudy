package scenes

import (
	"log"

	"stacew/gostudy/tuckersgame/12jangi/scenemanager"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) StartUp() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("images/start.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}
func (s *StartScene) Update(screen *ebiten.Image) {
	screen.DrawImage(s.startImg, nil)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})
	}
}
