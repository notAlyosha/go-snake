package snake

import (
	"image/color"
	"snake/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

type Snake struct {
	el     []types.SnakeElement
	dir    int
	size   int
	length int
}

func (s *Snake) Init(w, h, size int) {
	s.el = make([]types.SnakeElement, (w/size)*(h/size))
	s.el[0] = types.SnakeElement{X: 3, Y: 2, Color: color.RGBA{R: 40, G: 70, B: 140, A: 255}}
	s.length = 1
}

func (s *Snake) outOfBounds(w, h, size int) {

	if s.el[0].X > w/size {
		s.el[0].X = 0
		return
	}

	if s.el[0].X < 0 {
		s.el[0].X = w / size
		return
	}

	if s.el[0].Y > h/size {
		s.el[0].Y = 0
		return
	}

	if s.el[0].Y > 0 {
		s.el[0].Y = h / size
		return
	}
}

func (s *Snake) eatsItSelf() bool {
	for i := 0; i < s.length; i += 1 {
		if s.el[0].X == s.el[i].X && s.el[0].X == s.el[i].Y {
			return true
		}
	}
	return false
}

func (s *Snake) shift() {
	for i := 1; i < s.length; i += 1 {
		s.el[i] = s.el[i-1]
	}
}

func (s *Snake) draw(screen *ebiten.Image) {
	for i := 0; i < s.length; i += 1 {
		ebitenutil.DrawRect(screen, float64(s.el[i].X), float64(s.el[i].Y), float64(s.size), float64(s.size), s.el[i].Color)
	}
}
