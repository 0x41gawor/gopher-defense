package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	SCREENWIDTH  = 1280
	SCREENHEIGHT = 800
	TITLE        = "Gopher Defence"
)

func main() {
	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle(TITLE)

	var game = NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
