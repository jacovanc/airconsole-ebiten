package main

import (
	"syscall/js"
)

type gameInputs struct {
	KeyPressed map[string]bool
}

type Controller struct {
	Id int
	Name string
	Inputs gameInputs
}

type ControllerManager struct {
	controllers map[int]*Controller
}

func NewControllerManager() *ControllerManager {
    cm := &ControllerManager{
        controllers: make(map[int]*Controller),
    }

	cm.setupAirconsoleInput()

	return cm
}

func (pm *ControllerManager) addController(id int) {
	pm.controllers[id] = &Controller{
		Id: id,
		Inputs: gameInputs{
			KeyPressed: make(map[string]bool),
		},
	}
}

func (pm *ControllerManager) removeController(ControllerId int) {
	delete(pm.controllers, ControllerId)
}

func (pm *ControllerManager) getController(ControllerId int) *Controller {
	return pm.controllers[ControllerId]
}

func (pm *ControllerManager) setupAirconsoleInput() {
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

	// Handle Controller connect
	js.Global().Set("controllerConnectToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            controllerId := args[0].Int()
            pm.addController(controllerId)
        }
        return nil
    }))

    // Handle Controller disconnect
    js.Global().Set("controllerDisconnectFromGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            controllerId := args[0].Int()
            pm.removeController(controllerId)
        }
        return nil
    }))
}

func (pm *ControllerManager) updateInput(controllerId int, input, direction string) {
	// Check if ControllerInputs already exists for the ControllerId
	pl := pm.getController(controllerId)
	if pl == nil {
		panic("Controller not found. Can't assume Controller exists as we won't have their details like name.")
		// pm.addController(ControllerId)
		// pl = pm.getController(ControllerId)
	}

	// Update the gameInput struct with the input
	if(direction == "press") {
		pl.Inputs.KeyPressed[input] = true
	} else {
		pl.Inputs.KeyPressed[input] = false
	}
}
