package core

// System Gameboy system.
type System struct {
	cpu *CPU
	mmu *MMU
	gpu *Gpu
	spu *Spu
}
