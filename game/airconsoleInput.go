package main

import "syscall/js"

func setupAirconsoleInput() {
	// Create a callback function for handling input passed from JS
	js.Global().Set("passInputToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			action := action{
				playerId: args[0].Int(),
				input:   args[1].String(),
				direction: args[2].String(), // pressed or released
			}
			lastAction = action
		}
		return nil
	}))

	// Create a function for pausing the game
	js.Global().Set("passPauseToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Should pause the game however needed
		return nil
	}))

	// Create a function for resuming the game
	js.Global().Set("passResumeToGame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Should resume the game however needed
		return nil
	}))

	//
}