package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type entity struct {
	position vector
	components []component
	collisions []rectangle
	tags []string
}

type vector struct {
	x, y float64
}

type rectangle struct {
	position vector
	width, height float64
}

type component interface {
	onUpdate() error
	onDraw(*ebiten.Image) error
	onCollision(*entity) error
	// A unique name identifier of this component to ensure that
	// there are no duplicates, without using reflection
	uniqueName() string 
}

func (e *entity) update() error {
	for _, component := range e.components {
		if err := component.onUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (e *entity) draw(screen *ebiten.Image) error {
	for _, component := range e.components {
		if err := component.onDraw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (e *entity) collision(otherEntity *entity) error {
	for _, component := range e.components {
		if err := component.onCollision(otherEntity); err != nil {
			return err
		}
	}
	return nil
}

func (e *entity) addComponent(component component) {
	e.components = append(e.components, component)
}

func (e *entity) removeComponent(component component) {
	for i, c := range e.components {
		if c.uniqueName() == component.uniqueName() {
			e.components = append(e.components[:i], e.components[i+1:]...)
		}
	}
}

func (e *entity) getComponent(uniqueName string) component {
	for _, c := range e.components {
		if(c.uniqueName() == uniqueName) {
			return c
		}
	}
	return nil
}

func (e *entity) addCollision(rectangle rectangle) {
	e.collisions = append(e.collisions, rectangle)
}


