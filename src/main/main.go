package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/valep27/GChip8/src/emu"
	"github.com/valep27/GChip8/src/io"
)

func main() {
	var path string
	app := cli.NewApp()

	app.Name = "GChip8"
	app.UsageText = fmt.Sprintf("%s [path]", app.Name)
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "game file path",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {
		args := c.Args()
		if len(args) != 1 {
			return fmt.Errorf("Usage: %s", app.UsageText)
		}

		path := args.Get(0)
		return run(path)
	}
	app.Run(os.Args)
}

func run(path string) error {
	var event *io.KeyEvent

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("cannot open file '%s': %s", path, err)
	}

	chip8 := emu.New()
	chip8.LoadRom(path)

	front := io.NewSdlFrontend()
	input := io.NewSdlInput()
	front.Initialize()
	defer front.Close()

	drawChan := make(chan []uint8)
	go draw(front, drawChan)

	for {
		chip8.Step()
		drawChan <- chip8.GetPixelFrameBuffer()

		for event = input.Poll(); event != nil; event = input.Poll() {

			if event.Key == io.KeyQuit {
				return nil
			}

			chip8.HandleKeyEvent(uint8(event.Key), event.Up)
		}
	}
}

func draw(front io.SdlFrontend, c chan []uint8) {
	for {
		buffer := <-c
		front.Draw(buffer)
	}
}
