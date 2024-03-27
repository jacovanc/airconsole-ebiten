package main

import (
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

func (c *renderSpriteComponent) onDraw(screen *ebiten.Image) error {
	ebitenutil.DrawRect(screen, c.entity.position.x, c.entity.position.y, float64(c.width), float64(c.height), color.White)
	return nil
}

func (c *renderSpriteComponent) onCollision(otherEntity *entity) error {
	return nil
}
