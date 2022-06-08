package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Gun struct {
	pos     Vector
	bullets []*Bullet
}

func NewGun() Gun {
	return Gun{pos: Vector{}}
}

func (g *Gun) Update(pos Vector) {
	g.pos = pos
	for _, bullet := range g.bullets {
		bullet.Update()
	}
}

func (g *Gun) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.pos.X, g.pos.Y)
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
}
