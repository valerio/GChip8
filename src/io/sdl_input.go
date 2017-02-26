package io

import "github.com/veandco/go-sdl2/sdl"

// SdlInput implements basic drawing using SDL2.
type SdlInput struct {
}

// NewSdlInput creates a new uninitialized Input that uses SDL2.
func NewSdlInput() SdlInput {
	return SdlInput{}
}

// Poll polls for an input event and return the key that was pressed (mapped to Chip8 keys)
// and whether the event was for a key up or down event.
func (i *SdlInput) Poll() (keyPressed Key, up bool) {
	event := sdl.PollEvent()
	up = true
	keyPressed = KeyNone

	switch t := event.(type) {
	case *sdl.QuitEvent:
	case *sdl.MouseMotionEvent:
		// fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
	case *sdl.MouseButtonEvent:
		// fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		// fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
		// 	t.Timestamp, t.Type, t.Which, t.X, t.Y)
	case *sdl.KeyDownEvent:
		up = false
	case *sdl.KeyUpEvent:
		// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
		// 	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

		switch t.Keysym.Sym {
		case sdl.K_1:
			keyPressed = Key1
		case sdl.K_2:
			keyPressed = Key2
		case sdl.K_3:
			keyPressed = Key3
		case sdl.K_4:
			keyPressed = Key4
		case sdl.K_q:
			keyPressed = Key5
		case sdl.K_w:
			keyPressed = Key6
		case sdl.K_e:
			keyPressed = Key7
		case sdl.K_r:
			keyPressed = Key8
		case sdl.K_a:
			keyPressed = Key9
		case sdl.K_s:
			keyPressed = KeyA
		case sdl.K_d:
			keyPressed = KeyB
		case sdl.K_f:
			keyPressed = KeyC
		case sdl.K_z:
			keyPressed = KeyD
		case sdl.K_x:
			keyPressed = Key0
		case sdl.K_c:
			keyPressed = KeyE
		case sdl.K_v:
			keyPressed = KeyF
		case sdl.K_ESCAPE:
			keyPressed = KeyQuit
		}
	}

	return
}
