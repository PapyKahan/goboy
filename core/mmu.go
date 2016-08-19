package core

type mmu struct {
	cartidge [0x8000]byte
	sram     [0x2000]byte
	io       [0x100]byte
	vram     [0x2000]byte
	oam      [0x100]byte
	wram     [0x2000]byte
	hram     [0x80]byte
}

func (mmu *mmu) readByte(address uint16) byte {
	return 0
}

func (mmu *mmu) readWord(address uint16) uint16 {
	return 0
}

func (mmu *mmu) writeByte(address uint16, value byte) {

}

func (mmu *mmu) writeWord(address uint16, value uint16) {

}
