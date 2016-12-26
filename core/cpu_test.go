package core

import "testing"

type instructionTestFunction func(t *testing.T, cpu *cpu) func()

func TestMiscControlInstructions(t *testing.T) {
	t.Run("NOP", noFlagModificationInstructionTestHandler(testNop, 0x00))
	t.Run("STOP", noFlagModificationInstructionTestHandler(testStop, 0x10))
	t.Run("STOP folowed by another instruction", noFlagModificationInstructionTestHandler(testStopIncBc, 0x10))
}

func TestLoadMoveStoreInstructions(t *testing.T) {
	t.Run("LD BC nn", noFlagModificationInstructionTestHandler(testLdBcNn, 0x01))
	t.Run("LD (BC) A", noFlagModificationInstructionTestHandler(testLdBcpA, 0x02))
	t.Run("LD B n", noFlagModificationInstructionTestHandler(testLdBn, 0x06))
	t.Run("LD (nn) SP", noFlagModificationInstructionTestHandler(testLdNnpSp, 0x08))
	t.Run("LD A (BC)", noFlagModificationInstructionTestHandler(testLdABcp, 0x0A))
	t.Run("LD C n", noFlagModificationInstructionTestHandler(testLdCN, 0x0E))
	t.Run("LD DE nn", noFlagModificationInstructionTestHandler(testLdDeNn, 0x11))
	t.Run("LD (DE) A", noFlagModificationInstructionTestHandler(testLdDepA, 0x12))
	t.Run("LD D n", noFlagModificationInstructionTestHandler(testLdDn, 0x16))
	t.Run("LD A (DE)", noFlagModificationInstructionTestHandler(testLdADep, 0x1A))
	t.Run("LD E n", noFlagModificationInstructionTestHandler(testLdEn, 0x1E))
	t.Run("LD HL nn", noFlagModificationInstructionTestHandler(testLdHlNn, 0x21))
	t.Run("LD (HL+) A", noFlagModificationInstructionTestHandler(testLdHlpAIncHl, 0x22))
	t.Run("LD H n", noFlagModificationInstructionTestHandler(testLdHn, 0x26))
	t.Run("LD A (HL+)", noFlagModificationInstructionTestHandler(testLdAHlpIncHl, 0x2A))
	t.Run("LD L n", noFlagModificationInstructionTestHandler(testLdLn, 0x2E))
	t.Run("LD SP nn", noFlagModificationInstructionTestHandler(testLdSpNn, 0x31))
	t.Run("LD (HL-) A", noFlagModificationInstructionTestHandler(testLdHlpADecHl, 0x32))
	t.Run("LD (HL) n", noFlagModificationInstructionTestHandler(testLdHlpn, 0x36))
	t.Run("LD A (HL+)", noFlagModificationInstructionTestHandler(testLdAHlpDecHl, 0x3A))
	t.Run("LD A n", noFlagModificationInstructionTestHandler(testLdAn, 0x3E))
	t.Run("LD B B", noFlagModificationInstructionTestHandler(testLdBb, 0x40))
	t.Run("LD B C", noFlagModificationInstructionTestHandler(testLdBc, 0x41))
	t.Run("LD B D", noFlagModificationInstructionTestHandler(testLdBd, 0x42))
	t.Run("LD B E", noFlagModificationInstructionTestHandler(testLdBe, 0x43))
	t.Run("LD B H", noFlagModificationInstructionTestHandler(testLdBh, 0x44))
	t.Run("LD B L", noFlagModificationInstructionTestHandler(testLdBl, 0x45))
	t.Run("LD B (HL)", noFlagModificationInstructionTestHandler(testLdBhlp, 0x46))
	t.Run("LD B A", noFlagModificationInstructionTestHandler(testLdBa, 0x47))
	t.Run("LD C B", noFlagModificationInstructionTestHandler(testLdCb, 0x48))
	t.Run("LD C C", noFlagModificationInstructionTestHandler(testLdCc, 0x49))
	t.Run("LD C D", noFlagModificationInstructionTestHandler(testLdCd, 0x4A))
	t.Run("LD C E", noFlagModificationInstructionTestHandler(testLdCe, 0x4B))
	t.Run("LD C H", noFlagModificationInstructionTestHandler(testLdCh, 0x4C))
	t.Run("LD C L", noFlagModificationInstructionTestHandler(testLdCl, 0x4D))
	t.Run("LD C (HL)", noFlagModificationInstructionTestHandler(testLdChlp, 0x4E))
	t.Run("LD C A", noFlagModificationInstructionTestHandler(testLdCa, 0x4F))
	t.Run("LD D B", noFlagModificationInstructionTestHandler(testLdDb, 0x50))
	t.Run("LD D C", noFlagModificationInstructionTestHandler(testLdDc, 0x51))
	t.Run("LD D D", noFlagModificationInstructionTestHandler(testLdDd, 0x52))
	t.Run("LD D E", noFlagModificationInstructionTestHandler(testLdDe, 0x53))
	t.Run("LD D H", noFlagModificationInstructionTestHandler(testLdDh, 0x54))
	t.Run("LD D L", noFlagModificationInstructionTestHandler(testLdDl, 0x55))
	t.Run("LD D (HL)", noFlagModificationInstructionTestHandler(testLdDhlp, 0x56))
	t.Run("LD D A", noFlagModificationInstructionTestHandler(testLdDa, 0x57))
	t.Run("LD E B", noFlagModificationInstructionTestHandler(testLdEb, 0x58))
	t.Run("LD E C", noFlagModificationInstructionTestHandler(testLdEc, 0x59))
	t.Run("LD E D", noFlagModificationInstructionTestHandler(testLdEd, 0x5A))
	t.Run("LD E E", noFlagModificationInstructionTestHandler(testLdEe, 0x5B))
	t.Run("LD E H", noFlagModificationInstructionTestHandler(testLdEh, 0x5C))
	t.Run("LD E L", noFlagModificationInstructionTestHandler(testLdEl, 0x5D))
	t.Run("LD E (HL)", noFlagModificationInstructionTestHandler(testLdEhlp, 0x5E))
	t.Run("LD E A", noFlagModificationInstructionTestHandler(testLdEa, 0x5F))
}

func Test16BitsArithmeticLogicalInstructions(t *testing.T) {
	t.Run("INC BC", noFlagModificationInstructionTestHandler(testIncBc, 0x03))
	t.Run("INC BC overflow", noFlagModificationInstructionTestHandler(testIncBcOverflow, 0x03))

	t.Run("ADD HL, BC", instructionTestHandler(testAddHlBc, 0x09))
	t.Run("ADD HL, BC carry and half carry flags trigger", instructionTestHandler(testAddHlBcCarryAndHalfCarryEnable, 0x09))

	t.Run("DEC BC", noFlagModificationInstructionTestHandler(testDecBc, 0x0B))
	t.Run("DEC BC underflow", noFlagModificationInstructionTestHandler(testDecBcUnderflow, 0x0B))

	t.Run("INC DE", noFlagModificationInstructionTestHandler(testIncDe, 0x13))
	t.Run("INC DE overflow", noFlagModificationInstructionTestHandler(testIncDeOverflow, 0x13))

	t.Run("ADD HL DE", instructionTestHandler(testAddHlDe, 0x19))
	t.Run("ADD HL DE carry and half carry flags trigger", instructionTestHandler(testAddHlDeCarryAndHalfCarryEnable, 0x19))

	t.Run("DEC DE", noFlagModificationInstructionTestHandler(testDecDe, 0x1B))
	t.Run("DEC DE underflow", noFlagModificationInstructionTestHandler(testDecDeUnderflow, 0x1B))

	t.Run("INC HL", noFlagModificationInstructionTestHandler(testIncHl, 0x23))
	t.Run("INC HL overflow", noFlagModificationInstructionTestHandler(testIncHlOverflow, 0x23))

	t.Run("ADD HL HL", instructionTestHandler(testAddHlHl, 0x29))
	t.Run("ADD HL HL carry and half carry flags trigger", instructionTestHandler(testAddHlHlCarryAndHalfCarryEnable, 0x29))

	t.Run("DEC HL", noFlagModificationInstructionTestHandler(testDecHl, 0x2B))
	t.Run("DEC HL underflow", noFlagModificationInstructionTestHandler(testDecHlUnderflow, 0x2B))

	t.Run("INC SP", noFlagModificationInstructionTestHandler(testIncSp, 0x33))
	t.Run("INC SP overflow", noFlagModificationInstructionTestHandler(testIncSpOverflow, 0x33))

	t.Run("ADD HL SP", instructionTestHandler(testAddHlSp, 0x39))
	t.Run("ADD HL SP carry and half carry flags trigger", instructionTestHandler(testAddHlSpCarryAndHalfCarryEnable, 0x39))

	t.Run("DEC SP", noFlagModificationInstructionTestHandler(testDecSp, 0x3B))
	t.Run("DEC SP underflow", noFlagModificationInstructionTestHandler(testDecSpUnderflow, 0x3B))
}

func Test8BitsArithmeticLogicalInstructions(t *testing.T) {
	t.Run("INC B", instructionTestHandler(testIncB, 0x04))
	t.Run("INC B register overflow and half carry flag trigger", instructionTestHandler(testIncBOverflowAndHalfCarry, 0x04))
	t.Run("DEC B", instructionTestHandler(testDecB, 0x05))
	t.Run("DEC B zero flag", instructionTestHandler(testDecBZeroFlag, 0x05))
	t.Run("DEC B underflow", instructionTestHandler(testDecBUnderflow, 0x05))

	t.Run("INC C", instructionTestHandler(testIncC, 0x0C))
	t.Run("INC C register overflow and half carry flag trigger", instructionTestHandler(testIncCOverflowAndHalfCarry, 0x0C))
	t.Run("DEC C", instructionTestHandler(testDecC, 0x0D))
	t.Run("DEC C zero flag", instructionTestHandler(testDecCZeroFlag, 0x0D))
	t.Run("DEC C underflow", instructionTestHandler(testDecCUnderflow, 0x0D))

	t.Run("INC D", instructionTestHandler(testIncD, 0x14))
	t.Run("INC D register overflow and half carry flag trigger", instructionTestHandler(testIncDOverflowAndHalfCarry, 0x14))
	t.Run("DEC D", instructionTestHandler(testDecD, 0x15))
	t.Run("DEC D zero flag", instructionTestHandler(testDecDZeroFlag, 0x15))
	t.Run("DEC D underflow", instructionTestHandler(testDecDUnderflow, 0x15))

	t.Run("INC E", instructionTestHandler(testIncE, 0x1C))
	t.Run("INC E register overflow and half carry flag trigger", instructionTestHandler(testIncEOverflowAndHalfCarry, 0x1C))
	t.Run("DEC E", instructionTestHandler(testDecE, 0x1D))
	t.Run("DEC E zero flag", instructionTestHandler(testDecEZeroFlag, 0x1D))
	t.Run("DEC E underflow", instructionTestHandler(testDecEUnderflow, 0x1D))

	t.Run("INC H", instructionTestHandler(testIncH, 0x24))
	t.Run("INC H register overflow and half carry flag trigger", instructionTestHandler(testIncHOverflowAndHalfCarry, 0x24))
	t.Run("DEC H", instructionTestHandler(testDecH, 0x25))
	t.Run("DEC H zero flag", instructionTestHandler(testDecHZeroFlag, 0x25))
	t.Run("DEC H underflow", instructionTestHandler(testDecHUnderflow, 0x25))

	t.Run("DAA", instructionTestHandler(testDaa, 0x27))
	t.Run("DAA overflow to zero", instructionTestHandler(testDaaOverflowToZero, 0x27))
	t.Run("DAA half carry flag enabled", instructionTestHandler(testDaaHalfCarryFlagEnabled, 0x27))
	t.Run("DAA carry flag enabled", instructionTestHandler(testDaaCarryFlagEnabled, 0x27))
	t.Run("DAA half carry and carry flag enabled", instructionTestHandler(testDaaHalfCarryAndCarryFlagEnabled, 0x27))
	t.Run("DAA undeflow negative flag and half carry flag enabled", instructionTestHandler(testDaaNegativeFlagAndHalfCarryFlagEnabled, 0x27))
	t.Run("DAA undeflow negative flag and carry flag enabled", instructionTestHandler(testDaaNegativeFlagAndCarryFlagEnabled, 0x27))
	t.Run("DAA undeflow negative flag, carry flag and half carry flag enabled", instructionTestHandler(testDaaNegativeFlagCarryFlagAndHalfCarryFlagEnabled, 0x27))

	t.Run("INC L", instructionTestHandler(testIncL, 0x2C))
	t.Run("INC L register overflow and half carry flag trigger", instructionTestHandler(testIncLOverflowAndHalfCarry, 0x2C))
	t.Run("DEC L", instructionTestHandler(testDecL, 0x2D))
	t.Run("DEC L zero flag", instructionTestHandler(testDecLZeroFlag, 0x2D))
	t.Run("DEC L underflow", instructionTestHandler(testDecLUnderflow, 0x2D))

	t.Run("CPL", instructionTestHandler(testCpl, 0x2F))

	t.Run("INC (HL)", instructionTestHandler(testIncHlp, 0x34))
	t.Run("INC (HL) register overflow and half carry flag trigger", instructionTestHandler(testIncHlpOverflowAndHalfCarry, 0x34))
	t.Run("DEC (HL)", instructionTestHandler(testDecHlp, 0x35))
	t.Run("DEC (HL) zero flag", instructionTestHandler(testDecHlpZeroFlag, 0x35))
	t.Run("DEC (HL) underflow", instructionTestHandler(testDecHlpUnderflow, 0x35))

	t.Run("SCF", instructionTestHandler(testScf, 0x37))

	t.Run("INC A", instructionTestHandler(testIncA, 0x3C))
	t.Run("INC A register overflow and half carry flag trigger", instructionTestHandler(testIncAOverflowAndHalfCarry, 0x3C))
	t.Run("DEC A", instructionTestHandler(testDecA, 0x3D))
	t.Run("DEC A zero flag", instructionTestHandler(testDecAZeroFlag, 0x3D))
	t.Run("DEC A underflow", instructionTestHandler(testDecAUnderflow, 0x3D))

	t.Run("CCF carry flag disabled", instructionTestHandler(testCcfDisabledCarryFlag, 0x3F))
	t.Run("CCF carry flag enabled", instructionTestHandler(testCcfEnabledCarryFlag, 0x3F))
}

func Test8bitRotationsshiftsAndBitInstructions(t *testing.T) {
	t.Run("RCLA", instructionTestHandler(testRlca, 0x07))
	t.Run("RCLA apply carry and enable carry flag", instructionTestHandler(testRlcaApplyCarryAndEnableCarryFlag, 0x07))

	t.Run("RRCA", instructionTestHandler(testRrca, 0x0F))
	t.Run("RRCA disable carry flag", instructionTestHandler(testRrcaDisableCarryFlag, 0x0F))

	t.Run("RLA", instructionTestHandler(testRla, 0x17))
	t.Run("RLA apply carry and disable carry flag", instructionTestHandler(testRlaApplyCarryAndDisableCarryFlag, 0x17))
	t.Run("RLA apply carry and enable carry flag", instructionTestHandler(testRlaApplyCarryAndEnableCarryFlag, 0x17))

	t.Run("RRA", instructionTestHandler(testRra, 0x1F))
	t.Run("RRA apply carry and disable carry flag", instructionTestHandler(testRraApplyCarryAndDisableCarryFlag, 0x1F))
	t.Run("RRA apply carry and enable carry flag", instructionTestHandler(testRraApplyCarryAndEnableCarryFlag, 0x1F))
}

func TestJumpCalls(t *testing.T) {
	t.Run("JR n", noFlagModificationInstructionTestHandler(testJr, 0x18))
	t.Run("JR n negative value", noFlagModificationInstructionTestHandler(testJrNegativeValue, 0x18))

	t.Run("JR NZ n", noFlagModificationInstructionTestHandler(testJrNzn, 0x20))
	t.Run("JR NZ n zero flag enabled", noFlagModificationInstructionTestHandler(testJrNznZeroFlagEnabled, 0x20))
	t.Run("JR NZ n negative value", noFlagModificationInstructionTestHandler(testJrNznNegativeValue, 0x20))

	t.Run("JR Z n", noFlagModificationInstructionTestHandler(testJrZn, 0x28))
	t.Run("JR Z n zero flag disabled", noFlagModificationInstructionTestHandler(testJrZnZeroFlagDisabled, 0x28))
	t.Run("JR Z n negative value", noFlagModificationInstructionTestHandler(testJrZnNegativeValue, 0x28))

	t.Run("JR NC n", noFlagModificationInstructionTestHandler(testJrNcn, 0x30))
	t.Run("JR NC n zero flag enabled", noFlagModificationInstructionTestHandler(testJrNcnCarryFlagEnabled, 0x30))
	t.Run("JR NC n negative value", noFlagModificationInstructionTestHandler(testJrNcnNegativeValue, 0x30))

	t.Run("JR C n", noFlagModificationInstructionTestHandler(testJrCn, 0x38))
	t.Run("JR C n carry flag disabled", noFlagModificationInstructionTestHandler(testJrCnCarryFlagDisabled, 0x38))
	t.Run("JR C n negative value", noFlagModificationInstructionTestHandler(testJrCnNegativeValue, 0x38))
}

func instructionTestHandler(test instructionTestFunction, opcode byte) func(t *testing.T) {
	return func(t *testing.T) {
		system := New()
		system.cpu.registers.pc = romBank00BaseAddress
		system.cpu.mmu.writeByte(romBank00BaseAddress, opcode)

		postExecuteCheck := test(t, system.cpu)
		system.Execute()
		postExecuteCheck()

		instruction := (*system.cpu.instructionSet)[int(opcode)]
		if system.cpu.registers.pc != uint16(instruction.length) {
			t.Errorf("system.cpu.registers.pc = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		}
	}
}

func noFlagModificationInstructionTestHandler(test instructionTestFunction, opcode byte) func(t *testing.T) {
	return func(t *testing.T) {
		system := New()
		system.cpu.registers.pc = romBank00BaseAddress
		system.cpu.mmu.writeByte(romBank00BaseAddress, opcode)
		system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

		postExecuteCheck := test(t, system.cpu)
		system.Execute()
		postExecuteCheck()

		instruction := (*system.cpu.instructionSet)[int(opcode)]
		if system.cpu.registers.pc != uint16(instruction.length) {
			t.Errorf("system.cpu.registers.pc = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		}

		if system.cpu.registers.F&zeroFlag == 0x0 {
			t.Error("Zero flag must stay untouched")
		}
		if system.cpu.registers.F&negativeFlag == 0x0 {
			t.Error("Negative flag must stay untouched")
		}
		if system.cpu.registers.F&halfCarryFlag == 0x0 {
			t.Error("Half carry flag must stay untouched")
		}
		if system.cpu.registers.F&carryFlag == 0x0 {
			t.Error("Carry flag must stay untouched")
		}
	}
}

func testNop(t *testing.T, cpu *cpu) func() {
	return func() {}
}

func testStop(t *testing.T, cpu *cpu) func() {
	return func() {
		if !cpu.stoped {
			t.Error("Cpu must be stoped")
		}
	}
}

func testStopIncBc(t *testing.T, cpu *cpu) func() {
	return func() {
		if !cpu.stoped {
			t.Error("Cpu must be stoped")
		}

		var previousProgramCounter = cpu.registers.pc
		cpu.registers.writeBC(0xBC)
		cpu.mmu.writeByte(cpu.registers.pc, 0x03)
		cpu.next()

		var value = cpu.registers.readBC()
		if value != 0xBC {
			t.Errorf("system.cpu.registers.BC = %0#2X, Expected value %0#2X", value, 0xBC)
		}

		if cpu.registers.pc != previousProgramCounter {
			t.Errorf("Cpu is stoped program counter must not be incremented, program counter value = %d", cpu.registers.pc)
		}
	}
}

func testLdBcNn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeWord(romBank00BaseAddress+1, 0xFF0F)
	cpu.registers.writeBC(0x0)

	return func() {
		if cpu.registers.B != 0xFF {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0xFF)
		}

		if cpu.registers.C != 0x0F {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x0F)
		}
	}
}

func testLdBcpA(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0xF0
	cpu.registers.writeBC(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xFF)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0xF0 {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xF0)
		}
	}
}

func testLdBn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.B = 0x00

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdNnpSp(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeWord(romBank00BaseAddress+1, workRAMBank0BaseAddress)
	cpu.registers.sp = 0xF0FF

	return func() {
		value := cpu.mmu.readWord(workRAMBank0BaseAddress)
		if value != 0xF0FF {
			t.Errorf("(%0#4X) address value =%0#4X, expected = %0#4X", workRAMBank0BaseAddress, value, 0xF0FF)
		}
	}
}

func testLdABcp(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeBC(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x0F {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0F)
		}
	}
}

func testLdCN(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xE)
	cpu.registers.C = 0x1

	return func() {
		if cpu.registers.C != 0xE {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0xE)
		}
	}
}

func testLdDeNn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeWord(romBank00BaseAddress+1, 0xF0FF)
	cpu.registers.writeDE(0x0102)

	return func() {
		if cpu.registers.D != 0xF0 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0xF0)
		}

		if cpu.registers.E != 0xFF {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0xFF)
		}
	}
}

func testLdDepA(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0xFF
	cpu.registers.writeDE(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0xFF {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xFF)
		}
	}
}

func testLdDn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.D = 0x00

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdADep(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeDE(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x0F {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0F)
		}
	}
}

func testLdEn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.E = 0x00

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdHlNn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeWord(romBank00BaseAddress+1, 0xF0FF)
	cpu.registers.writeHL(0x0102)

	return func() {
		if cpu.registers.H != 0xF0 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0xF0)
		}

		if cpu.registers.L != 0xFF {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0xFF)
		}
	}
}

func testLdHlpAIncHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0xFF
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0xFF {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xFF)
		}

		hl := cpu.registers.readHL()
		if hl != workRAMBank0BaseAddress+0x0001 {
			t.Errorf("cpu.register.HL = %0#4X, expected = %0#4X", hl, workRAMBank0BaseAddress+0x0001)
		}
	}
}

func testLdHn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.H = 0x00

	return func() {
		if cpu.registers.H != 0x06 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x06)
		}
	}
}

func testLdAHlpIncHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x0F {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0F)
		}
		hl := cpu.registers.readHL()
		if hl != workRAMBank0BaseAddress+0x0001 {
			t.Errorf("cpu.registers.HL = %0#4X, expected = %0#4X", hl, workRAMBank0BaseAddress+0x0001)
		}
	}
}

func testLdLn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.L = 0x00

	return func() {
		if cpu.registers.L != 0x06 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x06)
		}
	}
}

func testLdSpNn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeWord(romBank00BaseAddress+1, 0xF0FF)
	cpu.registers.sp = 0x0102

	return func() {
		if cpu.registers.sp != 0xF0FF {
			t.Errorf("cpu.registers.sp = %0#4X, expected = %0#4X", cpu.registers.sp, 0xF0FF)
		}
	}
}

func testLdHlpADecHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0xFF
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0xFF {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xFF)
		}

		hl := cpu.registers.readHL()
		if hl != workRAMBank0BaseAddress-0x0001 {
			t.Errorf("cpu.register.HL = %0#4X, expected = %0#4X", hl, workRAMBank0BaseAddress-0x0001)
		}
	}
}

func testLdHlpn(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF0)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0xF0 {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xF0)
		}
	}
}

func testLdAHlpDecHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress + 0x0001)
	cpu.mmu.writeByte(workRAMBank0BaseAddress+0x0001, 0x0F)
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x0F {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0F)
		}
		hl := cpu.registers.readHL()
		if hl != workRAMBank0BaseAddress {
			t.Errorf("cpu.registers.HL = %0#4X, expected = %0#4X", hl, workRAMBank0BaseAddress)
		}
	}
}

func testLdAn(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	cpu.registers.A = 0x00

	return func() {
		if cpu.registers.A != 0x06 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x06)
		}
	}
}

func testLdBb(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBc(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.C = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBd(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.D = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBe(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.E = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBh(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.H = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBl(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.L = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBhlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x06)
	cpu.registers.B = 0x00

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdBa(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x00
	cpu.registers.A = 0x06

	return func() {
		if cpu.registers.B != 0x06 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x06)
		}
	}
}

func testLdCb(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.B = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCc(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCd(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.D = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCe(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.E = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCh(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.H = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCl(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.L = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdChlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x06)
	cpu.registers.C = 0x00

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdCa(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x00
	cpu.registers.A = 0x06

	return func() {
		if cpu.registers.C != 0x06 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x06)
		}
	}
}

func testLdDb(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.B = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDc(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.C = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDd(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDe(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.E = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDh(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.H = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDl(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.L = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDhlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x06)
	cpu.registers.D = 0x00

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdDa(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x00
	cpu.registers.A = 0x06

	return func() {
		if cpu.registers.D != 0x06 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x06)
		}
	}
}

func testLdEb(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.B = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEc(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.C = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEd(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.D = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEe(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEh(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.H = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEl(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.L = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEhlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x06)
	cpu.registers.E = 0x00

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testLdEa(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x00
	cpu.registers.A = 0x06

	return func() {
		if cpu.registers.E != 0x06 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x06)
		}
	}
}

func testIncBc(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeBC(0x0101)

	return func() {
		if cpu.registers.B != 0x01 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x01)
		}

		if cpu.registers.C != 0x02 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x02)
		}
	}
}

func testIncBcOverflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeBC(0xFFFF)

	return func() {
		if cpu.registers.B != 0x0 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x00)
		}

		if cpu.registers.C != 0x0 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x00)
		}
	}
}

func testAddHlBc(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	cpu.registers.writeBC(0x0003)
	cpu.registers.writeHL(0x0200)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x0203 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x0203)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testAddHlBcCarryAndHalfCarryEnable(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	//   0100 0100 0000 0000
	// + 1100 1100 0000 0000
	cpu.registers.writeBC(0x4400)
	cpu.registers.writeHL(0xCC00)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x1000 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x1000)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecBc(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeBC(0x1)

	return func() {
		value := cpu.registers.readBC()
		if value != 0x0 {
			t.Errorf("cpu.registers.BC = %0#4X, expected = %0#4X", value, 0x0)
		}
	}
}

func testDecBcUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeBC(0x0)

	return func() {
		value := cpu.registers.readBC()
		if value != 0xFFFF {
			t.Errorf("cpu.registers.BC = %0#4X, expected = %0#4X", value, 0xFFFF)
		}
	}
}

func testIncDe(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeDE(0x0101)

	return func() {
		if cpu.registers.D != 0x01 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x01)
		}

		if cpu.registers.E != 0x02 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x02)
		}
	}
}

func testIncDeOverflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeDE(0xFFFF)

	return func() {
		if cpu.registers.D != 0x00 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x00)
		}

		if cpu.registers.E != 0x00 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x00)
		}
	}
}

func testAddHlDe(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	cpu.registers.writeHL(0x0200)
	cpu.registers.writeDE(0x0003)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x0203 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x0203)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testAddHlDeCarryAndHalfCarryEnable(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	//   0100 0100 0000 0000
	// + 1100 1100 0000 0000
	cpu.registers.writeDE(0x4400)
	cpu.registers.writeHL(0xCC00)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x1000 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x1000)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecDe(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeDE(0x1)

	return func() {
		value := cpu.registers.readDE()
		if value != 0x0 {
			t.Errorf("cpu.registers.DE = %0#4X, expected = %0#4X", value, 0x0)
		}
	}
}

func testDecDeUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeDE(0x0)

	return func() {
		value := cpu.registers.readDE()
		if value != 0xFFFF {
			t.Errorf("cpu.registers.DE = %0#4X, expected = %0#4X", value, 0xFFFF)
		}
	}
}

func testIncHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(0x0101)

	return func() {
		if cpu.registers.H != 0x01 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x01)
		}

		if cpu.registers.L != 0x02 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x02)
		}
	}
}

func testIncHlOverflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(0xFFFF)

	return func() {
		if cpu.registers.H != 0x00 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x00)
		}

		if cpu.registers.L != 0x00 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x00)
		}
	}
}

func testAddHlHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	cpu.registers.writeHL(0x0200)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x0400 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x0400)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testAddHlHlCarryAndHalfCarryEnable(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	//   1100 1100 0000 0000
	// + 1100 1100 0000 0000
	cpu.registers.writeHL(0xCC00)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x9800 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x9800)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHl(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(0x1)

	return func() {
		value := cpu.registers.readHL()
		if value != 0x0 {
			t.Errorf("cpu.registers.HL = %0#4X, expected = %0#4X", value, 0x0)
		}
	}
}

func testDecHlUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.writeHL(0x0)

	return func() {
		value := cpu.registers.readHL()
		if value != 0xFFFF {
			t.Errorf("cpu.registers.HL = %0#4X, expected = %0#4X", value, 0xFFFF)
		}
	}
}

func testIncSp(t *testing.T, cpu *cpu) func() {
	cpu.registers.sp = 0x0101

	return func() {
		if cpu.registers.sp != 0x0102 {
			t.Errorf("cpu.registers.sp = %0#4X, expected = %0#4X", cpu.registers.sp, 0x0102)
		}
	}
}

func testIncSpOverflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.sp = 0xFFFF

	return func() {
		if cpu.registers.sp != 0x00 {
			t.Errorf("cpu.registers.sp = %0#4X, expected = %0#4X", cpu.registers.sp, 0x0000)
		}
	}
}

func testAddHlSp(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	cpu.registers.writeHL(0x0200)
	cpu.registers.sp = 0x0200

	return func() {
		value := cpu.registers.readHL()
		if value != 0x0400 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x0400)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testAddHlSpCarryAndHalfCarryEnable(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag

	//   1100 1100 0000 0000
	// + 1100 1100 0000 0000
	cpu.registers.writeHL(0xCC00)
	cpu.registers.sp = 0xCC00

	return func() {
		value := cpu.registers.readHL()
		if value != 0x9800 {
			t.Errorf("HL register value = %0#4X, expected = %0#4X", value, 0x9800)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecSp(t *testing.T, cpu *cpu) func() {
	cpu.registers.sp = 0x0001

	return func() {
		if cpu.registers.sp != 0x0000 {
			t.Errorf("cpu.registers.sp = %0#4X, expected = %0#4X", cpu.registers.sp, 0x0000)
		}
	}
}

func testDecSpUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.sp = 0x0000

	return func() {
		if cpu.registers.sp != 0xFFFF {
			t.Errorf("cpu.registers.sp = %0#4X, expected = %0#4X", cpu.registers.sp, 0xFFFF)
		}
	}
}

func testIncB(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.B != 0x10 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncBOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.B = 0xFF

	return func() {
		if cpu.registers.B != 0x00 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecB(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0xF
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	return func() {
		if cpu.registers.B != 0x0E {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x0E)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecBZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x1
	cpu.registers.F = negativeFlag | halfCarryFlag | carryFlag

	return func() {
		if cpu.registers.B != 0x0 {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecBUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.B = 0x0
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	return func() {
		if cpu.registers.B != 0xFF {
			t.Errorf("cpu.registers.B = %0#2X, expected = %0#2X", cpu.registers.B, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncC(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0x1
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag
	return func() {
		if cpu.registers.C != 0x2 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x2)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncCOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.C = 0xFF
	cpu.registers.F = negativeFlag | carryFlag

	return func() {
		if cpu.registers.C != 0x0 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecC(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.C = 0x2

	return func() {
		if cpu.registers.C != 0x1 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecCZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.C = 0x01

	return func() {
		if cpu.registers.C != 0x0 {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecCUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.C = 0x0

	return func() {
		if cpu.registers.C != 0xFF {
			t.Errorf("cpu.registers.C = %0#2X, expected = %0#2X", cpu.registers.C, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testIncD(t *testing.T, cpu *cpu) func() {
	cpu.registers.D = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.D != 0x10 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncDOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.D = 0xFF

	return func() {
		if cpu.registers.D != 0x00 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecD(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.D = 0x2

	return func() {
		if cpu.registers.D != 0x1 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecDZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.D = 0x01

	return func() {
		if cpu.registers.D != 0x0 {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecDUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.D = 0x0

	return func() {
		if cpu.registers.D != 0xFF {
			t.Errorf("cpu.registers.D = %0#2X, expected = %0#2X", cpu.registers.D, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testIncE(t *testing.T, cpu *cpu) func() {
	cpu.registers.E = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.E != 0x10 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncEOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.E = 0xFF

	return func() {
		if cpu.registers.E != 0x00 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecE(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.E = 0x2

	return func() {
		if cpu.registers.E != 0x1 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecEZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.E = 0x01

	return func() {
		if cpu.registers.E != 0x0 {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecEUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.E = 0x0

	return func() {
		if cpu.registers.E != 0xFF {
			t.Errorf("cpu.registers.E = %0#2X, expected = %0#2X", cpu.registers.E, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testIncH(t *testing.T, cpu *cpu) func() {
	cpu.registers.H = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.H != 0x10 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncHOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.H = 0xFF

	return func() {
		if cpu.registers.H != 0x00 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecH(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.H = 0x2

	return func() {
		if cpu.registers.H != 0x1 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.H = 0x01

	return func() {
		if cpu.registers.H != 0x0 {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.H = 0x0

	return func() {
		if cpu.registers.H != 0xFF {
			t.Errorf("cpu.registers.H = %0#2X, expected = %0#2X", cpu.registers.H, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testDaa(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x0A
	cpu.registers.F &^= zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	return func() {
		if cpu.registers.A != 0x10 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x10)
		}
	}
}

func testDaaOverflowToZero(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0xFF - 0x65
	cpu.registers.F &^= zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	return func() {
		if cpu.registers.A != 0 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x00)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testDaaHalfCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = halfCarryFlag

	return func() {
		if cpu.registers.A != 0x06 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x06)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Errorf("Carry flag must be disabled")
		}
	}
}

func testDaaCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = carryFlag
	return func() {
		if cpu.registers.A != 0x60 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x60)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testDaaHalfCarryAndCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = halfCarryFlag | carryFlag
	return func() {
		if cpu.registers.A != 0x66 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x66)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testDaaNegativeFlagAndHalfCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = negativeFlag | halfCarryFlag
	return func() {
		if cpu.registers.A != 0xFA {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0xFA)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Errorf("Carry flag must be disabled")
		}
	}
}

func testDaaNegativeFlagAndCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = negativeFlag | carryFlag
	return func() {
		if cpu.registers.A != 0xA0 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0xA0)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testDaaNegativeFlagCarryFlagAndHalfCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x00
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	return func() {
		if cpu.registers.A != 0x9A {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x9A)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testIncL(t *testing.T, cpu *cpu) func() {
	cpu.registers.L = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.L != 0x10 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncLOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.L = 0xFF

	return func() {
		if cpu.registers.L != 0x00 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecL(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.L = 0x2

	return func() {
		if cpu.registers.L != 0x1 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecLZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.L = 0x01

	return func() {
		if cpu.registers.L != 0x0 {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecLUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.L = 0x0

	return func() {
		if cpu.registers.L != 0xFF {
			t.Errorf("cpu.registers.L = %0#2X, expected = %0#2X", cpu.registers.L, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testCpl(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = 0x0
	cpu.registers.A = 0xF9

	return func() {
		if cpu.registers.A != 0x06 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x06)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Errorf("Carry flag must be disabled")
		}
	}
}

func testIncHlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0x10 {
			t.Errorf("(cpu.registers.HL) = %0#2X, expected = %0#2X", value, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncHlpOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xFF)

	return func() {
		value := cpu.mmu.readByte(workRAMBank0BaseAddress)
		if value != 0x00 {
			t.Errorf("(cpu.registers.HL) = %0#2X, expected = %0#2X", value, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHlp(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x02)

	return func() {
		value := cpu.mmu.readByte(cpu.registers.readHL())
		if value != 0x01 {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0x01)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHlpZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x1)

	return func() {
		value := cpu.mmu.readByte(cpu.registers.readHL())
		if value != 0x00 {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0x00)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecHlpUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.writeHL(workRAMBank0BaseAddress)
	cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0)

	return func() {
		value := cpu.mmu.readByte(cpu.registers.readHL())
		if value != 0xFF {
			t.Errorf("cpu.mmu.memory[%0#4X] = %0#2X, expected = %0#2X", workRAMBank0BaseAddress, value, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testScf(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag

	return func() {
		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testIncA(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x0F
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	return func() {
		if cpu.registers.A != 0x10 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x010)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testIncAOverflowAndHalfCarry(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag
	cpu.registers.A = 0xFF

	return func() {
		if cpu.registers.A != 0x00 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecA(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.A = 0x2

	return func() {
		if cpu.registers.A != 0x1 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x1)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecAZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x0 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x0)
		}

		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) != halfCarryFlag {
			t.Error("Half carry flag must be enabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testDecAUnderflow(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	cpu.registers.A = 0x0

	return func() {
		if cpu.registers.A != 0xFF {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0xFF)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) != negativeFlag {
			t.Error("Negative flag must be enabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testCcfDisabledCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag

	return func() {
		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Errorf("Carry flag must be enabled")
		}
	}
}

func testCcfEnabledCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	return func() {
		if (cpu.registers.F & zeroFlag) != zeroFlag {
			t.Error("Zero flag must be enabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Errorf("Carry flag must be disabled")
		}
	}
}

func testRlca(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x02 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x02)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testRlcaApplyCarryAndEnableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag
	cpu.registers.A = 0x81

	return func() {
		if cpu.registers.A != 0x03 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x03)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testRrca(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x05
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag

	return func() {
		if cpu.registers.A != 0x82 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x82)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testRrcaDisableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.A = 0x02
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	return func() {
		if cpu.registers.A != 0x01 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x01)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must be disabled")
		}
	}
}

func testRla(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | zeroFlag | halfCarryFlag
	cpu.registers.A = 0x81

	return func() {
		if cpu.registers.A != 0x02 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x02)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testRlaApplyCarryAndDisableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	cpu.registers.A = 0x1

	return func() {
		if cpu.registers.A != 0x03 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x03)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must stay disabled")
		}
	}
}

func testRlaApplyCarryAndEnableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	cpu.registers.A = 0x81

	return func() {
		if cpu.registers.A != 0x03 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x03)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testRra(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = negativeFlag | zeroFlag | halfCarryFlag
	cpu.registers.A = 0x41

	return func() {
		if cpu.registers.A != 0x20 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x20)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testRraApplyCarryAndDisableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	cpu.registers.A = 0x02

	return func() {
		if cpu.registers.A != 0x81 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x81)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) == carryFlag {
			t.Error("Carry flag must stay disabled")
		}
	}
}

func testRraApplyCarryAndEnableCarryFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	cpu.registers.A = 0x01

	return func() {
		if cpu.registers.A != 0x80 {
			t.Errorf("cpu.registers.A = %0#2X, expected = %0#2X", cpu.registers.A, 0x80)
		}

		if (cpu.registers.F & zeroFlag) == zeroFlag {
			t.Error("Zero flag must be disabled")
		}

		if (cpu.registers.F & negativeFlag) == negativeFlag {
			t.Error("Negative flag must be disabled")
		}

		if (cpu.registers.F & halfCarryFlag) == halfCarryFlag {
			t.Error("Half carry flag must be disabled")
		}

		if (cpu.registers.F & carryFlag) != carryFlag {
			t.Error("Carry flag must be enabled")
		}
	}
}

func testJr(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x000A {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x000A)
		} else {
			cpu.registers.pc -= 0x8
		}
	}
}

func testJrNegativeValue(t *testing.T, cpu *cpu) func() {
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF1) // -15

	return func() {
		if cpu.registers.pc != 0xFFF3 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0xFFF3)
		} else {
			cpu.registers.pc += 15
		}
	}
}

func testJrNzn(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x000A {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x000A)
		} else {
			cpu.registers.F |= zeroFlag
			cpu.registers.pc -= 0x8
		}
		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 12)
		}
	}
}

func testJrNznZeroFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x0002 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x0002)
		}

		if cpu.ticks != 8 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrNznNegativeValue(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF1) // -15

	return func() {
		if cpu.registers.pc != 0xFFF3 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0xFFF3)
		} else {
			cpu.registers.pc += 15
			cpu.registers.F |= zeroFlag
		}

		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrZn(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x000A {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x000A)
		} else {
			cpu.registers.pc -= 0x8
		}
		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 12)
		}
	}
}

func testJrZnZeroFlagDisabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x0002 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x0002)
		} else {
			cpu.registers.F |= zeroFlag
		}

		if cpu.ticks != 8 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrZnNegativeValue(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= zeroFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF1) // -15

	return func() {
		if cpu.registers.pc != 0xFFF3 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0xFFF3)
		} else {
			cpu.registers.pc += 15
		}

		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrNcn(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x000A {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x000A)
		} else {
			cpu.registers.F |= carryFlag
			cpu.registers.pc -= 0x8
		}
		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 12)
		}
	}
}

func testJrNcnCarryFlagEnabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x0002 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x0002)
		}

		if cpu.ticks != 8 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrNcnNegativeValue(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF1) // -15

	return func() {
		if cpu.registers.pc != 0xFFF3 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0xFFF3)
		} else {
			cpu.registers.pc += 15
			cpu.registers.F |= carryFlag
		}

		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrCn(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x000A {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x000A)
		} else {
			cpu.registers.pc -= 0x8
		}
		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 12)
		}
	}
}

func testJrCnCarryFlagDisabled(t *testing.T, cpu *cpu) func() {
	cpu.registers.F &^= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0x8)

	return func() {
		if cpu.registers.pc != 0x0002 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0x0002)
		} else {
			cpu.registers.F |= carryFlag
		}

		if cpu.ticks != 8 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}

func testJrCnNegativeValue(t *testing.T, cpu *cpu) func() {
	cpu.registers.F |= carryFlag
	cpu.mmu.writeByte(romBank00BaseAddress+1, 0xF1) // -15

	return func() {
		if cpu.registers.pc != 0xFFF3 {
			t.Errorf("cpu.registers.pc = %0#4X, expected = %0#4X", cpu.registers.pc, 0xFFF3)
		} else {
			cpu.registers.pc += 15
		}

		if cpu.ticks != 12 {
			t.Errorf("cpu.ticks = %d, expected = %d", cpu.ticks, 8)
		}
	}
}
