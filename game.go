package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/model"
)

type BugfaceGame struct {
	BugTexture *common.Texture
	player model.Bugface
}

func NewGame(playerTexture *common.Texture) BugfaceGame {
	player := model.Bugface{BasicEntity: ecs.NewBasic()}

	// Initialize the components, set scale to 8x
	player.RenderComponent = common.RenderComponent{
		Drawable: playerTexture,
		Scale:    engo.Point{0.1, 0.1},
	}
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{204, 190},
		Width:    playerTexture.Width() * player.RenderComponent.Scale.X,
		Height:   playerTexture.Height() * player.RenderComponent.Scale.Y,
		Rotation: 30,
	}

	return BugfaceGame{playerTexture, player}
}

func (bfs *BugfaceSystem) RotatePlayer(f float32) {
	bfs.player.Rotation = bfs.player.Rotation + f

}
