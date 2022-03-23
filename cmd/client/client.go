package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/red-letter-day/habbo/internal/client"
	"log"
)

func main() {
	game := client.NewGame()
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Habbo Client")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
