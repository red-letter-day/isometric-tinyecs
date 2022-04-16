package room

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kaiaverkvist/tinyecs"
	"github.com/red-letter-day/habbo/internal/camera"
	"github.com/red-letter-day/habbo/internal/client/assets"
	"github.com/red-letter-day/habbo/internal/habbo"
	"github.com/red-letter-day/habbo/internal/render"
)

type RoomEntity struct {
	tinyecs.Entity

	Tilemap habbo.RoomTilemap
}

type RoomRenderableComponent struct {
	userHoveredTile habbo.RoomPosition
}

type RoomRenderSystem struct{}

func (s *RoomRenderSystem) Update(engine *tinyecs.Engine) {
	tinyecs.Each[RoomRenderableComponent](engine, func(id uint64, component RoomRenderableComponent) {
		cx, cy := ebiten.CursorPosition()
		coordX, coordY := s.screenToMap(cx, cy, 64, 32, camera.Cam)
		component.userHoveredTile = habbo.RoomPosition{X: coordX, Y: coordY}

		tinyecs.Set(engine, id, component)
	})
}

func (s *RoomRenderSystem) Draw(engine *tinyecs.Engine, screen *ebiten.Image) {
	tinyecs.EachEntity[RoomEntity, RoomRenderableComponent](engine, func(entity RoomEntity, obj RoomRenderableComponent) {
		for x, tileRow := range entity.Tilemap.Tiles {
			for y, tile := range tileRow {

				var tileImage *ebiten.Image
				switch tile.TileType {
				case habbo.TileTypeFloor:
					tileImage = assets.Get("room/floor_tile")
					s.drawTile(screen, float64(x), float64(y), tileImage)
					ebitenutil.DebugPrint(screen, fmt.Sprintf("%d, %d", obj.userHoveredTile.X, obj.userHoveredTile.Y))
					ebitenutil.DebugPrint(screen, fmt.Sprintf("\n%.2f, %.2f", camera.Cam.PosX, camera.Cam.PosY))

					cx, cy := ebiten.CursorPosition()
					ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n%d, %d", cx, cy))
					if obj.userHoveredTile.X == x && obj.userHoveredTile.Y == y {
						s.drawHighlight(screen, float64(x), float64(y))
					}
				case habbo.TileTypeWall:

					left := false

					if x == 0 || (x == 0 && y == 0) {
						left = true
					}
					s.drawWall(screen, float64(x), float64(y), left)
				case habbo.TileTypeDoorWay:
					s.drawDoor(screen, float64(x), float64(y))
				}
			}
		}
	})
}

func (s *RoomRenderSystem) cartesianToIsometric(tileSize int, x float64, y float64) (float64, float64) {
	ix := (x - y) * float64(tileSize/2)
	iy := (x + y) * float64(tileSize/4)
	return ix, iy
}

func (s *RoomRenderSystem) screenToMap(posX int, posY int, tileWidth int, tileHeight int, cam camera.Camera) (int, int) {
	/*posX -= tileWidth / 2
	posX += int(cam.PosX)
	posY += int(cam.PosY)
	tileHeight /= 2
	tileWidth /= 2
	x := (posX/tileWidth + posY/tileHeight) / 2
	y := (posY/tileHeight - (posX / tileWidth)) / 2
	*/
	posX += int(cam.PosX)
	posY += int(cam.PosY)
	pXOffset := -16
	pYOffset := -16

	x := (((posX - tileHeight) - pXOffset) / tileWidth) + ((posY - pYOffset) / tileHeight)
	y := (((posY - pYOffset) / tileHeight) - ((posX-tileHeight)-pXOffset)/tileWidth)

	return x, y
}

func (s *RoomRenderSystem) drawTile(screen *ebiten.Image, x float64, y float64, tile *ebiten.Image) {
	coordX, coordY := s.cartesianToIsometric(64, x, y)

	render.Image(camera.Cam, screen, tile, coordX, coordY)
}

func (s *RoomRenderSystem) drawHighlight(screen *ebiten.Image, x float64, y float64) {
	coordX, coordY := s.cartesianToIsometric(64, x, y)

	tile := assets.Get("room/tile_hover")
	render.Image(camera.Cam, screen, tile, coordX, coordY-5)
}

func (s *RoomRenderSystem) drawWall(screen *ebiten.Image, x float64, y float64, left bool) {
	coordX, coordY := s.cartesianToIsometric(64, x, y)

	var tile *ebiten.Image

	if left {
		tile = assets.Get("room/wall_l")
		coordX -= 8
		coordY -= 91
	} else {
		tile = assets.Get("room/wall_r")
		coordY -= 107
	}
	render.Image(camera.Cam, screen, tile, coordX, coordY)
}

func (s *RoomRenderSystem) drawDoor(screen *ebiten.Image, x float64, y float64) {
	coordX, coordY := s.cartesianToIsometric(64, x, y)

	var tileDoor = assets.Get("room/door")
	var tileFloor = assets.Get("room/door_floor")

	render.Image(camera.Cam, screen, tileDoor, coordX, coordY-91)
	render.Image(camera.Cam, screen, tileFloor, coordX, coordY+16)
}
