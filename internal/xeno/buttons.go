package xeno

import (
	"image/color"
	"math/rand/v2"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type button struct {
	assignedKey ebiten.Key
	label       string
	x, y        int
	width       int
	height      int
	onClick     func()
	hovered     bool
}

func (x *Xeno) loadButtons() {
	x.buttons = []*button{
		{
			label: "Change",
			x:     10, y: 10, width: 80, height: 30,
			onClick: func() {
				nextState := StateNames[rand.IntN(len(StateNames))]
				x.currentState = nextState
			},
			assignedKey: ebiten.KeyEnter,
		},
		{
			label: "Exit",
			x:     100, y: 10, width: 80, height: 30,
			onClick: func() {
				os.Exit(0)
			},
			assignedKey: ebiten.KeyEscape,
		},
	}
}

func (b *button) Update() {
	x, y := ebiten.CursorPosition()
	b.hovered = x >= b.x && x <= b.x+b.width && y >= b.y && y <= b.y+b.height

	if b.hovered && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.onClick()
	}

	if ebiten.IsKeyPressed(b.assignedKey) {
		b.onClick()
	}
}

func (b *button) Draw(screen *ebiten.Image) {
	col := color.RGBA{100, 100, 100, 255} // default
	if b.hovered {
		col = color.RGBA{150, 150, 150, 255}
	}
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), col, false)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(b.x+10), float64(b.y+20)) // 20 is baseline offset
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, b.label, text.NewGoXFace(basicfont.Face7x13), op)
}
