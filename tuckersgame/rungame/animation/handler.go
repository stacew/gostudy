package animation

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type animInfo struct {
	sprites        []*ebiten.Image
	frameChangeCnt int
}

type Handler struct {
	aniMap       map[string]*animInfo
	pCurrentAnim *animInfo
	lastImgIdx   int
	remainFrames int
}

func New() *Handler {
	h := &Handler{}
	h.aniMap = make(map[string]*animInfo)
	return h
}

func (h *Handler) Add(name string, sprites []*ebiten.Image, frameChangeCnt int) {
	if frameChangeCnt < 1 {
		log.Fatal("Frame Change Count Over 1")
	}

	h.aniMap[name] = &animInfo{sprites, frameChangeCnt}
}

func (h *Handler) Play(name string) {
	h.pCurrentAnim = h.aniMap[name]
	h.lastImgIdx = 0
	h.remainFrames = h.pCurrentAnim.frameChangeCnt
}

func (h *Handler) Update(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(h.pCurrentAnim.sprites[h.lastImgIdx], op)

	h.remainFrames--
	if h.remainFrames == 0 {
		h.remainFrames = h.pCurrentAnim.frameChangeCnt

		h.lastImgIdx++
		if len(h.pCurrentAnim.sprites) == h.lastImgIdx {
			h.lastImgIdx = 0
		}
	}
}
