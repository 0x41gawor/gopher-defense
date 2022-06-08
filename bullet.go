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
	pos  Vector
	msX  float64 // ms - movement speed
	msY  float64
	body *ebiten.Image
}

func NewBullet(startPos Vector, aimPos Vector) *Bullet {
	image := ebiten.NewImage(width, height)
	image.Fill(color.White)

	a := math.Abs(startPos.Y-aimPos.Y) / math.Abs(startPos.X-aimPos.X)
	angle := math.Atan(a)

	msX := 10 * math.Cos(angle)
	msY := 10 * math.Sin(angle)

	if aimPos.X <= startPos.X {
		msX *= -1
		if aimPos.Y < startPos.Y {
			msY *= -1
		}
	} else {
		if aimPos.Y < startPos.Y {
			msY *= -1
		}
	}

	return &Bullet{
		startPos,
		msX,
		msY,
		image,
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
