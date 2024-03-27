package main

func newPlatformEntity(position vector) *entity {
	platform := &entity{
		position:   position,
		components: []component{},
		collisions: []rectangle{},
		tags:       []string{"platform"},
	}
	platform.addComponent(&renderSpriteComponent{entity: platform, width: 32, height: 10})

	platform.addCollision(rectangle{position: position, width: 32, height: 10})

	return platform
}