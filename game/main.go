package main

import (
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	controllerManager *controllerManager
	entities []*entity
}

func (g *Game) Update() error {
	for _, entity := range g.entities {
		err := entity.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Loop through players and output their inputs
	for _, controller := range g.controllerManager.controllers {
		for input, pressed := range controller.Inputs.KeyPressed {
			if(pressed) {
				ebitenutil.DebugPrint(screen, "player " + strconv.Itoa(controller.Id + 1) + " is pressing key " + input)
			}
		}
	}

	for _, entity := range g.entities {
		entity.onDraw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	controllerManager := newControllerManager()

	// Sleep for a few seconds to allow the controller to connect
	time.Sleep(5 * time.Second)
	
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	entities := []*entity{}

	// Create a viewport for player 1
	// // Viewport should follow player 1 upwards when the player is close to the top of the screen
	// // Viewport should not follow player 1 downwards when the player is close to the bottom of the screen - player can fall off the bottom of the screen

	// Create player 1
	player1 := newPlayerEntity(vector{x: 0, y: 0}, controllerManager)
	entities = append(entities, player1)

	// // Player should constantly jump up and down as long as it's landing on a platform
	// // Player should be able to move left and right with the player 1 inputs
	// // Player should trigger lose condition when below the viewport by a certain amount

	// Viewports and players should be implemented such that we can add more players with their own viewports

	// Create a platform pool (maybe like 100)
	// // Platforms should be able to be recycled when they are off below the lowest viewport by a certain amount
	// // Platforms should be added above the viewport when the highest platform is below a certain distance above the highest view port (ensure we always have platforms above the viewport)

	if err := ebiten.RunGame(&Game{ 
		controllerManager: controllerManager,
		entities: entities,
	}); err != nil {
		log.Fatal(err)
	}
}