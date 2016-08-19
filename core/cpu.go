package core

// var instructionSet = map[int]instruction {
// 	instruction{cpu: nil, Name: "NOP", Ticks: 4, Length: 1, handler: handlerFunc(nop)}
// }

// nn = int16
// n = int8
var instructionSetDeclaration = map[int]*instruction{
	0x00: &instruction{Name: "NOP", Ticks: 4, Length: 1, handler: handlerFunc(nop)},
	0x01: &instruction{Name: "LD BC nn", Ticks: 12, Length: 3, handler: handlerFunc(ldBcNn)},
	0x02: &instruction{Name: "LD (BC) A", Ticks: 8, Length: 1, handler: handlerFunc(ldBcpA)},
	0x03: &instruction{Name: "INC BC", Ticks: 8, Length: 1, handler: handlerFunc(incBc)},
	0x04: &instruction{Name: "INC B", Ticks: 4, Length: 1, handler: handlerFunc(incB)},
	0x05: &instruction{Name: "DEC B", Ticks: 4, Length: 1, handler: handlerFunc(decB)},
	0x06: &instruction{Name: "LD B n", Ticks: 8, Length: 2, handler: handlerFunc(ldBn)},
	0x07: &instruction{Name: "RLCA", Ticks: 4, Length: 1, handler: handlerFunc(rlca)},
}

// CPU GameBoy cpu
type CPU struct {
	system         *System
	registers      registers
	ProgramCounter uint16
	StackPointer   uint16
	ticks          uint64
	instructionSet *map[int]*instruction
}

func (cpu *CPU) initialize(system *System) {
	cpu.system = system
	cpu.initializeInstructionset()
}

func (cpu *CPU) initializeInstructionset() error {
	cpu.instructionSet = &instructionSetDeclaration
	for _, instruction := range *cpu.instructionSet {
		instruction.Cpu = cpu
	}
	return nil
}

func (cpu *CPU) next() error {
	cpu.ProgramCounter++
	opcode := cpu.system.mmu.readByte(cpu.ProgramCounter)
	inst := (*cpu.instructionSet)[int(opcode&0xFF>>8)]
	inst.execute()
	return nil
}

func nop(cpu *CPU, _ uint16) {
	cpu.ProgramCounter += 2
}

func ldBcNn(cpu *CPU, parameter uint16) {
	cpu.ProgramCounter += 2
	cpu.registers.writeBC(parameter)
}

func ldBcpA(cpu *CPU, _ uint16) {
	cpu.system.mmu.writeByte(cpu.registers.readBC(), cpu.registers.A)
}

func incBc(cpu *CPU, _ uint16) {
	bc := cpu.registers.readBC()
	bc++
	cpu.registers.writeBC(bc)
}

func incB(cpu *CPU, _ uint16) {
	cpu.registers.B = cpu.registers.aluInc(cpu.registers.B)
}

func decB(cpu *CPU, _ uint16) {
	cpu.registers.B = cpu.registers.aluDec(cpu.registers.B)
}

func ldBn(cpu *CPU, parameter uint16) {
	cpu.registers.B = uint8(parameter & 0x00FF >> 8)
}

func rlca(cpu *CPU, _ uint16) {
	carry := (cpu.registers.A & 0x80) >> 7
	if carry == 1 {
		cpu.registers.F |= carryFlag
	} else {
		cpu.registers.F ^= carryFlag
	}

	cpu.registers.A <<= 1
	cpu.registers.A += carry

	cpu.registers.F ^= negativeFlag | zeroFlag | halfCarryFlag
}
