package core

type instruction struct {
	name                string
	actionTakenTicks    uint64
	actionNotTakenTicks uint64
	length              uint8
	handler             func(cpu *cpu, parameters uint16) bool
	subInstructionSet   *map[int]*instruction
}
