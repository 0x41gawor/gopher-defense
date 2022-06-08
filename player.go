package main

import (
	"github.com/0x41gawor/gopher-defense/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"log"
)

const (
	startX   = 20
	startY   = 20
	velocity = 10
)

type Player struct {
	body     *ebiten.Image
	pos      util.Vector
	velocity float64
}

func NewPlayer() Player {
	img, _, err := ebitenutil.NewImageFromFile("img/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return Player{
		img,
		util.Vector{X: startX, Y: startY},
		velocity,
	}
}

func (p *Player) Update() {
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		p.pos.Y -= p.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		p.pos.Y += p.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		p.pos.X -= p.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		p.pos.X += p.velocity
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.pos.X, p.pos.Y)
	screen.DrawImage(p.body, opt)
}
