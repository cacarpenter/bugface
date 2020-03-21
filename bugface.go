package main

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
	"log"
)

type bfScene struct{}

// define the bugface game type
func (*bfScene) Type() string { return "Bugface" }

// preload assets
func (*bfScene) Preload() {
	engo.Files.Load(PLAYER_TEXTURE_FILE)
}

func (*bfScene) Setup(u engo.Updater) {

	rb := engo.Input.RegisterButton
	rb(MOVE_LEFT, engo.KeyA, engo.KeyArrowLeft)
	rb(MOVE_RIGHT, engo.KeyD, engo.KeyArrowRight)
	rb(MOVE_UP, engo.KeyW, engo.KeyArrowUp)
	rb(MOVE_DOWN, engo.KeyS, engo.KeyArrowDown)
	rb(EXIT, engo.KeyEscape)

	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)

	bugfaceSystem := BugfaceSystem{}
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&bugfaceSystem)

	// Retrieve a bugTexture
	bugTexture, err := common.LoadedSprite(PLAYER_TEXTURE_FILE)
	if err != nil {
		log.Println(err, " bye")
		engo.Exit()
	}
	fmt.Printf("h = %v, w = %v\n", bugTexture.Height(), bugTexture.Width())

	game := NewGame(bugTexture)
	bugfaceSystem.BugfaceGame = &game

	// Add it to appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			bugfaceSystem.AddToSystem(sys)
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
