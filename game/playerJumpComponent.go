package main

import "github.com/hajimehoshi/ebiten/v2"

type playerJumpComponent struct {
	player    *entity
	jumpSpeed float64
	velocity  float64
}

func (c *playerJumpComponent) uniqueName() string {
	return "playerJumpComponent"
}

func (c *playerJumpComponent) onUpdate() error {
	c.player.position.y += c.velocity

	c.velocity -= 0.1
	if c.velocity < -c.jumpSpeed { // If the player is falling too fast, limit the speed
		c.velocity = -c.jumpSpeed
	}

	return nil
}

func (c *playerJumpComponent) onDraw(screen *ebiten.Image) error {
	return nil
}

func (c *playerJumpComponent) onCollision(otherEntity *entity) error {
	return nil
}
