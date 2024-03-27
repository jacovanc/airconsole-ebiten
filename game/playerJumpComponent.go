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
	if c.player.position.y > 0 {
		c.velocity -= 0.1
	} else {
		c.player.position.y = 0
		c.velocity = c.jumpSpeed
	}

	c.player.position.y += c.velocity

	return nil
}

func (c *playerJumpComponent) onDraw(screen *ebiten.Image) error {
	return nil
}

func (c *playerJumpComponent) onCollision(otherEntity *entity) error {
	return nil
}
