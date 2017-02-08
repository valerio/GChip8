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
	I      uint8
	pc     uint16
	sp     uint8
	stack  []uint16
	V      []uint8
	memory []uint8
	vram   []uint8
	delayt uint8
	soundt uint8
}

func New() Chip8 {
	return Chip8{
		0,
		0,
		0,
		make([]uint16, stackSize, stackSize),
		make([]uint8, registersNumber, registersNumber),
		make([]uint8, memorySize, memorySize),
		make([]uint8, vramSize, vramSize),
		0,
		0,
	}
}
