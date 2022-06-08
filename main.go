package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	var game = NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
