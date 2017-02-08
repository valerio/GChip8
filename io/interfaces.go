package io

import "github.com/valep27/chip8/emu"

// Frontend is the basic interface for graphical output.
// A frontend might be implemented by SDL, opengl or similar libraries.
type Frontend interface {
	Initialize()
	Draw(emulator emu.Chip8)
	Close()
}
