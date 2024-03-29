package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/jacovanc/airconsole-ebiten/game/collisions"
	"github.com/jacovanc/airconsole-ebiten/game/components"
	"github.com/jacovanc/airconsole-ebiten/game/controllers"
	"github.com/jacovanc/airconsole-ebiten/game/entities"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"

	_ "net/http/pprof"
)


const (
	levelWidth  = 250
	platformWidth = 32
	platformHeight = 10
)

type Game struct{
	controllerManager *controllers.ControllerManager
	entitiesArray []interfaces.Entity
	camerasArray []interfaces.Entity

	// Benchmarking
	updateAccumulatedTime time.Duration
    drawAccumulatedTime   time.Duration
    frameCount            int
	frameThreshold		 int
}

func (g *Game) Update() error {
	// Benchmarking
	startTime := time.Now()

	// For easy development we can use the arrow keys to control player 1 - but this essentially disables the airconsole controller
	overwritePlayer1ControllerWithArrowKeys(g.controllerManager)

	// Check all collisions
	collisions.CheckCollisions(g.entitiesArray)

	// Update cameras
	for _, camera := range g.camerasArray {
		err := camera.Update()
		if err != nil {
			return err
		}
	}

	// Update all entities
	for _, entity := range g.entitiesArray {
		err := entity.Update()
		if err != nil {
			return err
		}
	}

	// Benchmarking
	elapsedTime := time.Since(startTime)
	g.updateAccumulatedTime += elapsedTime
	g.frameCount++

	if g.frameCount >= g.frameThreshold {
		log.Printf("Total update accumulated time: %v ms\n", g.updateAccumulatedTime.Milliseconds())
		g.updateAccumulatedTime = 0 // Reset the accumulated time
		g.frameCount = 0 // Reset the frame count
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Benchmarking
	startTime := time.Now()

	// Loop through players and output their inputs
	for _, controller := range g.controllerManager.Controllers {
		for input, pressed := range controller.Inputs.KeyPressed {
			if(pressed) {
				ebitenutil.DebugPrint(screen, "player " + strconv.Itoa(controller.Id + 1) + " is pressing key " + input)
			}
		}
	}

	// Draw entities in a camera view
	for _, camera := range g.camerasArray {
		// Draw the camera view
		cameraComponent := camera.GetComponent("cameraComponent").(*components.CameraComponent)
		camera.Draw(screen, nil)
		for _, entity := range g.entitiesArray {
			if cameraComponent != nil && cameraComponent.IsInView(entity) {
				entity.Draw(screen, cameraComponent)
			}
		}
	}

	// Benchmarking
	elapsedTime := time.Since(startTime)
	g.drawAccumulatedTime += elapsedTime

	if g.frameCount >= g.frameThreshold {
		log.Printf("Total draw accumulated time: %v ms\n", g.drawAccumulatedTime.Milliseconds())
		g.drawAccumulatedTime = 0 // Reset the accumulated time
		// No need to reset g.frameCount here, as it's done in Update
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	controllerManager := controllers.NewControllerManager()

	// Create a controller for player 1 so that we can overwrite the controller with arrow keys, even if the controller is not connected
	controllerManager.AddController(0) // Dev only

	// Sleep for a few seconds to allow the controller to connect
	time.Sleep(5 * time.Second)
	
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	entitiesArray := []interfaces.Entity{}
	camerasArray := []interfaces.Entity{}

	// Create a viewport for player 1
	// // Viewport should follow player 1 upwards when the player is close to the top of the screen
	// // Viewport should not follow player 1 downwards when the player is close to the bottom of the screen - player can fall off the bottom of the screen

	// Create player 1
	player1 := entities.NewPlayerEntity(shapes.Vector{X: levelWidth / 2, Y: 0}, controllerManager)
	entitiesArray = append(entitiesArray, player1)

	// Platform guaranteed to be below the player on spawn
	platform := entities.NewPlatformEntity(shapes.Vector{X: 100, Y: 500}, shapes.Rectangle{Width: platformWidth, Height: platformHeight})
	entitiesArray = append(entitiesArray, platform)

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

		platform := entities.NewPlatformEntity(shapes.Vector{X: float64(xPos), Y: float64(yPos)}, shapes.Rectangle{Width: platformWidth, Height: platformHeight})
		entitiesArray = append(entitiesArray, platform)
	}

	// // Platforms should be able to be recycled when they are off below the lowest viewport by a certain amount
	// // Platforms should be added above the viewport when the highest platform is below a certain distance above the highest view port (ensure we always have platforms above the viewport)

	player1Camera := entities.NewCameraEntity(player1, shapes.Rectangle{Width: levelWidth, Height: 400})
	camerasArray = append(camerasArray, player1Camera)

	if err := ebiten.RunGame(&Game{ 
		controllerManager: controllerManager,
		entitiesArray: entitiesArray,
		camerasArray: camerasArray,

		// Benchmarking
		frameThreshold: 100,
	}); err != nil {
		log.Fatal(err)
	}
}

func overwritePlayer1ControllerWithArrowKeys(controllerManager *controllers.ControllerManager) {
	controller := controllerManager.GetController(0)
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