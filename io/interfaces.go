package io

import "github.com/valep27/GChip8/emu"

// Frontend is the basic interface for graphical output.
// A frontend might be implemented by SDL, opengl or similar libraries.
type Frontend interface {
	Initialize()
	Draw(emulator emu.Chip8)
	Close()
}

// Key is the type for identifying a key on the Chip8 keypad.
type Key uint8

// The possible values for keys
const (
	KEY_0 Key = iota
	KEY_1
	KEY_2
	KEY_3
	KEY_4
	KEY_5
	KEY_6
	KEY_7
	KEY_8
	KEY_9
	KEY_A
	KEY_B
	KEY_C
	KEY_D
	KEY_E
	KEY_F
	KEY_QUIT
	KEY_NONE
)

// Input is an interface for a provider of keypresses.
type Input interface {
	Poll() (Key, bool)
}
