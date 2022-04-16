package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/labstack/gommon/log"
	"github.com/red-letter-day/habbo/internal/client"
)

func main() {
	log.SetHeader("${time_rfc3339}  ${level}  <${short_file}:${line}>")

	game := client.NewGame()
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Habbo Client")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
