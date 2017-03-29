package io

import "github.com/valep27/GChip8/src/emu"

// Frontend is the basic interface for graphical output.
// A frontend might be implemented by SDL, opengl or similar libraries.
type Frontend interface {
	Initialize()
	Draw(emulator emu.Chip8)
	Close()
}

// Key is the type for identifying a key on the Chip8 keypad.
type Key uint8

// KeyEvent is a type for representing keydown or keyup events.
type KeyEvent struct {
	Key Key
	Up bool
}

// The possible values for keys
const (
	Key0 Key = iota
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyQuit
	KeyNone
)

// Input is an interface for a provider of keypresses.
type Input interface {
	Poll() *KeyEvent
}
