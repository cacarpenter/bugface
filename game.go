package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/bf"
	"github.com/cacarpenter/bugface/model"
	"github.com/cacarpenter/bugface/terrain"
)

type BugfaceGame struct {
	//BugTexture *common.Texture
	player *model.Bugface

	Points       int
	CurrentLevel int
}

func NewGame() BugfaceGame {
	player := &model.Bugface{BasicEntity: ecs.NewBasic()}

	terr := terrain.NewTerrain()
	playerTexture := terr.TrySome(terrain.EASTER_ISLAND_HEAD)

	// Initialize the components, set scale to 8x
	player.RenderComponent = common.RenderComponent{
		Drawable: playerTexture,
		Scale:    engo.Point{X: bf.DEFAULT_SCALE, Y: bf.DEFAULT_SCALE},
	}
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{204, 190},
		Width:    playerTexture.Width() * player.RenderComponent.Scale.X,
		Height:   playerTexture.Height() * player.RenderComponent.Scale.Y,
		Rotation: 0,
	}

	return BugfaceGame{ player, 0, 0}
}

func (bfs *BugfaceSystem) RotatePlayer(f float32) {
	bfs.player.Rotation = bfs.player.Rotation + f
}

func (bfs *BugfaceSystem) MovePlayer(f float32) {
	bfs.player.SpaceComponent.Position.Y = bfs.player.SpaceComponent.Position.Y + f
}
