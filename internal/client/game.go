package client

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaiaverkvist/tinyecs"
	"github.com/red-letter-day/habbo/internal/habbo"
	"github.com/red-letter-day/habbo/internal/room"
)

type Game struct {
	Engine  *tinyecs.Engine
	systems []System

	cameraPositionX float64
	cameraPositionY float64
}

const testTilemap = `xx*xx
x00000
x00+00
x00000
x00000`

func NewGame() Game {
	e := tinyecs.NewEngine()

	roomEntity := room.RoomEntity{
		Tilemap: habbo.NewRoomTilemap(testTilemap),
	}
	e.AddComponents(roomEntity,
		room.RoomRenderableComponent{},
	)

	var systems []System
	systems = append(systems, &room.RoomRenderSystem{})

	return Game{
		Engine:  &e,
		systems: systems,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	for _, system := range g.systems {
		system.Update(g.Engine)
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	// Write your game's rendering.
	for _, system := range g.systems {
		system.Draw(g.Engine, screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 1, outsideHeight / 1
}
