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

	pc uint16
	sp uint16
}

func (registers *registers) readBC() uint16 {
	return uint16(registers.B)<<8 + uint16(registers.C)
}

func (registers *registers) writeBC(value uint16) {
	registers.B = byte(value >> 8)
	registers.C = byte(value & 0x00FF)
}

func (registers *registers) readHL() uint16 {
	return uint16(registers.H)<<8 + uint16(registers.L)
}

func (registers *registers) writeHL(value uint16) {
	registers.H = byte(value >> 8)
	registers.L = byte(value & 0x00FF)
}

func (registers *registers) writeDE(value uint16) {
	registers.D = byte(value >> 8)
	registers.E = byte(value & 0x00FF)
}

func (registers *registers) readDE() uint16 {
	return uint16(registers.D)<<8 + uint16(registers.E)
}
