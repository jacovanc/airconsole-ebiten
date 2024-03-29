package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image, camera CameraComponent) error
	Collision(otherEntity Entity) error
	AddComponent(component Component)
	RemoveComponent(component Component)
	GetComponent(uniqueName string) Component
	AddCollision(collisionBox shapes.CollisionBox)
	GetCollisions() *[]shapes.CollisionBox
	GetPosition() *shapes.Vector
	GetDimensions() *shapes.Rectangle
	GetTags() []string
}

type Component interface {
	OnUpdate() error
	// CameraComponent can be passed in to handle camera offset and viewport. Can be nil.
	OnDraw(*ebiten.Image, CameraComponent) error 
	OnCollision(Entity) error
	// A unique name identifier of this component to ensure that
	// there are no duplicates, without using reflection
	UniqueName() string 

	GetEntity() Entity
}

type CameraComponent interface {
	Component
	IsInView(Entity) bool
	GetViewPort() shapes.CollisionBox
}