package main

import (
	"fmt"
	"time"

	"github.com/valep27/GChip8/emu"
	"github.com/valep27/GChip8/io"
	"os"
)

func main() {
	fmt.Println("Starting emulator...")
	path := ""
	if (len(os.Args) == 2) {
		path = os.Args[1]
	}

	chip8 := emu.New()
	chip8.LoadRom(path)

	front := io.NewSdlFrontend()
	input := io.NewSdlInput()
	front.Initialize()
	defer front.Close()


	running := true
	for running {
		chip8.Step()

		front.Draw(chip8)
		
		k := input.Poll()
		if k == io.KEY_QUIT {
			running = false
		}
		time.Sleep(33 * time.Millisecond)
	}
}
