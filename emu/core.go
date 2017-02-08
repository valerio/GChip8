package emu

const (
	MEMORY_SIZE      = 4096
	VRAM_SIZE        = 64 * 32
	REGISTERS_NUMBER = 16
	STACK_SIZE       = 16
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
		make([]uint16, STACK_SIZE, STACK_SIZE),
		make([]uint8, REGISTERS_NUMBER, REGISTERS_NUMBER),
		make([]uint8, MEMORY_SIZE, MEMORY_SIZE),
		make([]uint8, VRAM_SIZE, VRAM_SIZE),
		0,
		0,
	}
}
