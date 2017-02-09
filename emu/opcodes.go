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

// ClearScreen implements opcode 00E0.
// Resets the screen pixel values
func ClearScreen(c8 *Chip8) {
	for i := 0; i < len(c8.vram); i++ {
		c8.vram[i] = 0
	}
	c8.pc += 2
}

// ReturnFromSub implements opcode 00EE.
// Returns from a subroutine, meaning it will set the PC to the last stack value.
func ReturnFromSub(c8 *Chip8) { 
    c8.pc = c8.stack[c8.sp]
    c8.sp++
}

// JumpAddr implements opcode 1NNN.
// Sets the program counter to NNN.
func JumpAddr(c8 *Chip8) {
	c8.pc = c8.opcode & 0x0FFF
}

// CallSubAtNNN implements opcode 2NNN.
// It will call the subroutine at address NNN, i.e. move the PC to it.
func CallSubAtNNN(c8 *Chip8) {
	c8.stack[c8.sp] = c8.pc
	c8.sp--
	c8.pc = c8.opcode & 0x0FFF
}

// SkipIfVxEqualToNN implements opcode 3XNN.
// It will skip the next instruction if Vx == NN.
func SkipIfVxEqualToNN(c8 *Chip8) {
    x := (c8.opcode >> 8) & 0x000F
    nn := c8.opcode & 0x00FF

    if c8.V[x] == uint8(nn) {
        c8.pc += 4
    } else {
        c8.pc += 2
    }
}

// SkipIfVxNotEqualToNN implements opcode 4XNN.
// It will skip the next instruction if Vx != NN.
func SkipIfVxNotEqualToNN(c8 *Chip8) {
    x := (c8.opcode >> 8) & 0x000F
    nn := c8.opcode & 0x00FF

    if c8.V[x] != uint8(nn) {
        c8.pc += 4
    } else {
        c8.pc += 2
    }
}

// SkipIfVxEqualToVy implements opcode 5XY0.
// It will skip the next instruction if Vx == Vy.
func SkipIfVxEqualToVy(c8 *Chip8) {
    x := (c8.opcode >> 8) & 0x000F
    y := (c8.opcode >> 4) & 0x000F

    if c8.V[x] == c8.V[y] {
        c8.pc += 4
    } else {
        c8.pc += 2
    }
}

// AddNNToVx implements opcode 7XNN
// It will add NN to the Vx register
func AddNNToVx(c8 *Chip8) {
    x := (c8.opcode >> 8) & 0x000F
    nn := uint8(c8.opcode & 0x00FF)
    c8.V[x] += nn
    c8.pc += 2
}

// AssignVyToVx implements opcode 8XY0
// Assigns the value of Vy to Vx
func AssignVyToVx(c8 *Chip8)  {
    x := (c8.opcode >> 8) & 0x000F
    y := (c8.opcode >> 4) & 0x000F
    c8.V[x] = c8.V[y]
    c8.pc += 2
}

// VxOrVy implements opcode 8XY1
// Assigns the value of Vx | Vy to Vx
func VxOrVy(c8 *Chip8)  {
    x := (c8.opcode >> 8) & 0x000F
    y := (c8.opcode >> 4) & 0x000F
    c8.V[x] = c8.V[x] | c8.V[y]
    c8.pc += 2
}

// VxAndVy implements opcode 8XY2
// Assigns the value of Vx & Vy to Vx
func VxAndVy(c8 *Chip8)  {
    x := (c8.opcode >> 8) & 0x000F
    y := (c8.opcode >> 4) & 0x000F
    c8.V[x] = c8.V[x] & c8.V[y]
    c8.pc += 2
}

// VxXorVy implements opcode 8XY3
// Assigns the value of Vx xor Vy to Vx
func VxXorVy(c8 *Chip8)  {
    x := (c8.opcode >> 8) & 0x000F
    y := (c8.opcode >> 4) & 0x000F
    c8.V[x] = c8.V[x] ^ c8.V[y]
    c8.pc += 2
}