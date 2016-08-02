package core

type instruction struct {
	cpu     *CPU
	handler handlerFunc
	Name    string
	Ticks   uint64
	Length  uint8
}

func (inst *instruction) execute() {
	inst.handler.handle(inst.cpu, inst.Length)
	inst.cpu.ticks += inst.Ticks
}

type handlerFunc func(cpu *CPU, parameters uint16)

func (f handlerFunc) handle(cpu *CPU, instructionLength byte) {
	var parameters uint16

	switch instructionLength {
	case 0:
		f(cpu, 0)
	case 1:
		parameters = uint16(cpu.system.mmu.readByte(cpu.ProgramCounter))
	case 2:
		parameters = cpu.system.mmu.readWord(cpu.ProgramCounter)
	}

	f(cpu, parameters)
}
