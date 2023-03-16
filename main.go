package main

import (
	"snake/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.Game{}
	game.Init()
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
