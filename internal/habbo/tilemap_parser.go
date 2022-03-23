package habbo

import (
	"strings"
)

// ParseTilemap takes in a
func ParseTilemap(tiles string) [][]RoomTile {
	var tilemap [][]RoomTile

	lines := strings.Split(tiles, "\n")
	for _, line := range lines {
		slice := make([]RoomTile, len([]rune(line)))
		for ci, character := range []rune(line) {
			slice[ci] = RoomTile{TileType: character}
		}
		tilemap = append(tilemap, slice)
	}

	return tilemap
}
