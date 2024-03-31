package entities

import (
	"github.com/jacovanc/airconsole-ebiten/game/components"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

func NewPlatformEntity(position shapes.Vector, dimensions shapes.Rectangle) interfaces.Entity {
	platform := &Entity{
		position:   position,
		dimensions: shapes.Rectangle{Width: dimensions.Width, Height: dimensions.Height},
		components: []interfaces.Component{},
		collisions: []shapes.CollisionBox{},
		tags:       []string{"platform"},
	}
	platform.AddComponent(&components.RenderSpriteComponent{
		DefaultComponent: &components.DefaultComponent{
			Entity: platform,
		},
		Width: int(dimensions.Width), 
		Height: int(dimensions.Height),
	})

	platform.AddCollision(shapes.CollisionBox{Position: position, Box: shapes.Rectangle{Width: dimensions.Width, Height: dimensions.Height}})

	return platform
}