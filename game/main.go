package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	PlayerManager *PlayerManager
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Loop through players and output their inputs
	for _, player := range g.PlayerManager.players {
		for input, pressed := range player.Inputs.KeyPressed {
			if(pressed) {
				ebitenutil.DebugPrint(screen, "player " + strconv.Itoa(player.Id + 1) + "(" + player.Name + ") is pressing key " + input)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	playerManager := NewPlayerManager()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{ PlayerManager: playerManager }); err != nil {
		log.Fatal(err)
	}
}