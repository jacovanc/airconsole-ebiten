package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
)

// DefaultComponent provides default implementations for the Component interface
// This is useful for components that don't need to implement all the methods of
// the Component interface without writing empty methods
// It also handles the Entity reference and provides a default implementation for GetEntity
type DefaultComponent struct {
	Entity interfaces.Entity
}

func (c *DefaultComponent) OnUpdate() error {
	return nil
}

func (c *DefaultComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	return nil
}

func (c *DefaultComponent) OnCollision(otherEntity interfaces.Entity) error {
	return nil
}

func (c *DefaultComponent) GetEntity() interfaces.Entity {
	return c.Entity
}
