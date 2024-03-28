package main

const (
	width  = 300
	height = 400
)

func newCameraEntity(target *entity) *entity {
	camera := &entity{
		position:   vector{x: 0, y: 0}, // Doesn't matter, will be set by the camera component to follow the target
		components: []component{},
		collisions: []collisionBox{},
		tags:       []string{"camera"},
	}

	camera.addComponent(&cameraComponent{cameraEntity: camera, targetEntity: target, viewPort: collisionBox{position: vector{x: 0, y: 0}, box: rectangle{width: width, height: height}}})

	return camera
}