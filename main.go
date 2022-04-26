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
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
