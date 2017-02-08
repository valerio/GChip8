package emu

// The mapping of opcodes to the function implementing it.
var OpcodeMap = map[uint16]OpcodeFunc {
    0x0000: Nop,
}

// Nop does nothing
// This is not an actual opcode, just a placeholder.
func Nop(c8 *Chip8) {
    c8.pc += 2
}

// SetVxToImmediate implements opcode 6XNN.
// It will set NN (8 bit immediate) to the register Vx.
func SetVxToImmediate(c8 *Chip8) {
    x := (c8.opcode & 0x0F00) >> 8
    nn := uint8(c8.opcode & 0x00FF)

    c8.V[x] = nn
    c8.pc += 2
}