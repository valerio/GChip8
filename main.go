package main

import (
	"fmt"

	"github.com/valep27/GChip8/emu"
	"github.com/valep27/GChip8/io"
)

func main() {
	fmt.Println("Starting emulator...")

	chip8 := emu.New()
	front := io.NewSdlFrontend()

	front.Initialize()
	defer front.Close()

	front.Draw(chip8)
}
