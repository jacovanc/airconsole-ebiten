package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type cameraComponent struct {
	cameraEntity *entity
	targetEntity *entity
	// Viewport is a box that defines the camera's view. 
	// The position is where it is drawn on the screen, and the box is the size of the view. 
	// The size also affects calculation determining which entities are in view - even though it is the parent entity position that follows the player.
	viewPort collisionBox 
}

func (c *cameraComponent) uniqueName() string {
	return "cameraComponent"
}

func (c *cameraComponent) onUpdate() error {
	// Follow the target, centered on the viewport
	c.cameraEntity.position.x = c.targetEntity.position.x - c.viewPort.box.width / 2
	c.cameraEntity.position.y = c.targetEntity.position.y - c.viewPort.box.height / 2

	return nil
}

func (c *cameraComponent) onDraw(screen *ebiten.Image, offset vector) error {
	// Draw a box around the camera's view
	ebitenutil.DrawRect(screen, c.viewPort.position.x, c.viewPort.position.y, c.viewPort.box.width, c.viewPort.box.height, color.White)
	// Now another box in black 1 px inside the white box
	ebitenutil.DrawRect(screen, c.viewPort.position.x + 1, c.viewPort.position.y + 1, c.viewPort.box.width - 2, c.viewPort.box.height - 2, color.Black)
	return nil
}

func (c *cameraComponent) onCollision(otherEntity *entity) error {
	return nil
}

func (c *cameraComponent) isInView(entity *entity) bool {
	// Check if the entity is in the camera's view
	if entity.position.x > c.cameraEntity.position.x && entity.position.x < c.cameraEntity.position.x + c.viewPort.box.width &&
		entity.position.y > c.cameraEntity.position.y && entity.position.y < c.cameraEntity.position.y + c.viewPort.box.height {
		// The entity is in view
		return true
	} else {
		// The entity is not in view
		return false
	}
}
