package emu

import "github.com/valep27/GChip8/util"
import "fmt"

const (
	memorySize      = 4096
	vramSize        = 64 * 32
	registersNumber = 16
	stackSize       = 16
)

// Chip8 is the main struct holding all data relevant to the emulator.
// This includes registers (V0 to VF, PC, etc.), ram and framebuffer.
type Chip8 struct {
	I        uint16
	pc       uint16
	sp       uint16
	stack    []uint16
	V        []uint8
	memory   []uint8
	vram     []uint8
	keypad   []uint8
	delayt   uint8
	soundt   uint8
	opcode   uint16
	drawFlag bool
}

// OpcodeFunc is a function that implements an opcode for Chip8
type OpcodeFunc func(*Chip8)

// New initializes basic Chip8 data, but the emulator won't be in a runnable
// state until something is loaded.
func New() Chip8 {
	return Chip8{
		0,
		0,
		0,
		make([]uint16, stackSize, stackSize),
		make([]uint8, registersNumber, registersNumber),
		make([]uint8, memorySize, memorySize),
		make([]uint8, vramSize, vramSize),
		make([]uint8, 16, 16),
		0,
		0,
		0,
		false,
	}
}

// Step executes a single cycle of emulation.
func (c8 *Chip8) Step() {
	// fetch
	opcode := util.CombineBytes(c8.memory[c8.pc+1], c8.memory[c8.pc])
	c8.opcode = opcode

	// decode
	instr, ok := Decode(opcode)

	if ok {
		// exec
		instr(c8)
	} else {
		// opcode not found
		panic(fmt.Sprintf("No instruction for opcode: %v", opcode))
	}

	// update timers
	if c8.delayt > 0 {
		c8.delayt--
	}

	if c8.soundt > 0 {
		if c8.soundt == 1 {
			// TODO beep boop
		}
		c8.soundt--
	}
}
