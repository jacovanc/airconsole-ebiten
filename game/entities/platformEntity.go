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
		tags:       []string{"platform"},
	}

	baseComponent := components.NewBaseComponent(platform)

	platform.AddComponent(&components.RenderSpriteComponent{
		BaseComponent: baseComponent,
		Width:         int(dimensions.Width),
		Height:        int(dimensions.Height),
	})

	collisionComponent := components.NewCollisionComponent(platform, baseComponent)
	collisionComponent.AddCollisionBox(&shapes.CollisionBox{Position: position, Box: shapes.Rectangle{Width: dimensions.Width, Height: dimensions.Height}})

	platform.AddComponent(collisionComponent)

	return platform
}
