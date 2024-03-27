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
	fmt.Println(c.controllerId);
	controller := c.controllerManager.getController(c.controllerId)
	fmt.Println(controller.Inputs)
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