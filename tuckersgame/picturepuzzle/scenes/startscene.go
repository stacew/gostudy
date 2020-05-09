package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/stacew/gostudy/tuckersgame/picturepuzzle/font"
	"github.com/stacew/gostudy/tuckersgame/picturepuzzle/global"
	"github.com/stacew/gostudy/tuckersgame/picturepuzzle/scenemanager"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) StartUp() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}
func (s *StartScene) Update(screen *ebiten.Image) {
	screen.DrawImage(s.startImg, nil)

	fontSize := 2
	width := font.TextWidth(global.StartSceneFirstText, fontSize)
	font.DrawTextWithShadow(screen, global.StartSceneFirstText,
		global.ScreenWidth/2-width/2, global.ScreenHeight/2, fontSize, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{puzzleColumns: global.ScreenWidth / 100, puzzleRows: global.ScreenHeight / 100})
	}
}
