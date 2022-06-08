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
	sizeX    = 28
	sizeY    = 50
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

	if p.pos.Y > SCREENHEIGHT {
		p.pos.Y = SCREENHEIGHT
	} else if p.pos.Y < 0 {
		p.pos.Y = 0
	}
	if p.pos.X > SCREENWIDTH {
		p.pos.X = SCREENWIDTH
	} else if p.pos.X < 0 {
		p.pos.X = 0
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.pos.X-sizeX/2, p.pos.Y-sizeY/2)
	screen.DrawImage(p.body, opt)
}
