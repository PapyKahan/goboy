package core

func init() {

}

// CPU GameBoy cpu
type CPU struct {
	system         *System
	registers      registers
	ProgramCounter uint16
	StackPointer   uint16
	ticks          uint64
	m              uint64
	instructionSet map[byte]instruction
}

type instruction struct {
	Exec   func(cpu *CPU)
	Name   string
	Ticks  uint64
	Length uint8
}

// New creates a new instance of Gameboy's CPU
func New(system *System) *CPU {
	cpu := &CPU{}

	cpu.system = system
	cpu.instructionSet = initializeInstructionset()

	return cpu
}

func nop(cpu *CPU) {
	cpu.StackPointer += 2
}

func initializeInstructionset() map[byte]instruction {
	instructionSet := make(map[byte]instruction)

	instructionSet[0x0] = instruction{
		Name:   "NOP",
		Ticks:  4,
		Length: 1,
		Exec:   nop,
	}

	return instructionSet
}
