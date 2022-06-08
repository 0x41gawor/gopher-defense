package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 1280
	SCREENHEIGHT = 800
	TITLE        = "Gopher Defence"
)

type Game struct {
	player Player
}

func NewGame() Game {
	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle(TITLE)
	var player = NewPlayer()
	return Game{player: player}
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
