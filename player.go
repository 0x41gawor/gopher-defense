package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Player struct {
	body *ebiten.Image
}

func NewPlayer() Player {
	img, _, err := ebitenutil.NewImageFromFile("img/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return Player{
		img,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(p.body, opt)
}
