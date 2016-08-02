package core

// System Gameboy system.
type System struct {
	cpu *CPU
	mmu *MMU
	gpu *Gpu
	spu *Spu
}

// New creates a new gameboy system
func New() *System {
	system := &System{}
	system.cpu = &CPU{}
	system.cpu.initialize(system)
	return system
}

// Execute exectute Gameboy system
func (sys *System) Execute() {
	sys.cpu.next()
}
