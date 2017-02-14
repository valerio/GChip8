package main

import (
	"fmt"
	"time"

	"github.com/valep27/GChip8/emu"
	"github.com/valep27/GChip8/io"
)

func main() {
	fmt.Println("Starting emulator...")

	chip8 := emu.New()
	front := io.NewSdlFrontend()

	input := io.NewSdlInput()

	front.Initialize()

	defer front.Close()

	front.Draw(chip8)

	running := true
	for running {
		k := input.Poll()

		if k == io.KEY_QUIT {
			running = false
		}

		time.Sleep(33 * time.Millisecond)
	}
}
