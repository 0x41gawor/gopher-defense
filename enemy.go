package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Enemy struct {
	bodyBlack, bodyWhite, bodyGreen *ebiten.Image
	pos                             Vector
	sizeX, sizeY                    int
	ms                              float64
	health, maxHealth               float32
}

func NewEnemy(sizeX int, sizeY int, ms float64, health float32, maxHealth float32) Enemy {
	imageWhite := ebiten.NewImage(sizeX+2, sizeY+2)
	imageWhite.Fill(color.White)
	imageBlack := ebiten.NewImage(sizeX, sizeY)
	imageBlack.Fill(color.Black)
	imageGreen := ebiten.NewImage(sizeX, sizeY)
	imageGreen.Fill(color.RGBA{G: 0xff, A: 0xff})

	return Enemy{
		bodyBlack: imageBlack,
		bodyWhite: imageWhite,
		bodyGreen: imageGreen,
		pos:       Vector{X: 50, Y: 50},
		sizeX:     sizeX,
		sizeY:     sizeY,
		ms:        ms,
		health:    health,
		maxHealth: maxHealth,
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(e.pos.X, e.pos.Y)
	screen.DrawImage(e.bodyWhite, opt)
	opt.GeoM.Translate(1, 1)
	screen.DrawImage(e.bodyBlack, opt)
	opt.GeoM.Translate(0, float64((1-e.health/e.maxHealth)*float32(sizeY)))
	screen.DrawImage(e.bodyGreen, opt)
}

func (e *Enemy) Update(playerPos Vector) {
	// update moving dir
	var msX, msY float64
	if playerPos.X < e.pos.X {
		msX = e.ms * -1
	} else {
		msX = e.ms * 1
	}
	if playerPos.Y < e.pos.Y {
		msY = e.ms * -1
	} else {
		msY = e.ms * 1
	}
	// move
	e.pos.X += msX
	e.pos.Y += msY

	// update health image
	e.bodyGreen = ebiten.NewImage(sizeX+2, int(float32(sizeY)*e.health/e.maxHealth))
	e.bodyGreen.Fill(color.RGBA{G: 0xff, A: 0xff})
}
