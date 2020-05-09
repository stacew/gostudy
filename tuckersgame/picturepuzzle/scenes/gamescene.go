package scenes

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/stacew/gostudy/tuckersgame/picturepuzzle/global"
)

type GameScene struct {
	puzzleColumns, puzzleRows int

	bgImg          *ebiten.Image
	subImages      []*ebiten.Image
	board          [][]int
	blankX, blankY int
}

func (g *GameScene) StartUp() {
	var err error
	g.bgImg, _, err = ebitenutil.NewImageFromFile("images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	g.subImages = make([]*ebiten.Image, g.puzzleColumns*g.puzzleRows)
	g.board = make([][]int, g.puzzleColumns)
	for i := 0; i < g.puzzleColumns; i++ {
		g.board[i] = make([]int, g.puzzleRows)
	}

	width := global.ScreenWidth / g.puzzleColumns
	height := global.ScreenHeight / g.puzzleRows
	for i := 0; i < g.puzzleColumns; i++ {
		for j := 0; j < g.puzzleRows; j++ {
			g.subImages[j*g.puzzleColumns+i] =
				g.bgImg.SubImage(image.Rect(i*width, j*height, (i+1)*width, (j+1)*height)).(*ebiten.Image)
		}
	}
	//board & blank init
	for i := 0; i < g.puzzleColumns; i++ {
		for j := 0; j < g.puzzleRows; j++ {
			g.board[i][j] = j*g.puzzleColumns + i
		}
	}
	g.blankX, g.blankY = g.puzzleColumns-1, g.puzzleRows-1
	g.board[g.blankX][g.blankY] = -1
	//boar rand set
	nRandCount := g.puzzleColumns * g.puzzleRows
	for i := 0; i < nRandCount; i++ {
		randX1, randY1 := rand.Intn(g.puzzleColumns), rand.Intn(g.puzzleRows)
		randX2, randY2 := rand.Intn(g.puzzleColumns), rand.Intn(g.puzzleRows)
		//swap
		g.board[randX1][randY1], g.board[randX2][randY2] = g.board[randX2][randY2], g.board[randX1][randY1]
	}
	//find blank
	for i := 0; i < g.puzzleColumns; i++ {
		for j := 0; j < g.puzzleRows; j++ {
			if g.board[i][j] == -1 {
				g.blankX, g.blankY = i, j
				break
			}
		}
	}
}

func (g *GameScene) Update(screen *ebiten.Image) {

	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		if g.blankY > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY-1]
			g.board[g.blankX][g.blankY-1] = -1
			g.blankY--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		if g.blankY < g.puzzleRows-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][g.blankY+1]
			g.board[g.blankX][g.blankY+1] = -1
			g.blankY++
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		if g.blankX > 0 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX-1][g.blankY]
			g.board[g.blankX-1][g.blankY] = -1
			g.blankX--
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		if g.blankX < g.puzzleColumns-1 {
			g.board[g.blankX][g.blankY] = g.board[g.blankX+1][g.blankY]
			g.board[g.blankX+1][g.blankY] = -1
			g.blankX++
		}
	}

	width := global.ScreenWidth / g.puzzleColumns
	height := global.ScreenHeight / g.puzzleRows

	for i := 0; i < g.puzzleColumns; i++ {
		for j := 0; j < g.puzzleRows; j++ {
			if g.board[i][j] == -1 {
				continue
			}

			x := i * width
			y := j * height

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(g.subImages[g.board[i][j]], opts)
		}
	}
}
