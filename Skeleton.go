package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Skeleton struct {
	zombie Zombie
	gun    Gun
}

func NewSkeleton(sizeX int, sizeY int, ms float64, health float32, maxHealth float32) Skeleton {
	imageFrame := ebiten.NewImage(sizeX+2, sizeY+2)
	imageFrame.Fill(color.White)
	imageBlack := ebiten.NewImage(sizeX, sizeY)
	imageBlack.Fill(color.Black)
	imageHealth := ebiten.NewImage(sizeX, sizeY)
	imageHealth.Fill(color.RGBA{R: 0xff, A: 0xff})

	return Skeleton{
		Zombie{
			bodyBlack:  imageBlack,
			bodyFrame:  imageFrame,
			bodyHealth: imageHealth,
			pos:        Vector{X: 50, Y: 50},
			sizeX:      sizeX,
			sizeY:      sizeY,
			ms:         ms,
			health:     health,
			maxHealth:  maxHealth,
		},
		NewGun(2.0),
	}
}

func (this *Skeleton) Draw(screen *ebiten.Image) {
	this.zombie.Draw(screen)
	this.gun.Draw(screen)
}

func (this *Skeleton) Update(dt float64, playerPos Vector) {
	// update pos
	this.zombie.Update(playerPos)
	// update shooting
	this.gun.fire(playerPos, dt)
	this.gun.Update(dt, this.zombie.pos)
}
