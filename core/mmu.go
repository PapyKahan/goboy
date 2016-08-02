package core

// MMU Gameboy's memory mapping unit
type MMU struct {
}

func (mmu *MMU) readByte(address uint16) byte {
	return 0
}

func (mmu *MMU) readWord(address uint16) uint16 {
	return 0
}

func writeByte(value byte, address uint16) {

}

func writeWord(value uint16, address uint16) {

}
