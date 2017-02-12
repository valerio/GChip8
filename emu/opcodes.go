package emu

import (
	"math/rand"

	"github.com/valep27/GChip8/util"
)

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
func AssignVyToVx(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F
	c8.V[x] = c8.V[y]
	c8.pc += 2
}

// VxOrVy implements opcode 8XY1
// Assigns the value of Vx | Vy to Vx
func VxOrVy(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F
	c8.V[x] = c8.V[x] | c8.V[y]
	c8.pc += 2
}

// VxAndVy implements opcode 8XY2
// Assigns the value of Vx & Vy to Vx
func VxAndVy(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F
	c8.V[x] = c8.V[x] & c8.V[y]
	c8.pc += 2
}

// VxXorVy implements opcode 8XY3
// Assigns the value of Vx xor Vy to Vx
func VxXorVy(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F
	c8.V[x] = c8.V[x] ^ c8.V[y]
	c8.pc += 2
}

// AddVyToVx implements opcode 8XY4
// Math	Vx += Vy	Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
func AddVyToVx(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F

	result, carry := util.CheckedAdd(c8.V[x], c8.V[y])
	c8.V[0xF] = 0
	c8.V[x] = result

	if carry {
		c8.V[0xF] = 1
	}

	c8.pc += 2
}

// SubVyToVx implements opcode 8XY5
// Math	Vx -= Vy	VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func SubVyToVx(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F

	result, borrow := util.CheckedSub(c8.V[x], c8.V[y])
	c8.V[0xF] = 0
	c8.V[x] = result

	if borrow {
		c8.V[0xF] = 1
	}

	c8.pc += 2
}

// ShiftVxRight implements opcode 8XY6
// BitOp	Vx >> 1	Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.[2]
func ShiftVxRight(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F

	lsb := uint8(x) & 1
	c8.V[x] = c8.V[x] >> 1
	c8.V[0xF] = lsb

	c8.pc += 2
}

// SubVxToVy implements opcode 8XY7
// Math	Vx=Vy-Vx	Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
func SubVxToVy(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F

	result, borrow := util.CheckedSub(c8.V[y], c8.V[x])
	c8.V[0xF] = 0
	c8.V[x] = result

	if borrow {
		c8.V[0xF] = 1
	}

	c8.pc += 2
}

// ShiftVxLeft implements opcode 8XYE
// BitOp	Vx << 1	Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.[2]
func ShiftVxLeft(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F

	msb := uint8(x) & 0x80
	c8.V[x] = c8.V[x] << 1
	c8.V[0xF] = msb

	c8.pc += 2
}

// SkipIfVxNotEqualToVy implements opcode 9XY0
// Cond	if(Vx!=Vy)	Skips the next instruction if VX doesn't equal VY.
func SkipIfVxNotEqualToVy(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	y := (c8.opcode >> 4) & 0x000F

	if c8.V[x] != c8.V[y] {
		c8.pc += 4
	} else {
		c8.pc += 2
	}
}

// SetMemoryNNN implements opcode ANNN
// MEM	I = NNN	Sets I to the address NNN.
func SetMemoryNNN(c8 *Chip8) {
	c8.I = c8.opcode & 0x0FFF
	c8.pc += 2
}

// JumpAddrSum implements opcode BNNN
// Flow PC=V0+NNN	Jumps to the address NNN plus V0.
func JumpAddrSum(c8 *Chip8) {
	c8.pc = (c8.opcode & 0x0FFF) + uint16(c8.V[0])
}

// RandToVx implements opcode CXNN
// Rand Vx=rand()&NN	Sets VX to the result of a bitwise and operation on a random number (Typically: 0 to 255) and NN.
func RandToVx(c8 *Chip8) {
	x := (c8.opcode >> 8) & 0x000F
	nn := uint8(c8.opcode)

	c8.V[x] = uint8(rand.Intn(256)) & nn

	c8.pc += 2
}
