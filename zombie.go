package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Zombie struct {
	bodyBlack, bodyFrame, bodyHealth *ebiten.Image
	pos                              Vector
	sizeX, sizeY                     int
	ms                               float64
	health, maxHealth                float32
}

func NewZombie(sizeX int, sizeY int, ms float64, health float32, maxHealth float32) Zombie {
	imageFrame := ebiten.NewImage(sizeX+2, sizeY+2)
	imageFrame.Fill(color.White)
	imageBlack := ebiten.NewImage(sizeX, sizeY)
	imageBlack.Fill(color.Black)
	imageHealth := ebiten.NewImage(sizeX, sizeY)
	imageHealth.Fill(color.RGBA{R: 0xff, A: 0xff})

	return Zombie{
		bodyBlack:  imageBlack,
		bodyFrame:  imageFrame,
		bodyHealth: imageHealth,
		pos:        Vector{X: 50, Y: 50},
		sizeX:      sizeX,
		sizeY:      sizeY,
		ms:         ms,
		health:     health,
		maxHealth:  maxHealth,
	}
}

func (this *Zombie) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(this.pos.X, this.pos.Y)
	screen.DrawImage(this.bodyFrame, opt)
	opt.GeoM.Translate(1, 1)
	screen.DrawImage(this.bodyBlack, opt)
	opt.GeoM.Translate(0, float64((1-this.health/this.maxHealth)*float32(this.sizeY)))
	screen.DrawImage(this.bodyHealth, opt)
}

func (this *Zombie) Update(playerPos Vector) {
	// update moving dir
	var msX, msY float64
	if playerPos.X < this.pos.X {
		msX = this.ms * -1
	} else {
		msX = this.ms * 1
	}
	if playerPos.Y < this.pos.Y {
		msY = this.ms * -1
	} else {
		msY = this.ms * 1
	}
	// move
	this.pos.X += msX
	this.pos.Y += msY

	// update health image
	this.bodyHealth = ebiten.NewImage(this.sizeX, int(float32(this.sizeY)*this.health/this.maxHealth))
	this.bodyHealth.Fill(color.RGBA{R: 0xff, A: 0xff})
}
