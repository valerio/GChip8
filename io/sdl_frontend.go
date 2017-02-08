package io

import (
	"github.com/valep27/chip8/emu"
	"github.com/veandco/go-sdl2/sdl"
)

// SdlFrontend implements basic drawing using SDL2.
type SdlFrontend struct {
	window *sdl.Window
}

func NewSdlFrontend() SdlFrontend {
	return SdlFrontend{nil}
}

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

func (sf *SdlFrontend) Draw(emulator emu.Chip8) {
	surface, err := sf.window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 800, 600}
	surface.FillRect(&rect, 0xffff0000)
	sf.window.UpdateSurface()
}

func (sf *SdlFrontend) Close() {
	defer sf.window.Destroy()
	sdl.Delay(3000)
	sdl.Quit()
}
