package emu

// Decode maps an opcode to the function that implements it.
func Decode(opcode uint16) (instr OpcodeFunc, ok bool) {
    ok = true

    switch opcode & 0xF000 {
    case 0x0000:
        instr = Nop
    case 0x6000:
        instr = SetVxToImmediate
    default:
        ok = false
    }

    return
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