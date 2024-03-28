package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type renderSpriteComponent struct {
	entity    *entity
	width	 int
	height	 int
}

func (c *renderSpriteComponent) uniqueName() string {
	return "inputComponent"
}

func (c *renderSpriteComponent) onUpdate() error {
	return nil
}

func (c *renderSpriteComponent) onDraw(screen *ebiten.Image, camera *cameraComponent) error {
	offset := vector{x: 0, y: 0}
	viewportImage := screen

	// If we have a camera, we need to offset the entity by the camera position and the viewport position
	// Also we can ensure that we only render the entity inside the viewport by using a subImage of the screen
	if(camera != nil) {
		offset = vector{x: camera.cameraEntity.position.x - camera.viewPort.position.x, y: camera.cameraEntity.position.y - camera.viewPort.position.y}
		viewPortRect := image.Rect(int(camera.viewPort.position.x), int(camera.viewPort.position.y), int(camera.viewPort.position.x + camera.viewPort.box.width), int(camera.viewPort.position.y + camera.viewPort.box.height))
		viewportImage = screen.SubImage(viewPortRect).(*ebiten.Image)
	}

	// ebitenutil.DrawRect(screen, c.entity.position.x, c.entity.position.y, float64(c.width), float64(c.height), color.White)
	ebitenutil.DrawRect(viewportImage, c.entity.position.x - offset.x, c.entity.position.y - offset.y, float64(c.width), float64(c.height), color.White)
	return nil
}

func (c *renderSpriteComponent) onCollision(otherEntity *entity) error {
	return nil
}
