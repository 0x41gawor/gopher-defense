package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Gun struct {
	pos              Vector
	bullets          []*Bullet
	attackSpeed      float64
	timeFromLastShot float64
}

func (this *Gun) fire(targetPos Vector, dt float64) {
	if this.timeFromLastShot > this.attackSpeed {
		this.bullets = append(this.bullets, NewBullet(this.pos, targetPos))
		this.timeFromLastShot = 0.0
	}
}

func NewGun(attackSpeed float64) Gun {
	return Gun{pos: Vector{}, attackSpeed: attackSpeed}
}

func (this *Gun) Update(pos Vector, dt float64) {
	this.pos = pos
	for _, bullet := range this.bullets {
		bullet.Update()
	}
	this.timeFromLastShot += dt
}

func (this *Gun) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(this.pos.X, this.pos.Y)
	for _, bullet := range this.bullets {
		bullet.Draw(screen)
	}
}
