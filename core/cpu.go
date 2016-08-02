package core

// var instructionSet = map[int]instruction {
// 	instruction{cpu: nil, Name: "NOP", Ticks: 4, Length: 1, handler: handlerFunc(nop)}
// }

var instructionSet = map[int]instruction{
	0x00: instruction{Name: "NOP", Ticks: 4, Length: 1, handler: handlerFunc(nop)},
	0x01: instruction{Name: "LDBCnn", Ticks: 12, Length: 3, handler: handlerFunc(ldBcNn)},
}

// CPU GameBoy cpu
type CPU struct {
	system         *System
	registers      registers
	ProgramCounter uint16
	StackPointer   uint16
	ticks          uint64
	instructionSet map[int]instruction
}

func (cpu *CPU) initialize(system *System) {
	cpu.system = system
	cpu.initializeInstructionset()
}

func (cpu *CPU) initializeInstructionset() error {
	cpu.instructionSet = instructionSet
	for _, instruction := range cpu.instructionSet {
		instruction.cpu = cpu
	}
	return nil
}

func (cpu *CPU) next() error {
	cpu.ProgramCounter++
	opcode := cpu.system.mmu.readByte(cpu.ProgramCounter)
	inst := cpu.instructionSet[int(opcode&0xFF>>8)]
	inst.execute()
	return nil
}

func nop(cpu *CPU, parameters uint16) {
	cpu.ProgramCounter += 2
}

func ldBcNn(cpu *CPU, parameters uint16) {
	cpu.ProgramCounter += 2
	cpu.registers.writeBC(parameters)
}
