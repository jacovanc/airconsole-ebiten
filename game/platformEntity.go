package main

const (
	platformWidth  = 32
	platformHeight = 10
)

func newPlatformEntity(position vector) *entity {
	platform := &entity{
		position:   position,
		dimensions: rectangle{width: platformWidth, height: platformHeight},
		components: []component{},
		collisions: []collisionBox{},
		tags:       []string{"platform"},
	}
	platform.addComponent(&renderSpriteComponent{entity: platform, width: platformWidth, height: platformHeight})

	platform.addCollision(collisionBox{position: position, box: rectangle{width: platformWidth, height: platformHeight}})

	return platform
}