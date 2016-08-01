package core

type registers struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F byte
}

func (registers *registers) readAB() uint16 {
	return uint16(registers.A)<<8 + uint16(registers.B)
}

func (registers *registers) writeAB(value uint16) {
	registers.A = byte(value >> 8)
	registers.B = byte(value & 0xFF00)
}
