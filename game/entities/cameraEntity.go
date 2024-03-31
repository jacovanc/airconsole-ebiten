package entities

import (
	"github.com/jacovanc/airconsole-ebiten/game/components"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

func NewCameraEntity(target interfaces.Entity, dimensions shapes.Rectangle) interfaces.Entity {
	camera := &Entity{
		position:   shapes.Vector{X: 0, Y: 0},
		dimensions: shapes.Rectangle{Width: dimensions.Width, Height: dimensions.Height},
		components: []interfaces.Component{},
		collisions: []shapes.CollisionBox{},
		tags:       []string{"camera"},
	}

	cameraViewport := shapes.CollisionBox{
		Position: shapes.Vector{X: 50, Y: 50},
		Box:      shapes.Rectangle{Width: dimensions.Width, Height: dimensions.Height},
	}

	camera.AddComponent(&components.CameraComponent{
		DefaultComponent: &components.DefaultComponent{
			Entity: camera,
		},
		ViewPort:     cameraViewport,
	})

	camera.AddComponent(&components.CameraFollowComponent{
		DefaultComponent: &components.DefaultComponent{
			Entity: camera,
		},
		TargetEntity:       target,
		ViewPort:           cameraViewport,
		DistanceFromTop:    cameraViewport.Box.Height / 2,
		DistanceFromBottom: cameraViewport.Box.Height / 5,
	})

	return camera
}