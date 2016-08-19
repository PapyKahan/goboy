package core

// System Gameboy system.
type System struct {
	cpu *cpu
}

// New creates a new gameboy system
func New() *System {
	system := &System{}
	system.cpu = &cpu{}
	system.cpu.initialize()
	return system
}

// Execute exectute Gameboy system
func (sys *System) Execute() {
	sys.cpu.next()
}
