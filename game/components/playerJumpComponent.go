package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
)

// This is actually handles both jumping and falling
type PlayerJumpComponent struct {
	Player    interfaces.Entity
	JumpSpeed float64
	Velocity  float64
}

func (c *PlayerJumpComponent) UniqueName() string {
	return "playerJumpComponent"
}

func (c *PlayerJumpComponent) OnUpdate() error {
	c.Player.GetPosition().Y += c.Velocity

	c.Velocity += 0.15
	if c.Velocity > c.JumpSpeed*3 { // If the player is falling too fast, limit the speed
		c.Velocity = c.JumpSpeed*3
	}

	// Jumps are triggered in the playerCollisionComponent

	return nil
}

func (c *PlayerJumpComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	return nil
}

func (c *PlayerJumpComponent) OnCollision(otherEntity interfaces.Entity) error {
	return nil
}

func (c *PlayerJumpComponent) GetEntity() interfaces.Entity {
	return c.Player
}