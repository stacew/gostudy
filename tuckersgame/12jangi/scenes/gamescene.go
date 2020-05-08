package scenes

import (
	"log"
	"strconv"

	"github.com/stacew/gostudy/tuckersgame/12jangi/global"
	"github.com/stacew/gostudy/tuckersgame/12jangi/scenemanager"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	BoardWidth  = 4
	BoardHeight = 3
	GimulStartX = 20
	GimulStartY = 23
	GridWidth   = 116
	GridHeight  = 116
)

type GimulType int

const (
	GimulNone GimulType = -1 + iota
	GimulGreenWang
	GimulGreenJa
	GimulGreenJang
	GimulGreenSang
	GimulRedWang
	GimulRedJa
	GimulRedJang
	GimulRedSang
	GimulTypeMax
)

type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

type GameScene struct {
	bgimg       *ebiten.Image
	selectedImg *ebiten.Image
	gimulImgs   [GimulTypeMax]*ebiten.Image

	gameover bool

	board [BoardWidth][BoardHeight]GimulType

	currentTeam TeamType

	selected  bool
	selectedX int
	selectedY int
}

func (g *GameScene) StartUp() {
	var err error
	//img load
	g.bgimg, _, err = ebitenutil.NewImageFromFile("images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.selectedImg, _, err = ebitenutil.NewImageFromFile("images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	for i := GimulGreenWang; i < GimulTypeMax; i++ {
		imgPath := "images/" + strconv.Itoa(int(i+1)) + ".png"
		g.gimulImgs[i], _, err = ebitenutil.NewImageFromFile(imgPath, ebiten.FilterDefault)
		if err != nil {
			log.Fatalf("read file error : %v", err)
		}
	}

	g.gameover = false

	//init Board
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			g.board[i][j] = GimulNone
		}
	}
	g.board[0][0] = GimulGreenSang
	g.board[0][1] = GimulGreenWang
	g.board[0][2] = GimulGreenJang
	g.board[1][1] = GimulGreenJa
	g.board[3][0] = GimulRedSang
	g.board[3][1] = GimulRedWang
	g.board[3][2] = GimulRedJang
	g.board[2][1] = GimulRedJa

	g.currentTeam = TeamGreen
}

func (g *GameScene) Update(screen *ebiten.Image) {
	screen.DrawImage(g.bgimg, nil)

	if g.gameover {
		scenemanager.SetScene(&GameoverScene{})
		return
	}

	//input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight
		if !g.selected {
			if g.board[i][j] != GimulNone && g.currentTeam == getTeamType(g.board[i][j]) {
				g.selected = true
				g.selectedX, g.selectedY = i, j
			}
		} else {
			if g.selectedX == i && g.selectedY == j {
				g.selected = false
			} else {
				g.moveGimul(g.selectedX, g.selectedY, i, j)
			}

		}
	}

	//draw gimul
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			if g.board[i][j] == GimulNone {
				continue
			}

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridHeight*i), float64(GimulStartY+GridHeight*j))

			screen.DrawImage(g.gimulImgs[g.board[i][j]], opts)
		}
	}

	if g.selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(GimulStartX+GridHeight*g.selectedX), float64(GimulStartY+GridHeight*g.selectedY))

		screen.DrawImage(g.selectedImg, opts)
	}
}

func getTeamType(gimulType GimulType) TeamType {

	if gimulType == GimulGreenWang ||
		gimulType == GimulGreenJa ||
		gimulType == GimulGreenJang ||
		gimulType == GimulGreenSang {
		return TeamGreen
	}
	if gimulType == GimulRedWang ||
		gimulType == GimulRedJa ||
		gimulType == GimulRedJang ||
		gimulType == GimulRedSang {
		return TeamRed
	}
	return TeamNone
}

//OnDie calls when gimul is died
func (g *GameScene) onDie(gimulType GimulType) {
	if gimulType == GimulGreenWang ||
		gimulType == GimulRedWang {
		g.gameover = true
	}
}

func (g *GameScene) moveGimul(prevX, prevY, tarX, tarY int) {
	if g.isMovable(prevX, prevY, tarX, tarY) {
		g.onDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulNone, g.board[prevX][prevY]
		g.selected = false
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}

func (g *GameScene) isMovable(prevX, prevY, tarX, tarY int) bool {
	if getTeamType(g.board[prevX][prevY]) == getTeamType(g.board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}

	if tarX >= BoardWidth || tarY >= BoardHeight {
		return false
	}

	switch g.board[prevX][prevY] {
	case GimulGreenJa:
		return prevY == tarY && prevX+1 == tarX
	case GimulRedJa:
		return prevY == tarY && prevX-1 == tarX
	case GimulGreenJang, GimulRedJang:
		return global.Abs(prevX-tarX)+global.Abs(prevY-tarY) == 1
	case GimulGreenSang, GimulRedSang:
		return (global.Abs(prevX-tarX) == 1 && global.Abs(prevY-tarY) == 1)
	case GimulGreenWang, GimulRedWang:
		return (global.Abs(prevX-tarX) <= 1 && global.Abs(prevY-tarY) <= 1)
	}

	return false
}
