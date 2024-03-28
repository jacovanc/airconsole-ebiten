package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// This is actually handles both jumping and falling
type playerJumpComponent struct {
	player    *entity
	jumpSpeed float64
	velocity  float64
}

func (c *playerJumpComponent) uniqueName() string {
	return "playerJumpComponent"
}

func (c *playerJumpComponent) onUpdate() error {
	fmt.Println("Velcity: ", c.velocity)
	c.player.position.y += c.velocity

	c.velocity += 0.1
	if c.velocity > c.jumpSpeed { // If the player is falling too fast, limit the speed
		c.velocity = c.jumpSpeed
	}

	// Jumps are triggered in the playerCollisionComponent

	return nil
}

func (c *playerJumpComponent) onDraw(screen *ebiten.Image, offset vector) error {
	return nil
}

func (c *playerJumpComponent) onCollision(otherEntity *entity) error {
	return nil
}
