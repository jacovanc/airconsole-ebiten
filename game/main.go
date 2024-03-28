package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	levelWidth  = 250
)

type Game struct{
	controllerManager *controllerManager
	entities []*entity
	cameras []*entity
}

func (g *Game) Update() error {
	// For easy development we can use the arrow keys to control player 1 - but this essentially disables the airconsole controller
	overwritePlayer1ControllerWithArrowKeys(g.controllerManager)

	// Check all collisions
	checkCollisions(g.entities)

	// Update cameras
	for _, camera := range g.cameras {
		err := camera.update()
		if err != nil {
			return err
		}
	}

	// Update all entities
	for _, entity := range g.entities {
		err := entity.update()
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

	// Draw entities in a camera view
	for _, camera := range g.cameras {
		// Draw the camera view
		cameraComponent := camera.getComponent("cameraComponent").(*cameraComponent)
		camera.draw(screen, nil)
		for _, entity := range g.entities {
			if cameraComponent != nil && cameraComponent.isInView(entity) {
				entity.draw(screen, cameraComponent)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	controllerManager := newControllerManager()

	// Create a controller for player 1 so that we can overwrite the controller with arrow keys, even if the controller is not connected
	controllerManager.addController(0) // Dev only

	// Sleep for a few seconds to allow the controller to connect
	time.Sleep(5 * time.Second)
	
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	entities := []*entity{}
	cameras := []*entity{}

	// Create a viewport for player 1
	// // Viewport should follow player 1 upwards when the player is close to the top of the screen
	// // Viewport should not follow player 1 downwards when the player is close to the bottom of the screen - player can fall off the bottom of the screen

	// Create player 1
	player1 := newPlayerEntity(vector{x: levelWidth / 2, y: 400}, controllerManager)
	entities = append(entities, player1)

	// Platform guaranteed to be below the player on spawn
	platform := newPlatformEntity(vector{x: 100, y: 500})
	entities = append(entities, platform)

	// Create a platform pool (maybe like 100)
	previousXPos := 0
	for i := 0; i < 100; i++ {
		xPos := rand.Intn(levelWidth - platformWidth) // Don't place a platform outside the level width
		// Ensure that xPos is within 200 pixels previous platform
		if(i > 0) {
			if(xPos < previousXPos - 150) {
				xPos = previousXPos - 150
			}
			if(xPos > previousXPos + 150) {
				xPos = previousXPos + 150
			}
		}

		yPos := -(i * 100) + 1000 // Offset the platforms by 1000 so that they start below the player not above

		// Limit xPos to be between 0 and 640
		if(xPos < 0) {
			xPos = 0
		}
		if(xPos > 640) {
			xPos = 640
		}

		platform := newPlatformEntity(vector{x: float64(xPos), y: float64(yPos)})
		entities = append(entities, platform)
	}

	// // Platforms should be able to be recycled when they are off below the lowest viewport by a certain amount
	// // Platforms should be added above the viewport when the highest platform is below a certain distance above the highest view port (ensure we always have platforms above the viewport)

	player1Camera := newCameraEntity(player1)
	cameras = append(cameras, player1Camera)

	if err := ebiten.RunGame(&Game{ 
		controllerManager: controllerManager,
		entities: entities,
		cameras: cameras,
	}); err != nil {
		log.Fatal(err)
	}
}

func overwritePlayer1ControllerWithArrowKeys(controllerManager *controllerManager) {
	controller := controllerManager.getController(0)
	if(controller == nil) {
		return
	}
	if(ebiten.IsKeyPressed(ebiten.KeyLeft)) {
		controller.Inputs.KeyPressed["left"] = true
	} else {
		controller.Inputs.KeyPressed["left"] = false
	}
	if(ebiten.IsKeyPressed(ebiten.KeyRight)) {
		controller.Inputs.KeyPressed["right"] = true
	} else {
		controller.Inputs.KeyPressed["right"] = false
	}
}