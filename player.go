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
	gun      Gun
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
		NewGun(0.2),
	}
}

func (this *Player) Update(dt float64) {
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		this.pos.Y -= this.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		this.pos.Y += this.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		this.pos.X -= this.velocity
	}
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		this.pos.X += this.velocity
	}

	if this.pos.Y > SCREENHEIGHT {
		this.pos.Y = SCREENHEIGHT
	} else if this.pos.Y < 0 {
		this.pos.Y = 0
	}
	if this.pos.X > SCREENWIDTH {
		this.pos.X = SCREENWIDTH
	} else if this.pos.X < 0 {
		this.pos.X = 0
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		targetPos := Vector{X: float64(x), Y: float64(y)}
		this.gun.fire(targetPos, dt)
	}
	this.gun.Update(dt, this.pos)
}

func (this *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(this.pos.X-sizeX/2, this.pos.Y-sizeY/2)
	screen.DrawImage(this.body, opt)
	this.gun.Draw(screen)
}
