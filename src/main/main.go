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
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("cannot open file '%s': %s", path, err)
	}

	chip8 := emu.New()
	chip8.LoadRom(path)

	front := io.NewSdlFrontend()
	input := io.NewSdlInput()
	front.Initialize()
	defer front.Close()

	for {
		chip8.Step()
		front.Draw(chip8)

		k, up := input.Poll()
		if k == io.KeyQuit {
			return nil
		}

		chip8.HandleKeyEvent(uint8(k), up)
	}
}
