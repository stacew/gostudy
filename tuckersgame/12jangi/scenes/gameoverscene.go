package scenes

import (
	"log"

	"github.com/stacew/gostudy/tuckersgame/12jangi/scenemanager"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GameoverScene struct {
	gameOverImg *ebiten.Image
}

func (g *GameoverScene) StartUp() {
	var err error
	g.gameOverImg, _, err = ebitenutil.NewImageFromFile("images/gameover.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}
func (g *GameoverScene) Update(screen *ebiten.Image) {
	screen.DrawImage(g.gameOverImg, nil)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&StartScene{})
	}
}
