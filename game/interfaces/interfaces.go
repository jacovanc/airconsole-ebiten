package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image, camera CameraComponent) error
	AddComponent(component Component)
	RemoveComponent(component Component)
	GetComponent(uniqueName string) Component
	GetComponents() []Component
	GetPosition() *shapes.Vector
	GetDimensions() *shapes.Rectangle
	GetTags() []string
}

type Component interface {
	// A unique name identifier of this component to ensure that
	// there are no duplicates, without using reflection
	UniqueName() string
	GetEntity() Entity
}

// Updateable is a component that can be updated (in the Entity update call from the game loop)
type Updateable interface {
	Component
	OnUpdate() error
}

// Drawable is a component that can be drawn (in the Entity draw call from the game loop)
type Drawable interface {
	Component
	// CameraComponent can be passed in to handle camera offset and viewport. Can be nil.
	OnDraw(*ebiten.Image, CameraComponent) error
}

// Collidable is a component that can be collided with (from the Collisions system)
type Collidable interface {
	Component
	AddCollisionBox(*shapes.CollisionBox)
	GetCollisionBoxes() []*shapes.CollisionBox
	OnCollision(Entity) error
}

// CameraComponent is a component that can be used to handle camera offset and viewport
type CameraComponent interface {
	Component
	IsInView(Entity) bool
	GetViewPort() shapes.CollisionBox
}
