package habbo

type Room struct {
	name string

	RoomTilemap RoomTilemap
}

const (
	TileTypeEmpty   = '+'
	TileTypeDoorWay = '*'
	TileTypeWall    = 'x'
	TileTypeFloor   = '0'

	TileTypeInvalid = '?'
)

// RoomTile represents a single tile in a room.
type RoomTile struct {
	TileType rune
}

type RoomTilemap struct {
	Door  RoomPosition
	Tiles [][]RoomTile
}

func NewRoomTilemap(tiles string) RoomTilemap {
	return RoomTilemap{
		Tiles: ParseTilemap(tiles),
	}
}

// RoomPosition is used to describe a position within a room.
//
// This diagram describes how the axes are set up.
//
//                                |
//                                |      <-- Z axis
//                                |
//                                |
//                                |
//                                |
//                               --\
//                            --/   ---\
//          Y axis -->     --/          ---\        <-- X axis
//                      --/                 ---\
//                   --/                        ---\
//
type RoomPosition struct {
	X int
	Y int
	Z int
}
