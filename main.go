package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cards              map[Point]*Card
	flipped            []*Card
	numFlipped         int
	flipTimer          float64
	previousUpdateTime time.Time
}

func (g *Game) Update() error {
	currentTime := time.Now()
	deltaTime := currentTime.Sub(g.previousUpdateTime).Seconds()
	g.previousUpdateTime = currentTime
	if g.flipTimer > 0 {
		g.flipTimer -= deltaTime
	} else {
		if g.numFlipped == 2 {
			if g.flipped[0].Val == g.flipped[1].Val {
				g.cards[g.flipped[0].Position] = nil
				g.cards[g.flipped[1].Position] = nil
			} else {
				g.flipped[0].Shown = false
				g.flipped[1].Shown = false
			}
			g.flipped[0] = nil
			g.flipped[1] = nil
			g.numFlipped = 0
		}

	}
	for i := range g.cards {
		if g.cards[i] != nil {
			g.cards[i].Update(g)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, card := range g.cards {
		if card != nil {
			card.Draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) FlipCard(card *Card) {
	switch g.numFlipped {
	case 0:
		g.flipped[0] = card
		g.numFlipped++
	case 1:
		g.flipped[1] = card
		g.numFlipped++
		g.flipTimer = 0.8
	default:
		card.Shown = false
	}
}

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Hello, World!")

	numbers := make([]int, 30)
	for i := 0; i < 30; i++ {
		numbers[i] = i + 1
	}
	pool := append([]int{}, numbers...)
	pool = append(pool, numbers...)
	rand.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})

	cards := make(map[Point]*Card)
	for i := 0; i < 60; i++ {
		pos := Point{X: 150 + i%10*75, Y: 50 + i/10*100}
		cards[pos] = &Card{
			Val:      pool[i],
			Position: pos,
			Shown:    false,
		}
	}

	if err := ebiten.RunGame(&Game{cards: cards, flipped: make([]*Card, 2), flipTimer: 0.0}); err != nil {
		log.Fatal(err)
	}
}
