package io

import (
	"github.com/valep27/chip8/emu"
	"github.com/veandco/go-sdl2/sdl"
)

// SdlFrontend implements basic drawing using SDL2.
type SdlFrontend struct {
	window *sdl.Window
}

// NewSdlFrontend creates a new uninitialized frontend that uses SDL2.
func NewSdlFrontend() SdlFrontend {
	return SdlFrontend{nil}
}

// Initialize creates the window and sets up any internal state for the frontend.
func (sf *SdlFrontend) Initialize() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Chip8",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}

	sf.window = window
}

// Draw will draw on the window the contents of the emulator framebuffer.
func (sf *SdlFrontend) Draw(emulator emu.Chip8) {
	surface, err := sf.window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 800, 600}
	surface.FillRect(&rect, 0xffff0000)
	sf.window.UpdateSurface()
}

// Close will free any resources, the window and quit the application.
// Best used with defer.
func (sf *SdlFrontend) Close() {
	defer sf.window.Destroy()
	sdl.Delay(3000)
	sdl.Quit()
}
