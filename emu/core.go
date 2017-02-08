package emu

const (
	memorySize      = 4096
	vramSize        = 64 * 32
	registersNumber = 16
	stackSize       = 16
)

// Chip8 is the main struct holding all data relevant to the emulator.
// This includes registers (V0 to VF, PC, etc.), ram and framebuffer.
type Chip8 struct {
	I      uint16
	pc     uint16
	sp     uint16
	stack  []uint16
	V      []uint8
	memory []uint8
	vram   []uint8
	keypad []uint8
	delayt uint8
	soundt uint8
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
	}
}
