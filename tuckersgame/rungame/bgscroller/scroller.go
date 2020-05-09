package bgscroller

import (
	"github.com/hajimehoshi/ebiten"
)

func New(bgImg *ebiten.Image, speed int) *Scroller {
	return &Scroller{bgImg, speed, 0}
}

type Scroller struct {
	bgImg  *ebiten.Image
	speed  int
	frames int
}

func (s *Scroller) Set(speed int) {
	s.speed = speed
}

func (s *Scroller) Update(screen *ebiten.Image) {
	
	s.frames = s.frames + s.speed

	//scroll left
	bgWidth, _ := s.bgImg.Size()
	backX := -(float64((s.frames / 2) % bgWidth))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(backX, 0)
	screen.DrawImage(s.bgImg, op)

	//scroll right 정석
	op.GeoM.Translate(float64(bgWidth), 0) //op값을 이동 시킨다.
	screen.DrawImage(s.bgImg, op)
	//scroll right 쉬운 이해방법
	// op2 := &ebiten.DrawImageOptions{} //op를 새로 만든다.
	// backX2 := backX + float64(bgWidth)
	// op2.GeoM.Translate(backX2, 0)
	// screen.DrawImage(s.bgImg, op2)
}
