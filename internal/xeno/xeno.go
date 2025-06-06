package xeno

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Xeno struct {
	currentState string
	states       map[string]state
	buttons      []*button
}

func NewXeno() *Xeno {
	return &Xeno{
		currentState: "idle",
		states:       make(map[string]state),
		buttons:      make([]*button, 0),
	}
}

func (x *Xeno) Start() {
	x.loadStates()
	x.loadButtons()
}

func (x *Xeno) Update() error {
	for _, b := range x.buttons {
		b.Update()
	}
	return nil
}

func (x *Xeno) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	x.drawStateImage(screen)

	for _, b := range x.buttons {
		b.Draw(screen)
	}
}

func (x *Xeno) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // small screen size
}
