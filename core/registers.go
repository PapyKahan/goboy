package core

type flags byte

const (
	zeroFlag      = 1 << 7
	negativeFlag  = 1 << 6
	halfCarryFlag = 1 << 5
	carryFlag     = 1 << 4
)

type registers struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F flags
}

func (registers *registers) readBC() uint16 {
	return uint16(registers.B)<<8 + uint16(registers.C)
}

func (registers *registers) writeBC(value uint16) {
	registers.B = byte(value >> 8)
	registers.C = byte(value & 0x00FF)
}

func (registers *registers) aluInc(value byte) byte {
	if value&0x0F == 0 {
		registers.F |= halfCarryFlag
	} else {
		registers.F ^= halfCarryFlag
	}

	value++

	if value == 0 {
		registers.F |= zeroFlag
	} else {
		registers.F ^= zeroFlag
	}

	registers.F ^= negativeFlag

	return value
}

func (registers *registers) aluDec(value byte) byte {
	if value&0x0F == 0 {
		registers.F ^= halfCarryFlag
	} else {
		registers.F |= halfCarryFlag
	}

	value--

	if value == 0 {
		registers.F |= zeroFlag
	} else {
		registers.F ^= zeroFlag
	}

	registers.F |= negativeFlag

	return value
}
