package components

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacovanc/airconsole-ebiten/game/controllers"
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
)

type InputComponent struct {
	ControllerId	  int
	Entity    interfaces.Entity
	Speed     int
	ControllerManager *controllers.ControllerManager
}

func (c *InputComponent) UniqueName() string {
	return "inputComponent"
}

func (c *InputComponent) OnUpdate() error {
	controller := c.ControllerManager.GetController(c.ControllerId)
	if controller == nil {
		fmt.Println("Controller not found")
		return nil
	}
	if(controller.Inputs.KeyPressed["left"]) {
		c.Entity.GetPosition().X -= float64(c.Speed)
	}
	if(controller.Inputs.KeyPressed["right"]) {
		c.Entity.GetPosition().X += float64(c.Speed)
	}
	return nil
}

func (c *InputComponent) OnDraw(screen *ebiten.Image, camera interfaces.CameraComponent) error {
	return nil
}

func (c *InputComponent) OnCollision(otherEntity interfaces.Entity) error {
	return nil
}

func (c *InputComponent) GetEntity() interfaces.Entity {
	return c.Entity
}