package main

func newCameraEntity(target *entity) *entity {
	camera := &entity{
		position:   vector{x: 0, y: 0},
		dimensions: rectangle{width: levelWidth, height: 400},
		components: []component{},
		collisions: []collisionBox{},
		tags:       []string{"camera"},
	}

	camera.addComponent(&cameraComponent{
		cameraEntity: camera,
		targetEntity: target,
		viewPort: collisionBox{
			position: vector{x: 50, y: 50},
			box:      rectangle{width: levelWidth, height: 400},
		},
	})

	return camera
}