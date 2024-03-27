package main

const (
	playerSpeed     = 2
	playerJumpSpeed = 5
	playerWidth     = 16
	playerHeight    = 16
)

func newPlayerEntity(position vector, controllerManager *controllerManager) *entity {
	player := &entity{
		position:   position,
		components: []component{},
		collisions: []rectangle{},
		tags:       []string{"player"},
	}

	renderSpriteComponent := &renderSpriteComponent{
		entity: player,
		width:  playerWidth,
		height: playerHeight,
	}

	jumpComponent := &playerJumpComponent{
		player:    player,
		velocity:  playerJumpSpeed, // Start jumping immediately
		jumpSpeed: playerJumpSpeed,
	}

	inputComponent := &inputComponent{
		controllerId:      0,
		entity:            player,
		speed:             playerSpeed,
		controllerManager: controllerManager,
	}

	collisionComponent := &playerCollisionComponent{
		entity: player,
	}

	player.addComponent(renderSpriteComponent)
	player.addComponent(jumpComponent)
	player.addComponent(inputComponent)
	player.addComponent(collisionComponent)

	player.addCollision(rectangle{position: player.position, width: playerWidth, height: playerHeight})

	return player
}