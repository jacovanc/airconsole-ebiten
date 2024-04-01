package components

import (
	"github.com/jacovanc/airconsole-ebiten/game/collisions"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type PlayerCollisionComponent struct {
	*BaseComponent
	PlayerJumpComponent *PlayerJumpComponent
	CollisionsBoxes     []*shapes.CollisionBox
}

func NewPlayerCollisionComponent(entity interfaces.Entity, playerJumpComponent *PlayerJumpComponent, baseComponent *BaseComponent) *PlayerCollisionComponent {
	comp := &PlayerCollisionComponent{
		BaseComponent:       NewBaseComponent(entity),
		PlayerJumpComponent: playerJumpComponent,
	}

	collisions.AddToPool(comp)

	return comp
}

func (c *PlayerCollisionComponent) UniqueName() string {
	return "playerCollisionComponent"
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
			if c.GetEntity().GetPosition().Y-float64(c.GetEntity().GetDimensions().Height) > otherEntity.GetPosition().Y-float64(otherEntity.GetDimensions().Height) {
				return nil
			}

			// Trigger the jump
			c.PlayerJumpComponent.Velocity = -c.PlayerJumpComponent.JumpSpeed
		}
	}
	return nil
}

func (c *PlayerCollisionComponent) OnUpdate() error {
	// Move all the collisions with the player
	for i := range c.CollisionsBoxes {
		c.CollisionsBoxes[i].Position = *c.GetEntity().GetPosition()
	}
	return nil
}

func (e *PlayerCollisionComponent) AddCollisionBox(collisionBox *shapes.CollisionBox) {
	e.CollisionsBoxes = append(e.CollisionsBoxes, collisionBox)
}

func (e *PlayerCollisionComponent) GetCollisionBoxes() []*shapes.CollisionBox {
	return e.CollisionsBoxes
}
