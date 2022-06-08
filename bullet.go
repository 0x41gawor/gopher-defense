package main

import (
	"github.com/0x41gawor/gopher-defense/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Vector = util.Vector

const (
	width  = 5
	height = 5
)

type Bullet struct {
	body *ebiten.Image
	pos  Vector
	msX  float64 // ms - movement speed
	msY  float64
}

func NewBullet(startPos Vector, targetPos Vector) *Bullet {
	image := ebiten.NewImage(width, height)
	image.Fill(color.White)

	a := math.Abs(startPos.Y-targetPos.Y) / math.Abs(startPos.X-targetPos.X)
	angle := math.Atan(a)

	msX := 10 * math.Cos(angle)
	msY := 10 * math.Sin(angle)

	if targetPos.X <= startPos.X {
		msX *= -1
		if targetPos.Y < startPos.Y {
			msY *= -1
		}
	} else {
		if targetPos.Y < startPos.Y {
			msY *= -1
		}
	}

	return &Bullet{
		image,
		startPos,
		msX,
		msY,
	}
}

func (b *Bullet) Update() {
	b.pos.X += b.msX
	b.pos.Y += b.msY
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.pos.X, b.pos.Y)
	screen.DrawImage(b.body, opt)
}
