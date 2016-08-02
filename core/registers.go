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

func (registers *registers) readBC() uint16 {
	return uint16(registers.B)<<8 + uint16(registers.C)
}

func (registers *registers) writeBC(value uint16) {
	registers.B = byte(value >> 8)
	registers.C = byte(value & 0x00FF)
}
