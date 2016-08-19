package core

// nn = int16
// n = int8
var instructionSetDeclaration = map[int]*instruction{
	0x00: &instruction{name: "NOP", ticks: 4, length: 1, handler: handlerFunc(nop)},
	0x01: &instruction{name: "LD BC nn", ticks: 12, length: 3, handler: handlerFunc(ldBcNn)},
	0x02: &instruction{name: "LD (BC) A", ticks: 8, length: 1, handler: handlerFunc(ldBcpA)},
	0x03: &instruction{name: "INC BC", ticks: 8, length: 1, handler: handlerFunc(incBc)},
	0x04: &instruction{name: "INC B", ticks: 4, length: 1, handler: handlerFunc(incB)},
	0x05: &instruction{name: "DEC B", ticks: 4, length: 1, handler: handlerFunc(decB)},
	0x06: &instruction{name: "LD B n", ticks: 8, length: 2, handler: handlerFunc(ldBn)},
	0x07: &instruction{name: "RLCA", ticks: 4, length: 1, handler: handlerFunc(rlca)},
}

type cpu struct {
	registers      registers
	ProgramCounter uint16
	StackPointer   uint16
	ticks          uint64
	instructionSet *map[int]*instruction

	// Processing units
	mmu *mmu
	gpu *gpu
	spu *spu
}

func (cpu *cpu) initialize() {
	cpu.initializeInstructionset()
}

func (cpu *cpu) initializeInstructionset() error {
	cpu.instructionSet = &instructionSetDeclaration
	for _, instruction := range *cpu.instructionSet {
		instruction.cpu = cpu
	}
	return nil
}

func (cpu *cpu) next() error {
	opcode := cpu.mmu.readByte(cpu.ProgramCounter)
	cpu.ProgramCounter++
	inst := (*cpu.instructionSet)[int(opcode&0xFF>>8)]
	inst.execute()
	return nil
}

func nop(cpu *cpu, _ uint16) {
}

func ldBcNn(cpu *cpu, parameter uint16) {
	cpu.registers.writeBC(parameter)
}

func ldBcpA(cpu *cpu, _ uint16) {
	cpu.mmu.writeByte(cpu.registers.readBC(), cpu.registers.A)
}

func incBc(cpu *cpu, _ uint16) {
	bc := cpu.registers.readBC()
	bc++
	cpu.registers.writeBC(bc)
}

func incB(cpu *cpu, _ uint16) {
	cpu.registers.B = cpu.registers.aluInc(cpu.registers.B)
}

func decB(cpu *cpu, _ uint16) {
	cpu.registers.B = cpu.registers.aluDec(cpu.registers.B)
}

func ldBn(cpu *cpu, parameter uint16) {
	cpu.registers.B = uint8(parameter & 0x00FF >> 8)
}

func rlca(cpu *cpu, _ uint16) {
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
