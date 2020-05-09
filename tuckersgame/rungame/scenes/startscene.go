package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/stacew/gostudy/tuckersgame/rungame/actor"
	"github.com/stacew/gostudy/tuckersgame/rungame/bgscroller"
	"github.com/stacew/gostudy/tuckersgame/rungame/font"
	"github.com/stacew/gostudy/tuckersgame/rungame/global"
	"github.com/stacew/gostudy/tuckersgame/rungame/scenemanager"
)

type StartScene struct {
	bgScroller *bgscroller.Scroller
	runner     *actor.Runner
}

func (s *StartScene) StartUp() {
	bgImg, _, err := ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	s.bgScroller = bgscroller.New(bgImg, 0)
	s.runner = actor.NewRunner(0, global.ScreenHeight/2)
	s.runner.SetState(actor.Idle)
}

func (s *StartScene) Update(screen *ebiten.Image) {

	s.bgScroller.Update(screen)

	//Idle Animation
	s.runner.Update(screen)

	//text
	fontSize := 2
	width := font.TextWidth(global.StartSceneFirstText, fontSize)
	font.DrawTextWithShadow(screen, global.StartSceneFirstText,
		global.ScreenWidth/2-width/2, global.ScreenHeight/2, fontSize, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})
	}
}
