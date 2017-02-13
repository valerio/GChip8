package io

import "github.com/veandco/go-sdl2/sdl"
import "fmt"

// SdlInput implements basic drawing using SDL2.
type SdlInput struct {
}

// NewSdlInput creates a new uninitialized Input that uses SDL2.
func NewSdlInput() SdlInput {
	return SdlInput{}
}

// Poll polls polls
func (i *SdlInput) Poll() (keyPressed Key) {
	event := sdl.PollEvent()
	keyPressed = KEY_NONE

	switch t := event.(type) {
	case *sdl.QuitEvent:
	case *sdl.MouseMotionEvent:
		fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
	case *sdl.MouseButtonEvent:
		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y)
	case *sdl.KeyUpEvent:
		fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

		switch t.Keysym.Sym {
		case sdl.K_1:
			keyPressed = KEY_1
		case sdl.K_2:
			keyPressed = KEY_2
		case sdl.K_3:
			keyPressed = KEY_3
		case sdl.K_4:
			keyPressed = KEY_4
		case sdl.K_q:
			keyPressed = KEY_5
		case sdl.K_w:
			keyPressed = KEY_6
		case sdl.K_e:
			keyPressed = KEY_7
		case sdl.K_r:
			keyPressed = KEY_8
		case sdl.K_a:
			keyPressed = KEY_9
		case sdl.K_s:
			keyPressed = KEY_A
		case sdl.K_d:
			keyPressed = KEY_B
		case sdl.K_f:
			keyPressed = KEY_C
		case sdl.K_z:
			keyPressed = KEY_D
		case sdl.K_x:
			keyPressed = KEY_0
		case sdl.K_c:
			keyPressed = KEY_E
		case sdl.K_v:
			keyPressed = KEY_F
		case sdl.K_ESCAPE:
			keyPressed = KEY_QUIT
		}
	}

	return
}
