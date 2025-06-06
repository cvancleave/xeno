package xeno

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type state struct {
	name  string
	image *ebiten.Image
}

var StateNames = []string{"bleh", "blink", "idle", "meh", "sad", "surprise"}

func (x *Xeno) loadStates() {
	for _, name := range StateNames {
		img, err := loadImage(fmt.Sprintf("./data/png/%s.png", name))
		if err != nil {
			fmt.Println("failed to load image for state:", name)
			continue
		}

		x.states[name] = state{
			name:  name,
			image: img,
		}
	}
}

func (x *Xeno) drawStateImage(screen *ebiten.Image) {
	s := x.states[x.currentState]

	screenHeight := screen.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(screenHeight-120))
	screen.DrawImage(s.image, op)
}
