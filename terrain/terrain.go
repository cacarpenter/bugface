package terrain

import (
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/bf"
)

var spriteRegions = []common.SpriteRegion{
	{
		engo.Point{0, 0}, 95, 160,
	},
	{
		engo.Point{95, 0}, 65, 125,
	},
	{
		engo.Point{160, 0}, 95, 160,
	},
	{
		engo.Point{160, 0}, 93, 160,
	},
}

type Terrain struct {
	spritesheet *common.Spritesheet
}

func NewTerrain() Terrain {
	terrainSheet := common.NewAsymmetricSpritesheetFromFile(bf.TERRAIN_TEXTURE_SHEET_FILE, spriteRegions)
	return Terrain{terrainSheet}
}

func (t *Terrain) TryMe() common.Texture {
	return t.spritesheet.Cell(2)
}
