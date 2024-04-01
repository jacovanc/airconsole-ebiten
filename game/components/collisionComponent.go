package components

import (
	"github.com/jacovanc/airconsole-ebiten/game/collisions"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type CollisionComponent struct {
	*BaseComponent
	collisionsBoxes []*shapes.CollisionBox
}

func NewCollisionComponent(entity interfaces.Entity, baseComponent *BaseComponent) *CollisionComponent {
	comp := &CollisionComponent{
		BaseComponent: baseComponent,
	}

	collisions.AddToPool(comp)

	return comp
}

func (c *CollisionComponent) UniqueName() string {
	return "collisionComponent"
}

func (c *CollisionComponent) OnUpdate() error {
	// Move all the collisions with the player
	for i := range c.collisionsBoxes {
		c.collisionsBoxes[i].Position = *c.GetEntity().GetPosition()
	}
	return nil
}

func (c *CollisionComponent) OnCollision(otherEntity interfaces.Entity) error {
	return nil
}

func (e *CollisionComponent) AddCollisionBox(collisionBox *shapes.CollisionBox) {
	e.collisionsBoxes = append(e.collisionsBoxes, collisionBox)
}

func (e *CollisionComponent) GetCollisionBoxes() []*shapes.CollisionBox {
	return e.collisionsBoxes
}
