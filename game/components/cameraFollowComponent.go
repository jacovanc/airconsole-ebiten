package components

import (
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type CameraFollowComponent struct {
	*BaseComponent
	TargetEntity interfaces.Entity

	// Viewport is a box that defines the camera's view.
	// The position is where it is drawn on the screen, and the box is the size of the view.
	// The size also affects calculation determining which entities are in view - even though it is the parent entity position that follows the player.
	ViewPort shapes.CollisionBox
	// Variables to represent how close to the top or bottom of the viewport the target entity should be before the camera follows it
	DistanceFromTop    float64
	DistanceFromBottom float64
}

func (c *CameraFollowComponent) UniqueName() string {
	return "cameraComponent"
}

func (c *CameraFollowComponent) OnUpdate() error {
	velocity := c.TargetEntity.GetComponent("playerJumpComponent").(*PlayerJumpComponent).Velocity

	targetPosition := c.TargetEntity.GetPosition()
	cameraPosition := c.GetEntity().GetPosition()

	// Only follow if near the top or bottom of the viewport
	// if c.targetEntity.position.y < c.cameraEntity.position.y + c.distanceFromTop && velocity < 0 {
	// 	// Follow the target vertically, centered on the viewport
	// 	c.cameraEntity.position.y = c.cameraEntity.position.y + velocity
	// } else if c.targetEntity.position.y > c.cameraEntity.position.y + c.viewPort.box.height - c.distanceFromBottom && velocity > 0 {
	// 	// Follow the target vertically, centered on the viewport
	// 	c.cameraEntity.position.y = c.cameraEntity.position.y + velocity
	// }

	if targetPosition.Y < cameraPosition.Y+c.DistanceFromTop && velocity < 0 {
		// Follow the target vertically, centered on the viewport
		cameraPosition.Y = cameraPosition.Y + velocity
	} else if targetPosition.Y > cameraPosition.Y+c.ViewPort.Box.Height-c.DistanceFromBottom && velocity > 0 {
		// Follow the target vertically, centered on the viewport
		cameraPosition.Y = cameraPosition.Y + velocity
	}

	return nil
}
