package actor

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/stacew/gostudy/tuckersgame/rungame/animation"
)

//runner.png frame
const (
	//runner.png 1 runner frame size
	FrameWidth  = 32
	FrameHeight = 32
	//runner frame count
	RunningFrameCount = 8
	IdleFrameCount    = 5
	//running img state frame
	RunningX = 0
	RunningY = 32
	//idle img state frame
	IdleX = 0
	IdleY = 0
)

//러너 상태
type RunnerState int

const (
	Idle RunnerState = iota
	Running
	Fast
)

//동작 변화 카운트
const (
	idleFrameChangeCnt    = 8
	runningFrameChangeCnt = 4
	fastFrameChangeCnt    = 2
)

type Runner struct {
	X, Y      float64
	state     RunnerState
	animation *animation.Handler
}

func NewRunner(x, y float64) *Runner {
	r := &Runner{}
	r.X, r.Y = x, y
	r.state = Idle

	r.animation = animation.New()

	runnerImg, _, err := ebitenutil.NewImageFromFile("images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	//sprites
	sprites := make([]*ebiten.Image, IdleFrameCount)
	for i := 0; i < IdleFrameCount; i++ {
		sx := IdleX + FrameWidth*i
		sy := IdleY

		subRect := image.Rect(sx, sy, sx+FrameWidth, sy+FrameHeight)
		sprites[i] = runnerImg.SubImage(subRect).(*ebiten.Image)
	}

	r.animation.Add("Idle", sprites, idleFrameChangeCnt)

	//sprites
	sprites = make([]*ebiten.Image, RunningFrameCount)
	for i := 0; i < RunningFrameCount; i++ {
		sx := RunningX + FrameWidth*i
		sy := RunningY

		subRect := image.Rect(sx, sy, sx+FrameWidth, sy+FrameHeight)
		sprites[i] = runnerImg.SubImage(subRect).(*ebiten.Image)
	}

	r.animation.Add("Run", sprites, runningFrameChangeCnt)
	r.animation.Add("Fast", sprites, fastFrameChangeCnt)

	return r
}
func (r *Runner) SetState(state RunnerState) {
	r.state = state
	switch r.state {
	case Idle:
		r.animation.Play("Idle")
	case Running:
		r.animation.Play("Run")
	case Fast:
		r.animation.Play("Fast")
	}
}
func (r *Runner) Update(screen *ebiten.Image) {
	r.animation.Update(screen, r.X, r.Y)
}
