package scenes

import (
	"log"

	"github.com/stacew/gostudy/tuckersgame/rungame/actor"
	"github.com/stacew/gostudy/tuckersgame/rungame/bgscroller"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/stacew/gostudy/tuckersgame/rungame/global"
)

type GameScene struct {
	bgScroller *bgscroller.Scroller
	runner     *actor.Runner
}

func (g *GameScene) StartUp() {
	bgImg, _, err := ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.bgScroller = bgscroller.New(bgImg, 2)

	g.runner = actor.NewRunner(0, global.ScreenHeight/2)
	g.runner.SetState(actor.Running)

	//Set이랑 Fast 상태 추가해 봄
	// g.bgScroller.Set(10)
	// g.runner.SetState(actor.Fast)
}

func (g *GameScene) Update(screen *ebiten.Image) {

	g.bgScroller.Update(screen)

	//Running Animation
	g.runner.Update(screen)
}
