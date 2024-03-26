package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	ControllerManager *ControllerManager
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Loop through players and output their inputs
	for _, controller := range g.ControllerManager.controllers {
		for input, pressed := range controller.Inputs.KeyPressed {
			if(pressed) {
				ebitenutil.DebugPrint(screen, "player " + strconv.Itoa(controller.Id + 1) + " is pressing key " + input)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	controllerManager := NewControllerManager()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// Create a viewport for player 1
	// // Viewport should follow player 1 upwards when the player is close to the top of the screen
	// // Viewport should not follow player 1 downwards when the player is close to the bottom of the screen - player can fall off the bottom of the screen

	// Create player 1
	// // Player should constantly jump up and down as long as it's landing on a platform
	// // Player should be able to move left and right with the player 1 inputs
	// // Player should trigger lose condition when below the viewport by a certain amount

	// Viewports and players should be implemented such that we can add more players with their own viewports

	// Create a platform pool (maybe like 100)
	// // Platforms should be able to be recycled when they are off below the lowest viewport by a certain amount
	// // Platforms should be added above the viewport when the highest platform is below a certain distance above the highest view port (ensure we always have platforms above the viewport)

	if err := ebiten.RunGame(&Game{ ControllerManager: controllerManager }); err != nil {
		log.Fatal(err)
	}
}