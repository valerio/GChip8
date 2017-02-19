package io

import (
	"unsafe"

	"github.com/valep27/GChip8/emu"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width        = 64
	height       = 32
	textureDepth = 4
	renderScale  = 4
)

// SdlFrontend implements basic drawing using SDL2.
type SdlFrontend struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	fb       []uint32
}

// NewSdlFrontend creates a new uninitialized frontend that uses SDL2.
func NewSdlFrontend() SdlFrontend {
	return SdlFrontend{nil, nil, make([]uint32, width*height)}
}

// Initialize creates the window and sets up any internal state for the frontend.
func (sf *SdlFrontend) Initialize() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Chip8",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width*renderScale,
		height*renderScale,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)

	if err != nil {
		panic(err)
	}

	sf.window = window
	sf.renderer = renderer
}

// Draw will draw on the window the contents of the emulator framebuffer.
func (sf *SdlFrontend) Draw(emulator *emu.Chip8) {
	pixels := width * height

	for i := 0; i < pixels; i++ {
		if emulator.GetPixelFrameBuffer()[i] == 0 {
			sf.fb[i] = 0
		} else {
			sf.fb[i] = 0xFFFFFFFF
		}
	}

	surface, err := sdl.CreateRGBSurfaceFrom(
		unsafe.Pointer(&sf.fb[0]),
		width,
		height,
		32,
		textureDepth*width,
		0x000000FF,
		0x0000FF00,
		0x00FF0000,
		0xFF000000)

	if err != nil {
		panic(err)
	}

	surface.Lock()
	sf.renderer.Clear()
	txt, err := sf.renderer.CreateTextureFromSurface(surface)
	surface.Unlock()

	if err != nil {
		panic(err)
	}

	sf.renderer.Copy(txt, nil, nil)
	sf.renderer.Present()
}

// Close will free any resources, the window and quit the application.
// Best used with defer.
func (sf *SdlFrontend) Close() {
	defer sf.window.Destroy()
	defer sf.renderer.Destroy()
	sdl.Quit()
}
