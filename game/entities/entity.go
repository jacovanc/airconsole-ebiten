package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type Entity struct {
	position shapes.Vector
	dimensions shapes.Rectangle
	components []interfaces.Component
	collisions []shapes.CollisionBox
	tags []string
}

func (e *Entity) Update() error {
	for _, component := range e.components {
		if err := component.OnUpdate(); err != nil {
			return err
		}
	}
	return nil
}

// Offset handles the camera offset to ensure we render it inside the camera viewport
func (e *Entity) Draw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	for _, component := range e.components {
		if err := component.OnDraw(screen, camera); err != nil {
			return err
		}
	}
	return nil
}

func (e *Entity) Collision(otherEntity interfaces.Entity) error {
	for _, component := range e.components {
		if err := component.OnCollision(otherEntity); err != nil {
			return err
		}
	}
	return nil
}

func (e *Entity) AddComponent(component interfaces.Component) {
	e.components = append(e.components, component)
}

func (e *Entity) RemoveComponent(component interfaces.Component) {
	for i, c := range e.components {
		if c.UniqueName() == component.UniqueName() {
			e.components = append(e.components[:i], e.components[i+1:]...)
		}
	}
}

func (e *Entity) GetComponent(uniqueName string) interfaces.Component {
	for _, c := range e.components {
		if(c.UniqueName() == uniqueName) {
			return c
		}
	}
	return nil
}

func (e *Entity) AddCollision(collisionBox shapes.CollisionBox) {
	e.collisions = append(e.collisions, collisionBox)
}

func (e *Entity) GetCollisions() *[]shapes.CollisionBox {
	return &e.collisions
}

func (e *Entity) GetPosition() *shapes.Vector {
	return &e.position
}
func (e *Entity) GetDimensions() *shapes.Rectangle {
	return &e.dimensions
}

func (e *Entity) GetTags() []string {
	return e.tags
}