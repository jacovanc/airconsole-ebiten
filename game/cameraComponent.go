package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type cameraComponent struct {
	cameraEntity *entity
	// Viewport is a box that defines the camera's view. 
	// The position is where it is drawn on the screen, and the box is the size of the view. 
	// The size also affects calculation determining which entities are in view - even though it is the parent entity position that follows the player.
	viewPort collisionBox
}

func (c *cameraComponent) uniqueName() string {
	return "cameraComponent"
}

func (c *cameraComponent) onUpdate() error {
	return nil
}

func (c *cameraComponent) onDraw(screen *ebiten.Image, camera *cameraComponent) error {
	// Draw a box around the camera's view
	ebitenutil.DrawRect(screen, c.viewPort.position.x, c.viewPort.position.y, c.viewPort.box.width, c.viewPort.box.height, color.White)
	// Now another box in black 1 px inside the white box
	ebitenutil.DrawRect(screen, c.viewPort.position.x + 1, c.viewPort.position.y + 1, c.viewPort.box.width - 2, c.viewPort.box.height - 2, color.Black)
	return nil
}

func (c *cameraComponent) onCollision(otherEntity *entity) error {
	return nil
}

// Considers the entities width and hight to ensure it doesn't pop in and out of view
func (c *cameraComponent) isInView(entity *entity) bool {
	// Check if the entity is in the camera's view
	return entity.position.x + entity.dimensions.width > c.cameraEntity.position.x && entity.position.x < c.cameraEntity.position.x + c.viewPort.box.width &&
		entity.position.y + entity.dimensions.height > c.cameraEntity.position.y && entity.position.y < c.cameraEntity.position.y + c.viewPort.box.height
}
