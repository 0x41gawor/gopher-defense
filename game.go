package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 1920 / 2
	SCREENHEIGHT = 1080 / 2
	TITLE        = "Gopher Defence"
)

type Game struct {
	player   Player
	zombie   Zombie
	skeleton Skeleton
	time     float64
}

func NewGame() Game {
	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle(TITLE)
	var player = NewPlayer()
	var zombie = NewZombie(30, 50, 1, 80, 100)
	var skeleton = NewSkeleton(40, 60, 0.5, 40, 200)
	return Game{player: player, zombie: zombie, skeleton: skeleton}
}

func (this *Game) Update() error {
	dt := 1.0 / 60.0

	this.player.Update(dt)
	this.zombie.Update(this.player.pos)
	this.skeleton.Update(dt, this.player.pos)
	return nil
}

func (this *Game) Draw(screen *ebiten.Image) {
	this.player.Draw(screen)
	this.zombie.Draw(screen)
	this.skeleton.Draw(screen)
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
