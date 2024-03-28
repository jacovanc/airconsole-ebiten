package main

import (
	"syscall/js"
)

type gameInputs struct {
	KeyPressed map[string]bool
}

type controller struct {
	Id int
	Inputs gameInputs
}

type controllerManager struct {
	controllers map[int]*controller
}

func newControllerManager() *controllerManager {
    cm := &controllerManager{
        controllers: make(map[int]*controller),
    }

	cm.setupAirconsoleInput()

	return cm
}

func (pm *controllerManager) addController(id int) {
	pm.controllers[id] = &controller{
		Id: id,
		Inputs: gameInputs{
			KeyPressed: make(map[string]bool),
		},
	}
}

func (pm *controllerManager) removeController(controllerId int) {
	delete(pm.controllers, controllerId)
}

func (pm *controllerManager) getController(controllerId int) *controller {
	return pm.controllers[controllerId]
}

func (pm *controllerManager) setupAirconsoleInput() {
	// Create a callback function for handling input passed from JS
	js.Global().Set("passInputToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			controllerId := args[0].Int()
			input := args[1].String()
			direction := args[2].String() // pressed or released

			pm.updateInput(controllerId, input, direction)
		}
		return nil
	}))

	// Handle controller connect
	js.Global().Set("controllerConnectToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            controllerId := args[0].Int()
            pm.addController(controllerId)
        }
        return nil
    }))

    // Handle controller disconnect
    js.Global().Set("controllerDisconnectFromGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            controllerId := args[0].Int()
            pm.removeController(controllerId)
        }
        return nil
    }))
}

func (pm *controllerManager) updateInput(controllerId int, input, direction string) {
	// Check if controllerInputs already exists for the controllerId
	pl := pm.getController(controllerId)
	if pl == nil {
		// pm.addcontroller(controllerId)
		// pl = pm.getcontroller(controllerId)
	}

	// Update the gameInput struct with the input
	if(direction == "press") {
		pl.Inputs.KeyPressed[input] = true
	} else {
		pl.Inputs.KeyPressed[input] = false
	}
}
