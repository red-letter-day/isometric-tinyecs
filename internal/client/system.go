package client

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaiaverkvist/tinyecs"
)

type System interface {
	Update(engine *tinyecs.Engine)
	Draw(engine *tinyecs.Engine, screen *ebiten.Image)
}
