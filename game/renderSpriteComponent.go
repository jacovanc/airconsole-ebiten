package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type renderSpriteComponent struct {
	entity    *entity
}

func (c *renderSpriteComponent) uniqueName() string {
	return "inputComponent"
}

func (c *renderSpriteComponent) onUpdate() error {
	return nil
}

func (c *renderSpriteComponent) onDraw(screen *ebiten.Image) error {
	ebitenutil.DrawRect(screen, c.entity.position.x, c.entity.position.y, 16, 16, color.White)
	return nil
}