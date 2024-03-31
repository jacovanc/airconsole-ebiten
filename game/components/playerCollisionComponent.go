package components

import (
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
)

type PlayerCollisionComponent struct {
	*DefaultComponent
	PlayerJumpComponent *PlayerJumpComponent
}

func (c *PlayerCollisionComponent) UniqueName() string {
	return "playerCollisionComponent"
}

func (c *PlayerCollisionComponent) OnUpdate() error {
	collisions := c.Entity.GetCollisions()
	// Move all the collisions with the player
	for i := range *collisions {
		(*collisions)[i].Position = *c.Entity.GetPosition()
	}
	return nil
}

func (c *PlayerCollisionComponent) OnCollision(otherEntity interfaces.Entity) error {
	// If otherEntity.tags array contains "platform"
	for _, tag := range otherEntity.GetTags() {

		if tag == "platform" {
			// Don't do anything if the player is not falling
			if c.PlayerJumpComponent.Velocity < 0 {
				return nil
			}

			// If the the bottom of the player is lower than the bottom of the platform, don't do anything
			if c.Entity.GetPosition().Y - float64(c.Entity.GetDimensions().Height) > otherEntity.GetPosition().Y - float64(otherEntity.GetDimensions().Height) {
				return nil
			}

			// Trigger the jump
			c.PlayerJumpComponent.Velocity = -c.PlayerJumpComponent.JumpSpeed
		}
	}
	return nil
}
