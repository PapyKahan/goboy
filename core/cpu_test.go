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
}

func Test8bitRotationsshiftsAndBitInstructions(t *testing.T) {
	t.Run("RCLA", instructionTestHandler(testRlca, 0x07))
	t.Run("RCLA apply carry and enable carry flag", instructionTestHandler(testRlcaApplyCarryAndEnableCarryFlag, 0x07))

	t.Run("RRCA", instructionTestHandler(testRrca, 0x0F))
	t.Run("RRCA disable carry flag", instructionTestHandler(testRrcaDisableCarryFlag, 0x0F))

	t.Run("RLA", instructionTestHandler(testRla, 0x17))
	t.Run("RLA apply carry and disable carry flag", instructionTestHandler(testRlaApplyCarryAndDisableCarryFlag, 0x17))
	t.Run("RLA apply carry and enable carry flag", instructionTestHandler(testRlaApplyCarryAndEnableCarryFlag, 0x17))
}

func TestJumpCalls(t *testing.T) {
	t.Run("JR n", noFlagModificationInstructionTestHandler(testJr, 0x18))
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
			t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
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
			t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
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
	cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

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
	cpu.registers.C = 0x1

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

func testDecCZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
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
	cpu.registers.D = 0x1

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

func testDecDZeroFlag(t *testing.T, cpu *cpu) func() {
	cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
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
