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
	0x21: &instruction{name: "LD HL nn", actionTakenTicks: 12, length: 3, handler: ldHlNn},
	0x22: &instruction{name: "LD (HL+) A", actionTakenTicks: 8, length: 1, handler: ldHlpAIncHl},
	0x23: &instruction{name: "INC HL", actionTakenTicks: 8, length: 1, handler: incHl},
	0x24: &instruction{name: "INC H", actionTakenTicks: 4, length: 1, handler: incH},
	0x25: &instruction{name: "DEC H", actionTakenTicks: 4, length: 1, handler: decH},
	0x26: &instruction{name: "LD H n", actionTakenTicks: 8, length: 2, handler: ldHn},
	0x27: &instruction{name: "DAA", actionTakenTicks: 4, length: 1, handler: daa},
	0x28: &instruction{name: "JR Z n", actionTakenTicks: 12, actionNotTakenTicks: 8, length: 2, handler: jrZn},
	0x29: &instruction{name: "ADD HL HL", actionTakenTicks: 8, length: 1, handler: addHlHl},
	0x2A: &instruction{name: "LD A (HL+)", actionTakenTicks: 8, length: 1, handler: ldAHlpIncHl},
	0x2B: &instruction{name: "DEC HL", actionTakenTicks: 8, length: 1, handler: decHl},
	0x2C: &instruction{name: "INC L", actionTakenTicks: 4, length: 1, handler: incL},
	0x2D: &instruction{name: "DEC L", actionTakenTicks: 4, length: 1, handler: decL},
	0x2E: &instruction{name: "LD L n", actionTakenTicks: 8, length: 2, handler: ldLn},
	0x2F: &instruction{name: "CPL", actionTakenTicks: 4, length: 1, handler: cpl},
	0x30: &instruction{name: "JR NC n", actionTakenTicks: 12, actionNotTakenTicks: 8, length: 2, handler: jrNcn},
	0x31: &instruction{name: "LD SP nn", actionTakenTicks: 12, length: 3, handler: ldSpNn},
	0x32: &instruction{name: "LD (HL-) A", actionTakenTicks: 8, length: 1, handler: ldHlpADecHl},
	0x33: &instruction{name: "INC SP", actionTakenTicks: 8, length: 1, handler: incSp},
	0x34: &instruction{name: "INC (HL)", actionTakenTicks: 12, length: 1, handler: incHlp},
	0x35: &instruction{name: "DEC (HL)", actionTakenTicks: 12, length: 1, handler: decHlp},
	0x36: &instruction{name: "LD (HL), n", actionTakenTicks: 12, length: 2, handler: ldHlpn},
	0x37: &instruction{name: "SCF", actionTakenTicks: 4, length: 1, handler: scf},
	0x38: &instruction{name: "JR C n", actionTakenTicks: 12, actionNotTakenTicks: 8, length: 2, handler: jrCn},
	0x39: &instruction{name: "ADD HL SP", actionTakenTicks: 8, length: 1, handler: addHlSp},
	0x3A: &instruction{name: "LD A (HL-)", actionTakenTicks: 8, length: 1, handler: ldAHlpDecHl},
	0x3B: &instruction{name: "DEC SP", actionTakenTicks: 8, length: 1, handler: decSp},
	0x3C: &instruction{name: "INC A", actionTakenTicks: 4, length: 1, handler: incA},
	0x3D: &instruction{name: "DEC A", actionTakenTicks: 4, length: 1, handler: decA},
	0x3E: &instruction{name: "LD A n", actionTakenTicks: 8, length: 2, handler: ldAn},
	0x3F: &instruction{name: "CCF", actionTakenTicks: 4, length: 1, handler: ccf},
	0x40: &instruction{name: "LD B B", actionTakenTicks: 4, length: 1, handler: ldBb},
	0x41: &instruction{name: "LD B C", actionTakenTicks: 4, length: 1, handler: ldBc},
	0x42: &instruction{name: "LD B D", actionTakenTicks: 4, length: 1, handler: ldBd},
	0x43: &instruction{name: "LD B E", actionTakenTicks: 4, length: 1, handler: ldBe},
	0x44: &instruction{name: "LD B H", actionTakenTicks: 4, length: 1, handler: ldBh},
	0x45: &instruction{name: "LD B L", actionTakenTicks: 4, length: 1, handler: ldBl},
	0x46: &instruction{name: "LD B (HL)", actionTakenTicks: 8, length: 1, handler: ldBhlp},
	0x47: &instruction{name: "LD B A", actionTakenTicks: 4, length: 1, handler: ldBa},
	0x48: &instruction{name: "LD C B", actionTakenTicks: 4, length: 1, handler: ldCb},
	0x49: &instruction{name: "LD C C", actionTakenTicks: 4, length: 1, handler: ldCc},
	0x4A: &instruction{name: "LD C D", actionTakenTicks: 4, length: 1, handler: ldCd},
	0x4B: &instruction{name: "LD C E", actionTakenTicks: 4, length: 1, handler: ldCe},
	0x4C: &instruction{name: "LD C H", actionTakenTicks: 4, length: 1, handler: ldCh},
	0x4D: &instruction{name: "LD C L", actionTakenTicks: 4, length: 1, handler: ldCl},
	0x4E: &instruction{name: "LD C (HL)", actionTakenTicks: 8, length: 1, handler: ldChlp},
	0x4F: &instruction{name: "LD C A", actionTakenTicks: 4, length: 1, handler: ldCa},
	0x50: &instruction{name: "LD D B", actionTakenTicks: 4, length: 1, handler: ldDb},
	0x51: &instruction{name: "LD D C", actionTakenTicks: 4, length: 1, handler: ldDc},
	0x52: &instruction{name: "LD D D", actionTakenTicks: 4, length: 1, handler: ldDd},
	0x53: &instruction{name: "LD D E", actionTakenTicks: 4, length: 1, handler: ldDe},
	0x54: &instruction{name: "LD D H", actionTakenTicks: 4, length: 1, handler: ldDh},
	0x55: &instruction{name: "LD D L", actionTakenTicks: 4, length: 1, handler: ldDl},
	0x56: &instruction{name: "LD D (HL)", actionTakenTicks: 8, length: 1, handler: ldDhlp},
	0x57: &instruction{name: "LD D A", actionTakenTicks: 4, length: 1, handler: ldDa},
	0x58: &instruction{name: "LD E B", actionTakenTicks: 4, length: 1, handler: ldEb},
	0x59: &instruction{name: "LD E C", actionTakenTicks: 4, length: 1, handler: ldEc},
	0x5A: &instruction{name: "LD E D", actionTakenTicks: 4, length: 1, handler: ldEd},
	0x5B: &instruction{name: "LD E E", actionTakenTicks: 4, length: 1, handler: ldEe},
	0x5C: &instruction{name: "LD E H", actionTakenTicks: 4, length: 1, handler: ldEh},
	0x5D: &instruction{name: "LD E L", actionTakenTicks: 4, length: 1, handler: ldEl},
	0x5E: &instruction{name: "LD E (HL)", actionTakenTicks: 8, length: 1, handler: ldEhlp},
	0x5F: &instruction{name: "LD E A", actionTakenTicks: 4, length: 1, handler: ldEa},
	0x60: &instruction{name: "LD H B", actionTakenTicks: 4, length: 1, handler: ldHb},
	0x61: &instruction{name: "LD H C", actionTakenTicks: 4, length: 1, handler: ldHc},
	0x62: &instruction{name: "LD H D", actionTakenTicks: 4, length: 1, handler: ldHd},
	0x63: &instruction{name: "LD H E", actionTakenTicks: 4, length: 1, handler: ldHe},
	0x64: &instruction{name: "LD H H", actionTakenTicks: 4, length: 1, handler: ldHh},
	0x65: &instruction{name: "LD H L", actionTakenTicks: 4, length: 1, handler: ldHl},
	0x66: &instruction{name: "LD H (HL)", actionTakenTicks: 8, length: 1, handler: ldHhlp},
	0x67: &instruction{name: "LD H A", actionTakenTicks: 4, length: 1, handler: ldHa},
	0x68: &instruction{name: "LD L B", actionTakenTicks: 4, length: 1, handler: ldLb},
	0x69: &instruction{name: "LD L C", actionTakenTicks: 4, length: 1, handler: ldLc},
	0x6A: &instruction{name: "LD L D", actionTakenTicks: 4, length: 1, handler: ldLd},
	0x6B: &instruction{name: "LD L E", actionTakenTicks: 4, length: 1, handler: ldLe},
	0x6C: &instruction{name: "LD L H", actionTakenTicks: 4, length: 1, handler: ldLh},
	0x6D: &instruction{name: "LD L L", actionTakenTicks: 4, length: 1, handler: ldLl},
	0x6E: &instruction{name: "LD L (HL)", actionTakenTicks: 8, length: 1, handler: ldLhlp},
	0x6F: &instruction{name: "LD L A", actionTakenTicks: 4, length: 1, handler: ldLa},
	0x70: &instruction{name: "LD (HL) B", actionTakenTicks: 8, length: 1, handler: ldHlpB},
	0x71: &instruction{name: "LD (HL) C", actionTakenTicks: 8, length: 1, handler: ldHlpC},
	0x72: &instruction{name: "LD (HL) D", actionTakenTicks: 8, length: 1, handler: ldHlpD},
	0x73: &instruction{name: "LD (HL) E", actionTakenTicks: 8, length: 1, handler: ldHlpE},
	0x74: &instruction{name: "LD (HL) H", actionTakenTicks: 8, length: 1, handler: ldHlpH},
	0x75: &instruction{name: "LD (HL) L", actionTakenTicks: 8, length: 1, handler: ldHlpL},
	0x76: &instruction{name: "HALT", actionTakenTicks: 4, length: 1, handler: halt},
	0x77: &instruction{name: "LD (HL) A", actionTakenTicks: 8, length: 1, handler: ldHlpA},
}

type cpu struct {
	registers      registers
	ticks          uint64
	stoped         bool
	halt           bool
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
	cpu.halt = false
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

func ldHlNn(cpu *cpu, value uint16) bool {
	cpu.registers.writeHL(value)
	return true
}

func ldHlpAIncHl(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	cpu.mmu.writeByte(hl, cpu.registers.A)
	hl++
	cpu.registers.writeHL(hl)
	return true
}

func incHl(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	hl++
	cpu.registers.writeHL(hl)
	return true
}

func incH(cpu *cpu, _ uint16) bool {
	cpu.registers.H = cpu.aluInc(cpu.registers.H)
	return true
}

func decH(cpu *cpu, _ uint16) bool {
	cpu.registers.H = cpu.aluDec(cpu.registers.H)
	return true
}

func ldHn(cpu *cpu, value uint16) bool {
	cpu.registers.H = byte(value & 0x00FF)
	return true
}

func daa(cpu *cpu, _ uint16) bool {
	correctionFactor := uint16(cpu.registers.A)

	if cpu.registers.F&negativeFlag != negativeFlag {
		if (cpu.registers.F&halfCarryFlag == halfCarryFlag) || (correctionFactor&0x0F > 0x09) {
			correctionFactor += 0x06
		}

		if (cpu.registers.F&carryFlag == carryFlag) || (correctionFactor > 0x9F) {
			correctionFactor += 0x60
		}
	} else {
		if cpu.registers.F&halfCarryFlag == halfCarryFlag {
			correctionFactor = (correctionFactor - 0x06) & 0xFF
		}

		if cpu.registers.F&carryFlag == carryFlag {
			correctionFactor -= 0x60
		}
	}

	cpu.registers.F &^= halfCarryFlag

	if correctionFactor&0x100 == 0x100 {
		cpu.registers.F |= carryFlag
	}

	cpu.registers.A = byte(correctionFactor & 0xFF)

	if cpu.registers.A == 0 {
		cpu.registers.F |= zeroFlag
	} else {
		cpu.registers.F &^= zeroFlag
	}

	return true
}

func jrZn(cpu *cpu, value uint16) bool {
	if cpu.registers.F&zeroFlag == zeroFlag {
		cpu.relativeJump(byte(value & 0x00FF))
		return true
	}
	return false
}

func addHlHl(cpu *cpu, value uint16) bool {
	hl1 := cpu.registers.readHL()
	hl2 := hl1
	cpu.registers.writeHL(cpu.addWord(hl1, hl2))
	return true
}

func ldAHlpIncHl(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	cpu.registers.A = cpu.mmu.readByte(hl)
	hl++
	cpu.registers.writeHL(hl)
	return true
}

func decHl(cpu *cpu, _ uint16) bool {
	value := cpu.registers.readHL()
	value--
	cpu.registers.writeHL(value)
	return true
}

func incL(cpu *cpu, _ uint16) bool {
	cpu.registers.L = cpu.aluInc(cpu.registers.L)
	return true
}

func decL(cpu *cpu, _ uint16) bool {
	cpu.registers.L = cpu.aluDec(cpu.registers.L)
	return true
}

func ldLn(cpu *cpu, value uint16) bool {
	cpu.registers.L = byte(value & 0x00FF)
	return true
}

func cpl(cpu *cpu, value uint16) bool {
	cpu.registers.A = cpu.registers.A ^ 0xFF
	cpu.registers.F |= negativeFlag | halfCarryFlag
	return true
}

func jrNcn(cpu *cpu, value uint16) bool {
	if cpu.registers.F&carryFlag != carryFlag {
		cpu.relativeJump(byte(value & 0x00FF))
		return true
	}
	return false
}

func ldSpNn(cpu *cpu, value uint16) bool {
	cpu.registers.sp = value
	return true
}

func ldHlpADecHl(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	cpu.mmu.writeByte(hl, cpu.registers.A)
	hl--
	cpu.registers.writeHL(hl)
	return true
}

func incSp(cpu *cpu, _ uint16) bool {
	cpu.registers.sp++
	return true
}

func incHlp(cpu *cpu, _ uint16) bool {
	address := cpu.registers.readHL()
	value := cpu.mmu.readByte(address)
	cpu.mmu.writeByte(address, cpu.aluInc(value))
	return true
}

func decHlp(cpu *cpu, _ uint16) bool {
	address := cpu.registers.readHL()
	value := cpu.mmu.readByte(address)
	cpu.mmu.writeByte(address, cpu.aluDec(value))
	return true
}

func ldHlpn(cpu *cpu, value uint16) bool {
	address := cpu.registers.readHL()
	cpu.mmu.writeByte(address, byte(value&0x00FF))
	return true
}

func scf(cpu *cpu, _ uint16) bool {
	cpu.registers.F |= carryFlag
	cpu.registers.F &^= negativeFlag | halfCarryFlag
	return true
}

func jrCn(cpu *cpu, value uint16) bool {
	if cpu.registers.F&carryFlag == carryFlag {
		cpu.relativeJump(byte(value & 0x00FF))
		return true
	}
	return false
}

func addHlSp(cpu *cpu, value uint16) bool {
	hl := cpu.registers.readHL()
	cpu.registers.writeHL(cpu.addWord(hl, cpu.registers.sp))
	return true
}

func ldAHlpDecHl(cpu *cpu, _ uint16) bool {
	hl := cpu.registers.readHL()
	cpu.registers.A = cpu.mmu.readByte(hl)
	hl--
	cpu.registers.writeHL(hl)
	return true
}

func decSp(cpu *cpu, _ uint16) bool {
	cpu.registers.sp--
	return true
}

func incA(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.aluInc(cpu.registers.A)
	return true
}

func decA(cpu *cpu, _ uint16) bool {
	cpu.registers.A = cpu.aluDec(cpu.registers.A)
	return true
}

func ldAn(cpu *cpu, value uint16) bool {
	cpu.registers.A = byte(value & 0x00FF)
	return true
}

func ccf(cpu *cpu, value uint16) bool {
	if (cpu.registers.F & carryFlag) == carryFlag {
		cpu.registers.F &^= carryFlag
	} else {
		cpu.registers.F |= carryFlag
	}
	cpu.registers.F &^= negativeFlag | halfCarryFlag
	return true
}

func ldBb(cpu *cpu, value uint16) bool {
	return true
}

func ldBc(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.C
	return true
}

func ldBd(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.D
	return true
}

func ldBe(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.E
	return true
}

func ldBh(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.H
	return true
}

func ldBl(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.L
	return true
}

func ldBhlp(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldBa(cpu *cpu, value uint16) bool {
	cpu.registers.B = cpu.registers.A
	return true
}

func ldCb(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.B
	return true
}

func ldCc(cpu *cpu, value uint16) bool {
	return true
}

func ldCd(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.D
	return true
}

func ldCe(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.E
	return true
}

func ldCh(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.H
	return true
}

func ldCl(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.L
	return true
}

func ldChlp(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldCa(cpu *cpu, value uint16) bool {
	cpu.registers.C = cpu.registers.A
	return true
}

func ldDb(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.B
	return true
}

func ldDc(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.C
	return true
}

func ldDd(cpu *cpu, value uint16) bool {
	return true
}

func ldDe(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.E
	return true
}

func ldDh(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.H
	return true
}

func ldDl(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.L
	return true
}

func ldDhlp(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldDa(cpu *cpu, value uint16) bool {
	cpu.registers.D = cpu.registers.A
	return true
}

func ldEb(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.B
	return true
}

func ldEc(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.C
	return true
}

func ldEd(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.D
	return true
}

func ldEe(cpu *cpu, value uint16) bool {
	return true
}

func ldEh(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.H
	return true
}

func ldEl(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.L
	return true
}

func ldEhlp(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldEa(cpu *cpu, value uint16) bool {
	cpu.registers.E = cpu.registers.A
	return true
}

//--

func ldHb(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.B
	return true
}

func ldHc(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.C
	return true
}

func ldHd(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.D
	return true
}

func ldHe(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.E
	return true
}

func ldHh(cpu *cpu, value uint16) bool {
	return true
}

func ldHl(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.L
	return true
}

func ldHhlp(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldHa(cpu *cpu, value uint16) bool {
	cpu.registers.H = cpu.registers.A
	return true
}

func ldLb(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.B
	return true
}

func ldLc(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.C
	return true
}

func ldLd(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.D
	return true
}

func ldLe(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.E
	return true
}

func ldLh(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.H
	return true
}

func ldLl(cpu *cpu, value uint16) bool {
	return true
}

func ldLhlp(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.mmu.readByte(cpu.registers.readHL())
	return true
}

func ldLa(cpu *cpu, value uint16) bool {
	cpu.registers.L = cpu.registers.A
	return true
}

func ldHlpB(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.B)
	return true
}

func ldHlpC(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.C)
	return true
}

func ldHlpD(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.D)
	return true
}

func ldHlpE(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.E)
	return true
}

func ldHlpH(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.H)
	return true
}

func ldHlpL(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.L)
	return true
}

func halt(cpu *cpu, value uint16) bool {
	cpu.halt = true
	return true
}

func ldHlpA(cpu *cpu, value uint16) bool {
	cpu.mmu.writeByte(cpu.registers.readHL(), cpu.registers.A)
	return true
}
