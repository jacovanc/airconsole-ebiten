package main

import (
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
	for _, tag := range otherEntity.tags {

		if tag == "platform" {
			if playerJumpComponent := c.entity.getComponent("playerJumpComponent").(*playerJumpComponent); playerJumpComponent != nil {
				// Don't do anything if the player is not falling
				if playerJumpComponent.velocity > 0 {
					return nil
				}

				// If the the bottom of the player is lower than the bottom of the platform, don't do anything
				if c.entity.position.y < otherEntity.position.y {
					return nil
				}

				// Trigger the jump
				playerJumpComponent.velocity = playerJumpComponent.jumpSpeed
			}
		}
	}
	return nil
}
