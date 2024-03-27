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
	return nil
}

func (c *playerCollisionComponent) onDraw(screen *ebiten.Image) error {
	return nil
}

func (c *playerCollisionComponent) onCollision(otherEntity *entity) error {
	// If otherEntity.tags array contains "platform"
	for _, tag := range otherEntity.tags {
		if tag == "platform" {
			playerJumpComponent := c.entity.getComponent("playerJumpComponent").(*playerJumpComponent)
			if playerJumpComponent != nil {
				// Trigger the jump
				playerJumpComponent.velocity = playerJumpComponent.jumpSpeed
			}
		}
	}
	return nil
}
