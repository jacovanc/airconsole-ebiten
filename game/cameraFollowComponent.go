package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type cameraFollowComponent struct {
	cameraEntity *entity
	targetEntity *entity
	// Viewport is a box that defines the camera's view. 
	// The position is where it is drawn on the screen, and the box is the size of the view. 
	// The size also affects calculation determining which entities are in view - even though it is the parent entity position that follows the player.
	viewPort collisionBox
	// Variables to represent how close to the top or bottom of the viewport the target entity should be before the camera follows it
	distanceFromTop float64 
	distanceFromBottom float64
}

func (c *cameraFollowComponent) uniqueName() string {
	return "cameraComponent"
}

func (c *cameraFollowComponent) onUpdate() error {
	velocity := c.targetEntity.getComponent("playerJumpComponent").(*playerJumpComponent).velocity
	fmt.Println(velocity)
	
	// Only follow if near the top or bottom of the viewport
	// || c.targetEntity.position.y > c.cameraEntity.position.y + c.viewPort.box.height - 20
	if c.targetEntity.position.y < c.cameraEntity.position.y + c.distanceFromTop && velocity < 0 {
		// Follow the target vertically, centered on the viewport
		c.cameraEntity.position.y = c.cameraEntity.position.y + velocity
	} else if c.targetEntity.position.y > c.cameraEntity.position.y + c.viewPort.box.height - c.distanceFromBottom && velocity > 0 {
		// Follow the target vertically, centered on the viewport
		c.cameraEntity.position.y = c.cameraEntity.position.y + velocity
	}
	return nil
}

func (c *cameraFollowComponent) onDraw(screen *ebiten.Image, camera *cameraComponent) error {
	return nil
}

func (c *cameraFollowComponent) onCollision(otherEntity *entity) error {
	return nil
}