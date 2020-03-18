package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/cacarpenter/bugface/model"
	"image/color"
	"log"
)

type bfScene struct{}

// define the bugface game type
func (*bfScene) Type() string { return "Bugface" }

// preload assets
func (*bfScene) Preload() {
	engo.Files.Load("bug.png")
}

func (*bfScene) Setup(u engo.Updater) {

	rb := engo.Input.RegisterButton
	rb(model.MOVE_LEFT, engo.KeyA, engo.KeyArrowLeft)
	rb(model.MOVE_RIGHT, engo.KeyD, engo.KeyArrowRight)
	rb(model.MOVE_UP, engo.KeyW, engo.KeyArrowUp)
	rb(model.MOVE_DOWN, engo.KeyS, engo.KeyArrowDown)
	rb(model.EXIT, engo.KeyEscape)

	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)

	bugfaceSystem := BugfaceSystem{}
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&bugfaceSystem)

	// Retrieve a bugTexture
	bugTexture, err := common.LoadedSprite("bug.png")
	if err != nil {
		log.Println(err)
	}

	game := NewGame(bugTexture)
	bugfaceSystem.BugfaceGame = &game

	// Add it to appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			bugfaceSystem.AddToSystem(sys)
			// sys.Add(&game.Bugface.BasicEntity, &game.Bugface.RenderComponent, &game.Bugface.SpaceComponent)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title:  "Bugface",
		Width:  1024,
		Height: 640,
	}
	engo.Run(opts, &bfScene{})
}
