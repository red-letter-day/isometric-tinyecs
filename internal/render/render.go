package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/red-letter-day/habbo/internal/camera"
)

func Image(camera camera.Camera, screen *ebiten.Image, image *ebiten.Image, posX float64, posY float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(posX, posY)
	op.GeoM.Translate(-camera.PosX, -camera.PosY)
	screen.DrawImage(image, &op)
}
