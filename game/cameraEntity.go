package main

func newCameraEntity(target *entity) *entity {
	camera := &entity{
		position:   vector{x: 0, y: 0},
		dimensions: rectangle{width: levelWidth, height: 400},
		components: []component{},
		collisions: []collisionBox{},
		tags:       []string{"camera"},
	}

	cameraViewport := collisionBox{
		position: vector{x: 50, y: 50},
		box:      rectangle{width: levelWidth, height: 400},
	}

	camera.addComponent(&cameraComponent{
		cameraEntity: camera,
		viewPort:     cameraViewport,
	})

	camera.addComponent(&cameraFollowComponent{
		cameraEntity:       camera,
		targetEntity:       target,
		viewPort:           cameraViewport,
		distanceFromTop:    cameraViewport.box.height / 2,
		distanceFromBottom: cameraViewport.box.height / 5,
	})

	return camera
}