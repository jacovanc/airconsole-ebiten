package entities

import (
	"github.com/jacovanc/airconsole-ebiten/game/components"
	"github.com/jacovanc/airconsole-ebiten/game/controllers"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

const (
	playerSpeed     = 2
	playerJumpSpeed = 6
	playerWidth     = 16
	playerHeight    = 16
)

func NewPlayerEntity(position shapes.Vector, controllerManager *controllers.ControllerManager) interfaces.Entity {
	player := &Entity{
		position:   position,
		dimensions: shapes.Rectangle{Width: playerWidth, Height: playerHeight},
		components: []interfaces.Component{},
		collisions: []shapes.CollisionBox{},
		tags:       []string{"player"},
	}

	renderSpriteComponent := &components.RenderSpriteComponent{
		Entity: player,
		Width:  playerWidth,
		Height: playerHeight,
	}

	jumpComponent := &components.PlayerJumpComponent{
		Player:    player,
		Velocity:  playerJumpSpeed, // Start jumping immediately
		JumpSpeed: playerJumpSpeed,
	}

	inputComponent := &components.InputComponent{
		ControllerId:      0,
		Entity:            player,
		Speed:             playerSpeed,
		ControllerManager: controllerManager,
	}

	collisionComponent := &components.PlayerCollisionComponent{
		Entity: player,
	}

	player.AddComponent(renderSpriteComponent)
	player.AddComponent(jumpComponent)
	player.AddComponent(inputComponent)
	player.AddComponent(collisionComponent)

	player.AddCollision(shapes.CollisionBox{Position: player.position, Box: shapes.Rectangle{Width: playerWidth, Height: playerHeight}})

	return player
}