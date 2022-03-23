package habbo_test

import (
	"github.com/red-letter-day/habbo/internal/habbo"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Tilemap_1 = `xxxxx
x0000
x0000
x0000`

	Tilemap_2 = `xxxxx
x0000
x0000
x0000
x0000
x0000`
)

func Test_ParseTilemap(t *testing.T) {
	tilemap := habbo.ParseTilemap(Tilemap_1)
	assert.Len(t, tilemap, 4)

	tilemap2 := habbo.ParseTilemap(Tilemap_2)
	assert.Len(t, tilemap2, 6)
}
