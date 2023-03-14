package main

import (
	"snake/screen"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := screen.Game{}
	game.Init()
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
