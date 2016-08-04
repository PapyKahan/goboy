package core

type instruction struct {
	Cpu     *CPU
	handler handlerFunc
	Name    string
	Ticks   uint64
	Length  uint8
}

func (inst *instruction) execute() {
	inst.handler.handle(inst.Cpu, inst.Length)
	inst.Cpu.ticks += inst.Ticks
}

type handlerFunc func(cpu *CPU, parameters uint16)

func (f handlerFunc) handle(cpu *CPU, instructionLength byte) {
	var parameters uint16

	switch instructionLength - 1 {
	case 0:
		f(cpu, 0)
	case 1:
		parameters = uint16(cpu.system.mmu.readByte(cpu.ProgramCounter))
	case 2:
		parameters = cpu.system.mmu.readWord(cpu.ProgramCounter)
	}
	f(cpu, parameters)
	cpu.ProgramCounter += uint16(instructionLength - 1)
}
