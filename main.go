package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	cards []Card
}

func (g *Game) Update() error {
	for i := range g.cards {
		g.cards[i].Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for _, card := range g.cards {
		card.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Hello, World!")

	var cards []Card
	for i := 0; i < 60; i++ {
		card := Card{
			Val: 'a',
			Position: Point{X: 150 + i % 10 * 75, Y: 50 + i / 10 * 100},
			Shown: false,
		}
		cards = append(cards, card)
	}

	if err := ebiten.RunGame(&Game{cards: cards}); err != nil {
		log.Fatal(err)
	}
}
