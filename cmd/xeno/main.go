package main

import (
	"github.com/cvancleave/xeno/internal/xeno"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	x := xeno.NewXeno()
	x.Start()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Xeno")
	ebiten.SetTPS(10)

	if err := ebiten.RunGame(x); err != nil {
		panic(err)
	}
}
