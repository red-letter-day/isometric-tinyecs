package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/labstack/gommon/log"
	_ "image/gif"
	_ "image/png"
	"strings"
)

var (
	assets map[string]*ebiten.Image
)

func Register(paths ...string) {
	if assets == nil {
		assets = make(map[string]*ebiten.Image)
	}

	for _, path := range paths {
		img, err := loadImage(path)
		if err != nil {
			log.Error("Failed to register image at path: ", path, ": ", err)
		}

		// Clean out the things we don't want in our path.
		s := strings.ReplaceAll(path, ".png", "")
		s = strings.ReplaceAll(s, "assets/", "")
		cleanPath := s

		// Register the asset
		assets[cleanPath] = img

		log.Info("Registering asset for path ", path, " as <", cleanPath, ">")
	}
}

func Get(path string) *ebiten.Image {
	img := assets[path]
	if img == nil {
		log.Error("Unrecognized image path: ", path)
	}

	return img
}

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	return img, nil
}
