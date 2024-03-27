package main

const (
	playerSpeed     = 2
	playerJumpSpeed = 5
	playerSize      = 16
)

func newPlayerEntity(position vector, controllerManager *controllerManager) *entity {
	player := &entity{
		position:   position,
		components: []component{},
	}

	renderSpriteComponent := &renderSpriteComponent{
		entity: player,
	}

	jumpComponent := &playerJumpComponent{
		player:    player,
		velocity:  0,
		jumpSpeed: playerJumpSpeed,
	}

	inputComponent := &inputComponent{
		controllerId:      0,
		entity:            player,
		speed:             playerSpeed,
		controllerManager: controllerManager,
	}

	player.addComponent(renderSpriteComponent)
	player.addComponent(jumpComponent)
	player.addComponent(inputComponent)

	return player
}