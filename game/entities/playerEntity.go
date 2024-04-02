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
		tags:       []string{"player"},
	}

	baseComponent := components.NewBaseComponent(player)

	renderSpriteComponent := &components.RenderSpriteComponent{
		BaseComponent: baseComponent,
		Width:         playerWidth,
		Height:        playerHeight,
	}

	jumpComponent := &components.PlayerJumpComponent{
		BaseComponent: baseComponent,
		Velocity:      playerJumpSpeed, // Start jumping immediately
		JumpSpeed:     playerJumpSpeed,
	}

	inputComponent := &components.InputComponent{
		BaseComponent:     baseComponent,
		ControllerId:      0,
		Speed:             playerSpeed,
		ControllerManager: controllerManager,
	}

	playerCollisionComponent := components.NewPlayerCollisionComponent(player, jumpComponent, baseComponent)
	// Add the collision box
	playerCollisionComponent.AddCollisionBox(&shapes.CollisionBox{Position: player.position, Box: shapes.Rectangle{Width: playerWidth, Height: playerHeight}})

	player.AddComponent(renderSpriteComponent)
	player.AddComponent(jumpComponent)
	player.AddComponent(inputComponent)
	player.AddComponent(playerCollisionComponent)

	return player
}
