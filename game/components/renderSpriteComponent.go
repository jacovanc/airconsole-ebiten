package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type RenderSpriteComponent struct {
	*DefaultComponent
	Width	 int
	Height	 int
}

func (c *RenderSpriteComponent) UniqueName() string {
	return "inputComponent"
}

func (c *RenderSpriteComponent) OnUpdate() error {
	return nil
}

// func (c *RenderSpriteComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
func (c *RenderSpriteComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	offset := shapes.Vector{X: 0, Y: 0}
	viewportImage := screen

	// If we have a camera, we need to offset the entity by the camera position and the viewport position
	// Also we can ensure that we only render the entity inside the viewport by using a subImage of the screen
	if(camera != nil) {
		cameraEntityPosition := camera.GetEntity().GetPosition()
		cameraViewPortPosition := camera.GetViewPort().Position
		cameraViewPortBox := camera.GetViewPort().Box

		offset = shapes.Vector{X: cameraEntityPosition.X - cameraViewPortPosition.X, Y: cameraEntityPosition.Y - cameraViewPortPosition.Y}
		viewPortRect := image.Rect(int(cameraViewPortPosition.X), int(cameraViewPortPosition.Y), int(cameraViewPortPosition.X + cameraViewPortBox.Width), int(cameraViewPortPosition.Y + cameraViewPortBox.Height))
		viewportImage = screen.SubImage(viewPortRect).(*ebiten.Image)

	}

	// ebitenutil.DrawRect(screen, c.entity.position.x, c.entity.position.y, float64(c.width), float64(c.height), color.White)
	ebitenutil.DrawRect(viewportImage, c.Entity.GetPosition().X - offset.X, c.Entity.GetPosition().Y - offset.Y, float64(c.Width), float64(c.Height), color.White)
	return nil
}