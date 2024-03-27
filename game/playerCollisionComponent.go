package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type playerCollisionComponent struct {
	entity *entity
}

func (c *playerCollisionComponent) uniqueName() string {
	return "playerCollisionComponent"
}

func (c *playerCollisionComponent) onUpdate() error {
	// Move all the collisions with the player
	for i := range c.entity.collisions {
		c.entity.collisions[i].position = c.entity.position
	}
	return nil
}

func (c *playerCollisionComponent) onDraw(screen *ebiten.Image) error {
	return nil
}

func (c *playerCollisionComponent) onCollision(otherEntity *entity) error {
	// If otherEntity.tags array contains "platform"
	fmt.Println("Player collided with platform")
	for _, tag := range otherEntity.tags {
		if tag == "platform" {
			playerJumpComponent := c.entity.getComponent("playerJumpComponent").(*playerJumpComponent)
			if playerJumpComponent != nil {
				// Don't do anything if the player is not falling
				if playerJumpComponent.velocity <= 0 {
					// Trigger the jump
					playerJumpComponent.velocity = playerJumpComponent.jumpSpeed
				}
			}
		}
	}
	return nil
}
