package types

import "image/color"

type Position struct {
	X int
	Y int
}

type SnakeElement struct {
	X     int
	Y     int
	Color color.Color
}
