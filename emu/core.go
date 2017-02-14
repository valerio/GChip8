package emu

import "github.com/valep27/GChip8/util"
import "fmt"
import "io/ioutil"

const (
	memorySize      = 4096
	vramSize        = 64 * 32
	registersNumber = 16
	stackSize       = 16
)

// Sprites representing hex numbers from 0 to F
var fontSet = [...]uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

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
	stopped  bool
}

// OpcodeFunc is a function that implements an opcode for Chip8
type OpcodeFunc func(*Chip8)

// New initializes basic Chip8 data, but the emulator won't be in a runnable
// state until something is loaded.
func New() *Chip8 {
	c8 := &Chip8{
		0,
		0x200,
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
		false,
	}

	for i := 0; i < len(fontSet); i++ {
		c8.memory[i] = fontSet[i]
	}

	return c8
}

// LoadRom will load a rom file in memory, starting at address 0x200 (512).
func (c8 *Chip8) LoadRom(path string) {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		panic(fmt.Sprintf("Cannot read file %v, error: %s\n", path, err.Error()))
	}

	for i := 0; i < len(buffer); i++ {
		c8.memory[0x200 + i] = buffer[i]
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

// IsKeyPressed checks whether key 0 to 15 was pressed on the keypad.
func (c8 *Chip8) IsKeyPressed(key uint8) bool {
	return c8.keypad[key] != 0
}
