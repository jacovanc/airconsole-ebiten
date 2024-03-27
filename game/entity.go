package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type entity struct {
	position vector
	components []component
}

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(screen *ebiten.Image) error

	// A unique name identifier of this component to ensure that
	// there are no duplicates, without using reflection
	uniqueName() string 
}

func (e *entity) onUpdate() error {
	for _, component := range e.components {
		if err := component.onUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (e *entity) onDraw(screen *ebiten.Image) error {
	for _, component := range e.components {
		if err := component.onDraw(screen); err != nil {
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


