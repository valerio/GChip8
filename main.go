package main

import (
	"fmt"

	"github.com/valep27/chip8/io"
	"github.com/valep27/chip8/emu"
)

func main() {
	fmt.Println("Starting emulator...")

	chip8 := emu.New()
	front := io.NewSdlFrontend()

	front.Initialize()
	defer front.Close()

	front.Draw(chip8)
}
