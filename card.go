package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Point struct {
	X int
	Y int
}

type Card struct {
	Val     int 
	Position Point
	Shown    bool
}

func (this *Card) Draw(screen *ebiten.Image) {
	red := color.RGBA{R: 180, G: 120, B: 120, A: 255}
	grey := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	if this.Shown {

		ebitenutil.DrawRect(screen, float64(this.Position.X), float64(this.Position.Y), 50, 75, red)	

		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("%d", this.Val),
			int(this.Position.X+20),
			int(this.Position.Y+25),
		)
	} else {
		ebitenutil.DrawRect(screen, float64(this.Position.X), float64(this.Position.Y), 50, 75, grey)
	}
}

func (this *Card) Update(g *Game) {
	mx, my := ebiten.CursorPosition()

	if this.Position.X <= mx && this.Position.X+50 >= mx && this.Position.Y <= my && this.Position.Y+75 >= my {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			this.Flip(g)
		}
	}
}

func (this *Card) Flip(g *Game) {
	if !this.Shown {
		this.Shown = true
		g.FlipCard(this)
	}
}
