package io

import "github.com/veandco/go-sdl2/sdl"

// SdlInput implements basic drawing using SDL2.
type SdlInput struct {
}

// NewSdlInput creates a new uninitialized Input that uses SDL2.
func NewSdlInput() SdlInput {
	return SdlInput{}
}

func mapSymbolToKey(keycode sdl.Keycode) Key {
	keyPressed := KeyNone

	switch keycode {
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

	return keyPressed
}

// Poll polls for an input event and return the key that was pressed (mapped to Chip8 keys)
// and whether the event was for a key up or down event.
func (i *SdlInput) Poll() *KeyEvent {
	event := sdl.PollEvent()

	if event == nil {
		return nil
	}

	switch t := event.(type) {
	case *sdl.KeyDownEvent:
		return &KeyEvent{mapSymbolToKey(t.Keysym.Sym), false }
	case *sdl.KeyUpEvent:
		return &KeyEvent{mapSymbolToKey(t.Keysym.Sym), true }
	}

	return &KeyEvent{KeyNone, false }
}
