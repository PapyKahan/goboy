package core

// nn = int16
// n = int8
var instructionSetDeclaration = map[int]*instruction{
	0x00: &instruction{name: "NOP", actionTakenTicks: 4, length: 1, handler: nop},
	0x01: &instruction{name: "LD BC nn", actionTakenTicks: 12, length: 3, handler: ldBcNn},
	0x02: &instruction{name: "LD (BC) A", actionTakenTicks: 8, length: 1, handler: ldBcpA},
	0x03: &instruction{name: "INC BC", actionTakenTicks: 8, length: 1, handler: incBc},
	0x04: &instruction{name: "INC B", actionTakenTicks: 4, length: 1, handler: incB},
	0x05: &instruction{name: "DEC B", actionTakenTicks: 4, length: 1, handler: decB},
	0x06: &instruction{name: "LD B n", actionTakenTicks: 8, length: 2, handler: ldBn},
	0x07: &instruction{name: "RLCA", actionTakenTicks: 4, length: 1, handler: rlca},
	0x08: &instruction{name: "LD (nn) SP", actionTakenTicks: 20, length: 3, handler: ldNnpSp},
	0x09: &instruction{name: "ADD HL BC", actionTakenTicks: 8, length: 1, handler: addHlBc},
	0x0A: &instruction{name: "LD A (BC)", actionTakenTicks: 8, length: 1, handler: ldABcp},
	0x0B: &instruction{name: "DEC BC", actionTakenTicks: 8, length: 1, handler: decBc},
	0x0C: &instruction{name: "INC C", actionTakenTicks: 4, length: 1, handler: incC},
	0x0D: &instruction{name: "DEC C", actionTakenTicks: 4, length: 1, handler: decC},
	0x0E: &instruction{name: "LD C n", actionTakenTicks: 8, length: 2, handler: ldCN},
	0x0F: &instruction{name: "RRCA", actionTakenTicks: 4, length: 1, handler: rrca},
	0x10: &instruction{name: "STOP", actionTakenTicks: 4, length: 2, handler: stop},
	0x11: &instruction{name: "LD DE nn", actionTakenTicks: 12, length: 3, handler: ldDeNn},
	0x12: &instruction{name: "LD (DE) A", actionTakenTicks: 8, length: 1, handler: ldDepA},
	0x13: &instruction{name: "INC DE", actionTakenTicks: 8, length: 1, handler: incDe},
	0x14: &instruction{name: "INC D", actionTakenTicks: 4, length: 1, handler: incD},
	0x15: &instruction{name: "DEC D", actionTakenTicks: 4, length: 1, handler: decD},
	0x16: &instruction{name: "LD D n", actionTakenTicks: 8, length: 2, handler: ldDn},
	0x17: &instruction{name: "RLA", actionTakenTicks: 4, length: 1, handler: rla},
	0x18: &instruction{name: "JR n", actionTakenTicks: 12, length: 2, handler: jrn},
	0x19: &instruction{name: "ADD HL DE", actionTakenTicks: 8, length: 1, handler: addHlDe},
	0x1A: &instruction{name: "LD A (DE)", actionTakenTicks: 8, length: 1, handler: ldADep},
	0x1B: &instruction{name: "DEC DE", actionTakenTicks: 8, length: 1, handler: decDe},
	0x1C: &instruction{name: "INC E", actionTakenTicks: 4, length: 1, handler: incE},
	0x1D: &instruction{name: "DEC E", actionTakenTicks: 4, length: 1, handler: decE},
	0x1E: &instruction{name: "LD E n", actionTakenTicks: 8, length: 2, handler: ldEn},
	0x1F: &instruction{name: "RRA", actionTakenTicks: 4, length: 1, handler: rra},
	0x20: &instruction{name: "JR NZ n", actionTakenTicks: 12, actionNotTakenTicks: 8, length: 2, handler: jrNzn},
}

type cpu struct {
	registers      registers
	ticks          uint64
	stoped         bool
	instructionSet *map[int]*instruction

	// Processing units
	mmu *mmu
	gpu *gpu
	spu *spu
}

func (cpu *cpu) initialize() {
	cpu.initializeInstructionset()
	cpu.mmu = &mmu{}
	cpu.gpu = &gpu{}
	cpu.spu = &spu{}

	cpu.stoped = false
}

func (cpu *cpu) initializeInstructionset() error {
	cpu.instructionSet = &instructionSetDeclaration
	return nil
}

func (cpu *cpu) next() error {
	if cpu.stoped {
		return nil
	}

	opcode := cpu.mmu.readByte(cpu.registers.pc)
	cpu.registers.pc++
	inst := (*cpu.instructionSet)[int(opcode)]

	var parameters uint16

	switch inst.length {
	case 1:
		parameters = 0
	case 2:
		parameters = uint16(cpu.mmu.readByte(cpu.registers.pc))

	case 3:
		parameters = cpu.mmu.readWord(cpu.registers.pc)
	}

	condition := inst.handler(cpu, parameters)
	cpu.registers.pc += uint16(inst.length - 1)

	if condition {
		cpu.ticks += inst.actionTakenTicks
	} else {
		cpu.ticks += inst.actionNotTakenTicks
	}

	return nil
}

func (cpu *cpu) rotateLeftCarry(value byte) byte {
	carry := (value & 0x80) == 0x80

	value <<= 1
	if carry {
		value++
	}

	if carry {
		cpu.registers.F |= carryFlag
	} else {
		cpu.registers.F &^= carryFlag
	}

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F &^= negativeFlag | halfCarryFlag

	return value
}

func (cpu *cpu) rotateLeft(value byte) byte {
	carry := cpu.registers.F&carryFlag == carryFlag

	if value&0x80 == 0x80 {
		cpu.registers.F |= carryFlag
	} else {
		cpu.registers.F &^= carryFlag
	}

	value <<= 1
	if carry {
		value++
	}

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F &^= negativeFlag | halfCarryFlag

	return value
}

func (cpu *cpu) rotateRight(value byte) byte {
	carry := cpu.registers.F&carryFlag == carryFlag

	if value&0x01 == 0x01 {
		cpu.registers.F |= carryFlag
	} else {
		cpu.registers.F &^= carryFlag
	}

	value >>= 1
	if carry {
		value |= 0x80
	}

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F &^= negativeFlag | halfCarryFlag

	return value
}

func (cpu *cpu) aluRotateRightCarry(value byte) byte {
	carry := value&0x01 == 0x01

	if carry {
		cpu.registers.F |= carryFlag
	} else {
		cpu.registers.F &^= carryFlag
	}

	value >>= 1
	if carry {
		value |= 0x80
	}

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F &^= negativeFlag | halfCarryFlag

	return value
}

func (cpu *cpu) aluInc(value byte) byte {
	if value&0x0F == 0x0F {
		cpu.registers.F |= halfCarryFlag
	} else {
		cpu.registers.F &^= halfCarryFlag
	}

	value++

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F &^= negativeFlag

	return value
}

func (cpu *cpu) aluDec(value byte) byte {
	if value&0x0F == 0x0 {
		cpu.registers.F &^= halfCarryFlag
	} else {
		cpu.registers.F |= halfCarryFlag
	}

	value--

	if value == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	cpu.registers.F |= negativeFlag

	return value
}

func (cpu *cpu) addWord(a uint16, b uint16) uint16 {
	if uint(a)+uint(b) > 0xFFFF {
		cpu.registers.F |= carryFlag
	}

	if uint16(a&0x0FFF)+uint16(b&0x0FFF) > 0x0FFF {
		cpu.registers.F |= halfCarryFlag
	}

	cpu.registers.F &^= negativeFlag

	return a + b
}

func (cpu *cpu) relativeJump(value byte) {
	if value&0x80 == 0x80 {
		value--
		value = ^value
		cpu.registers.pc -= uint16(value)
	} else {
		cpu.registers.pc += uint16(value)
	}
}

// Instructions implementation :

func nop(cpu *cpu, _ uint16) bool {
	return true
}

func ldBcNn(cpu *cpu, parameter uint16) bool {
	cpu.registers.writeBC(parameter)
	return true
}

func ldBcpA(cpu *cpu, _ uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readBC(), cpu.registers.A)
	return true
}

func incBc(cpu *cpu, _ uint16) bool {
	bc := cpu.registers.readBC()
	bc++
	cpu.registers.writeBC(bc)
	return true
}

func incB(cpu *cpu, _ uint16) bool {
	cpu.registers.B = cpu.aluInc(cpu.registers.B)
	return true
}

func decB(cpu *cpu, _ uint16) bool {
	cpu.registers.B = cpu.aluDec(cpu.registers.B)
	return true
}

func ldBn(cpu *cpu, parameter uint16) bool {
	cpu.registers.B = uint8(parameter & 0x00FF)
	return true
}

func rlca(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.rotateLeftCarry(cpu.registers.A)
	cpu.registers.F &^= zeroFlag
	return true
}

func ldNnpSp(cpu *cpu, value uint16) bool {
	cpu.mmu.writeWord(value, cpu.registers.sp)
	return true
}

func addHlBc(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	bc := cpu.registers.readBC()
	cpu.registers.writeHL(cpu.addWord(hl, bc))
	return true
}

func ldABcp(cpu *cpu, address uint16) bool {
	cpu.registers.A = cpu.mmu.readByte(cpu.registers.readBC())
	return true
}

func decBc(cpu *cpu, _ uint16) bool {
	value := cpu.registers.readBC()
	value--
	cpu.registers.writeBC(value)
	return true
}

func incC(cpu *cpu, _ uint16) bool {
	cpu.registers.C = cpu.aluInc(cpu.registers.C)
	return true
}

func decC(cpu *cpu, _ uint16) bool {
	cpu.registers.C = cpu.aluDec(cpu.registers.C)
	return true
}

func ldCN(cpu *cpu, value uint16) bool {
	cpu.registers.C = byte(value & 0x00FF)
	return true
}

func rrca(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.aluRotateRightCarry(cpu.registers.A)
	cpu.registers.F &^= zeroFlag
	return true
}

func stop(cpu *cpu, _ uint16) bool {
	cpu.stoped = true
	return true
}

func ldDeNn(cpu *cpu, value uint16) bool {
	cpu.registers.writeDE(value)
	return true
}

func ldDepA(cpu *cpu, _ uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readDE(), cpu.registers.A)
	return true
}

func incDe(cpu *cpu, _ uint16) bool {
	de := cpu.registers.readDE()
	de++
	cpu.registers.writeDE(de)
	return true
}

func incD(cpu *cpu, _ uint16) bool {
	cpu.registers.D = cpu.aluInc(cpu.registers.D)
	return true
}

func decD(cpu *cpu, _ uint16) bool {
	cpu.registers.D = cpu.aluDec(cpu.registers.D)
	return true
}

func ldDn(cpu *cpu, value uint16) bool {
	cpu.registers.D = byte(value & 0x00FF)
	return true
}

func rla(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.rotateLeft(cpu.registers.A)
	cpu.registers.F &^= zeroFlag
	return true
}

func jrn(cpu *cpu, value uint16) bool {
	cpu.relativeJump(byte(value & 0x00FF))
	return true
}

func addHlDe(cpu *cpu, value uint16) bool {
	hl := cpu.registers.readHL()
	de := cpu.registers.readDE()
	cpu.registers.writeHL(cpu.addWord(hl, de))
	return true
}

func ldADep(cpu *cpu, address uint16) bool {
	cpu.registers.A = cpu.mmu.readByte(cpu.registers.readDE())
	return true
}

func decDe(cpu *cpu, _ uint16) bool {
	value := cpu.registers.readDE()
	value--
	cpu.registers.writeDE(value)
	return true
}

func incE(cpu *cpu, _ uint16) bool {
	cpu.registers.E = cpu.aluInc(cpu.registers.E)
	return true
}

func decE(cpu *cpu, _ uint16) bool {
	cpu.registers.E = cpu.aluDec(cpu.registers.E)
	return true
}

func ldEn(cpu *cpu, value uint16) bool {
	cpu.registers.E = byte(value & 0x00FF)
	return true
}

func rra(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.rotateRight(cpu.registers.A)
	cpu.registers.F &^= zeroFlag
	return true
}

func jrNzn(cpu *cpu, value uint16) bool {
	if cpu.registers.F&zeroFlag != zeroFlag {
		cpu.relativeJump(byte(value & 0x00FF))
		return true
	}
	return false
}
