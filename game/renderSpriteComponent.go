package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type renderSpriteComponent struct {
	entity    *entity
	width	 int
	height	 int
}

func (c *renderSpriteComponent) uniqueName() string {
	return "inputComponent"
}

func (c *renderSpriteComponent) onUpdate() error {
	return nil
}

func (c *renderSpriteComponent) onDraw(screen *ebiten.Image, offset vector) error {
	// ebitenutil.DrawRect(screen, c.entity.position.x, c.entity.position.y, float64(c.width), float64(c.height), color.White)
	ebitenutil.DrawRect(screen, c.entity.position.x - offset.x, c.entity.position.y - offset.y, float64(c.width), float64(c.height), color.White)
	fmt.Println("Drawing sprite at: ", c.entity.position.x - offset.x, c.entity.position.y - offset.y)
	return nil
}

func (c *renderSpriteComponent) onCollision(otherEntity *entity) error {
	return nil
}
