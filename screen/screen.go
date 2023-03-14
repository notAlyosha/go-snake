package screen

import (
	"image/color"

	"snake/apple"
	"snake/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 300
	screenHeight = 300
	step         = 20
)

type Game struct {
	apple apple.Apple
}

func (g *Game) Init() {
	g.apple.Init(types.Position{X: 2, Y: 2}, step, color.RGBA{R: 90, G: 200, B: 60 * 4, A: 255})
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Test")
	ebiten.SetTPS(5)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.apple.GenNewPos(screenWidth, screenHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawGrid(screen, color.RGBA{R: 90, G: 200, B: 240, A: 255})
	g.apple.Draw(screen)
}

/*==================================================================*/

func drawGrid(screen *ebiten.Image, color color.RGBA) {

	for y := 0; y < screenHeight; y += step {

		ebitenutil.DrawLine(screen, 0, float64(y), screenWidth, float64(y), color)
	}

	for x := 0; x < screenWidth; x += step {

		ebitenutil.DrawLine(screen, float64(x), 0, float64(x), screenHeight, color)
	}
}
