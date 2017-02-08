package emu

// The mapping of opcodes to the function implementing it.
var OpcodeMap = map[uint16]OpcodeFunc {
    0x0000: Nop,
}

// Nop does nothing
// This is not an actual opcode, just a placeholder.
func Nop(c8 *Chip8) {
    
}