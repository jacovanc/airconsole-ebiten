package components

import (
	"fmt"

	"github.com/jacovanc/airconsole-ebiten/game/controllers"
)

type InputComponent struct {
	*BaseComponent
	ControllerId      int
	Speed             int
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
	if controller.Inputs.KeyPressed["left"] {
		c.GetEntity().GetPosition().X -= float64(c.Speed)
	}
	if controller.Inputs.KeyPressed["right"] {
		c.GetEntity().GetPosition().X += float64(c.Speed)
	}
	return nil
}
