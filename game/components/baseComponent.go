package components

import (
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
)

// BaseComponent provides default implementations for the Component interface
// It also handles the Entity reference and provides a default implementation for GetEntity
// Attaching this to a component will handle the Entity reference for you, or it can not be provided
// to provide its own Entity reference.
type BaseComponent struct {
	entity interfaces.Entity
}

func NewBaseComponent(entity interfaces.Entity) *BaseComponent {
	return &BaseComponent{
		entity: entity,
	}
}

func (c *BaseComponent) GetEntity() interfaces.Entity {
	return c.entity
}
