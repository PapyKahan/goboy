package core

type instruction struct {
	cpu     *cpu
	handler handlerFunc
	name    string
	ticks   uint64
	length  uint8
}

func (inst *instruction) execute() {
	inst.handler.handle(inst.cpu, inst.length)
	inst.cpu.ticks += inst.ticks
}

type handlerFunc func(cpu *cpu, parameters uint16)

func (f handlerFunc) handle(cpu *cpu, instructionLength byte) {
	var parameters uint16

	switch instructionLength {
	case 1:
		parameters = 0
	case 2:
		parameters = uint16(cpu.mmu.readByte(cpu.ProgramCounter))

	case 3:
		parameters = cpu.mmu.readWord(cpu.ProgramCounter)
	}
	f(cpu, parameters)
	cpu.ProgramCounter += uint16(instructionLength - 1)
}
