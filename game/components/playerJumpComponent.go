package components

// This is actually handles both jumping and falling
type PlayerJumpComponent struct {
	*DefaultComponent
	JumpSpeed float64
	Velocity  float64
}

func (c *PlayerJumpComponent) UniqueName() string {
	return "playerJumpComponent"
}

func (c *PlayerJumpComponent) OnUpdate() error {
	c.Entity.GetPosition().Y += c.Velocity

	c.Velocity += 0.15
	if c.Velocity > c.JumpSpeed*3 { // If the player is falling too fast, limit the speed
		c.Velocity = c.JumpSpeed*3
	}

	// Jumps are triggered in the playerCollisionComponent

	return nil
}
