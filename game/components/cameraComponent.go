package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type CameraComponent struct {
	*DefaultComponent
	CameraEntity interfaces.Entity
	// Viewport is a box that defines the camera's view. 
	// The position is where it is drawn on the screen, and the box is the size of the view. 
	// The size also affects calculation determining which entities are in view - even though it is the parent entity position that follows the player.
	ViewPort shapes.CollisionBox
}

func (c *CameraComponent) UniqueName() string {
	return "cameraComponent"
}

func (c *CameraComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	// Draw a box around the camera's view
	ebitenutil.DrawRect(screen, c.ViewPort.Position.X, c.ViewPort.Position.Y, c.ViewPort.Box.Width, c.ViewPort.Box.Height, color.White)
	// Now another box in black 1 px inside the white box
	ebitenutil.DrawRect(screen, c.ViewPort.Position.X + 1, c.ViewPort.Position.Y + 1, c.ViewPort.Box.Width - 2, c.ViewPort.Box.Height - 2, color.Black)
	return nil
}

// Considers the entities width and hight to ensure it doesn't pop in and out of view
func (c *CameraComponent) IsInView(entity interfaces.Entity) bool {
	entityPosition := entity.GetPosition()
	entityDimensions := entity.GetDimensions()

	cameraEntityPosition := c.Entity.GetPosition()

	// Check if the entity is in the camera's view
	return entityPosition.X + entityDimensions.Width > cameraEntityPosition.X && entityPosition.X < cameraEntityPosition.X + c.ViewPort.Box.Width &&
		entityPosition.Y + entityDimensions.Height > cameraEntityPosition.Y && entityPosition.Y < cameraEntityPosition.Y + c.ViewPort.Box.Height
}

func (c *CameraComponent) GetViewPort() shapes.CollisionBox {
	return c.ViewPort
}
