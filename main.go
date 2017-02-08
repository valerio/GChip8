package main

import "fmt"
import "github.com/veandco/go-sdl2/sdl"

func main() {
	fmt.Println("Starting emulator...")

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Chip8",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN | sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 800, 600}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(3000)
	sdl.Quit()
}
