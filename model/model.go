package model

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type Direction byte

const (
	Up Direction = iota
	Down
	Left
	Right
)

type BlockType byte

// order here matches the old values
const (
	Clear BlockType = iota
	Wall
	RegularPellet
	PowerPellet
	Trail
	HorizontalDoor
	VerticalDoor
)

type MazePosition struct {
	X, Y int
}

type Bugface struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent

	Direction
	MazePosition
	Invincible bool
}

type Enemy struct {
	MazePosition
	Direction
	PrevDir Direction
}

type Board struct {
	//Maze          [60][40]BlockType
	// odd shaped board
	Maze          [40][60]BlockType
	StartingPoint MazePosition
}
