package apple

import (
	"fmt"
	"image/color"
	"math/rand"
	"snake/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Apple struct {
	Pos   types.Position
	size  int
	color color.RGBA
}

func (a *Apple) Init(p types.Position, s int, c color.RGBA) {
	a.Pos = p
	a.size = s
	a.color = c
}

func (a *Apple) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen, float64(a.Pos.X)*float64(a.size), float64(a.Pos.Y)*float64(a.size),
		float64(a.size), float64(a.size), a.color)
}

func (a *Apple) GenNewPos(sizeX, sizeY int) {
	if ebiten.IsKeyPressed(ebiten.KeyN) {
		a.Pos.Y = rand.Intn(sizeY / a.size)
		a.Pos.X = rand.Intn(sizeX / a.size)
	}
}
