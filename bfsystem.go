package main

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/model"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type PlayerEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	model.Bugface
}

type BugfaceSystem struct {
	world *ecs.World

	mouseTracker MouseTracker

	*BugfaceGame
}

func (*BugfaceSystem) Remove(ecs.BasicEntity) {
	fmt.Println("BugfaceSystem.Remove")
}

func (bfs *BugfaceSystem) Update(dt float32) {
	if engo.Input.Button(model.MOVE_LEFT).JustPressed() {
		fmt.Println("LEFT Pressed")
		bfs.RotatePlayer(-5)
	}
	if engo.Input.Button(model.MOVE_RIGHT).JustPressed() {
		fmt.Println("RIGHT Pressed")
		bfs.RotatePlayer(5)
	}
}

func (bfs *BugfaceSystem) AddToSystem(sys *common.RenderSystem) {
	sys.Add(&bfs.player.BasicEntity, &bfs.player.RenderComponent, &bfs.player.SpaceComponent)
}
