package terrain

import (
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/bf"
)

const (
	PILLAR_OPEN        = 0
	PILLAR_OPEN_BRIDGE = 1
	PILLAR_CLOSED      = 2
	//PILLAR_CLOSED_BRIDGE = 3
	EASTER_ISLAND_HEAD = 3
	PILLAR_GAP         = 4
	WATER_HOLE_DARK_1  = 5
	WATER_HOLE_DARK_2  = 6
	WATER_DARK_1       = 7
	WATER_DARK_2       = 8
)

var easterHead = common.SpriteRegion{
	engo.Point{484, 387}, 59, 87,
}

var spriteRegions = []common.SpriteRegion{
	{
		engo.Point{0, 0}, 95, 159,
	},
	{
		engo.Point{93, 0}, 67, 125,
	},
	{
		engo.Point{160, 0}, 95, 160,
	},
	easterHead,
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

func (t *Terrain) TrySome(i int) common.Texture {
	return t.spritesheet.Cell(i)
}
