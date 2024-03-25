package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type action struct {
	playerId int
	input string
	direction string
}
var lastAction action = action{playerId: -1, input: "", direction: ""}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Print the player ID and action to the screen
	outputString := "Player ID: " + strconv.Itoa(lastAction.playerId) + " Action: " + lastAction.input + " Direction: " + lastAction.direction
	ebitenutil.DebugPrint(screen, outputString);
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	setupAirconsoleInput()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}