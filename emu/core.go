package emu

const (
	MemorySize      = 4096
	VramSize        = 64 * 32
	RegistersNumber = 16
	StackSize       = 16
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
		make([]uint16, StackSize, StackSize),
		make([]uint8, RegistersNumber, RegistersNumber),
		make([]uint8, MemorySize, MemorySize),
		make([]uint8, VramSize, VramSize),
		0,
		0,
	}
}
