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
	enemy  Enemy
}

func NewGame() Game {
	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle(TITLE)
	var player = NewPlayer()
	var enemy = NewEnemy(30, 50, 1, 80, 100)
	return Game{player: player, enemy: enemy}
}

func (g *Game) Update() error {
	g.player.Update()
	g.enemy.Update(g.player.pos)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.enemy.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}
