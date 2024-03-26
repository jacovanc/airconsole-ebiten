package main

import (
	"syscall/js"
)

type gameInputs struct {
	KeyPressed map[string]bool
}

type Player struct {
	Id int
	Name string
	Inputs gameInputs
}

type PlayerManager struct {
	players map[int]*Player
}

func NewPlayerManager() *PlayerManager {
    pm := &PlayerManager{
        players: make(map[int]*Player),
    }

	pm.setupAirconsoleInput()

	return pm
}

func (pm *PlayerManager) addPlayer(id int, name string) {
	pm.players[id] = &Player{
		Id: id,
		Name: name,
		Inputs: gameInputs{
			KeyPressed: make(map[string]bool),
		},
	}
}

func (pm *PlayerManager) removePlayer(playerId int) {
	delete(pm.players, playerId)
}

func (pm *PlayerManager) getPlayer(playerId int) *Player {
	return pm.players[playerId]
}

func (pm *PlayerManager) setupAirconsoleInput() {
	// Create a callback function for handling input passed from JS
	js.Global().Set("passInputToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			playerId := args[0].Int()
			input := args[1].String()
			direction := args[2].String() // pressed or released

			pm.updateInput(playerId, input, direction)
		}
		return nil
	}))

	 // Handle player connect
	 js.Global().Set("playerConnectToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            playerId := args[0].Int()
			name := args[1].String()
            pm.addPlayer(playerId, name)
        }
        return nil
    }))

    // Handle player disconnect
    js.Global().Set("playerDisconnectFromGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) > 0 {
            playerId := args[0].Int()
            pm.removePlayer(playerId)
        }
        return nil
    }))
}

func (pm *PlayerManager) updateInput(playerId int, input, direction string) {
	// Check if playerInputs already exists for the playerId
	pl := pm.getPlayer(playerId)
	if pl == nil {
		panic("Player not found. Can't assume player exists as we won't have their details like name.")
		// pm.addPlayer(playerId)
		// pl = pm.getPlayer(playerId)
	}

	// Update the gameInput struct with the input
	if(direction == "press") {
		pl.Inputs.KeyPressed[input] = true
	} else {
		pl.Inputs.KeyPressed[input] = false
	}
}
