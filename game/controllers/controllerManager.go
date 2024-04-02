package controllers

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

type ControllerManager struct {
	Controllers map[int]*controller
}

func NewControllerManager() *ControllerManager {
    cm := &ControllerManager{
        Controllers: make(map[int]*controller),
    }

	cm.setupAirconsoleInput()

	return cm
}

func (pm *ControllerManager) AddController(id int) {
	pm.Controllers[id] = &controller{
		Id: id,
		Inputs: gameInputs{
			KeyPressed: make(map[string]bool),
		},
	}
}

func (pm *ControllerManager) removeController(controllerId int) {
	delete(pm.Controllers, controllerId)
}

func (pm *ControllerManager) GetController(controllerId int) *controller {
	return pm.Controllers[controllerId]
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

	// Handle controller connect
	js.Global().Set("controllerConnectToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            controllerId := args[0].Int()
            pm.AddController(controllerId)
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

func (pm *ControllerManager) updateInput(controllerId int, input, direction string) {
	// Check if controllerInputs already exists for the controllerId
	pl := pm.GetController(controllerId)
	if pl == nil {
		// pm.Addcontroller(controllerId)
		// pl = pm.getcontroller(controllerId)
	}

	// Update the gameInput struct with the input
	if(direction == "press") {
		pl.Inputs.KeyPressed[input] = true
	} else {
		pl.Inputs.KeyPressed[input] = false
	}
}
