package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type inputComponent struct {
	controllerId	  int
	entity    *entity
	speed     int
	controllerManager *controllerManager
}

func (c *inputComponent) uniqueName() string {
	return "inputComponent"
}

func (c *inputComponent) onUpdate() error {
	controller := c.controllerManager.getController(c.controllerId)
	if controller == nil {
		fmt.Println("Controller not found")
		return nil
	}
	if(controller.Inputs.KeyPressed["left"]) {
		c.entity.position.x -= float64(c.speed)
	}
	if(controller.Inputs.KeyPressed["right"]) {
		c.entity.position.x += float64(c.speed)
	}
	return nil
}

func (c *inputComponent) onDraw(screen *ebiten.Image) error {
	return nil
}

func (c *inputComponent) onCollision(otherEntity *entity) error {
	return nil
}
