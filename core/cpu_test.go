package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0)
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}
}

func TestLdBcNn(t *testing.T) {
	system := New()
	system.cpu.registers.writeBC(0x0)
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x1)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, 0xFF0F)

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x1]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0xFF {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0xFF)
		t.Fail()
	}

	if system.cpu.registers.C != 0x0F {
		t.Logf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x0F)
		t.Fail()
	}
}

func TestLdBcpA(t *testing.T) {
	system := New()
	system.cpu.registers.A = 0xF0
	system.cpu.registers.writeBC(workRAMBank0BaseAddress)
	system.cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x2)
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x2]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	value := system.cpu.mmu.readByte(workRAMBank0BaseAddress)
	if value != 0xF0 {
		t.Logf("system.cpu.mmu.memory[0xC000] = %X, expected = %X", value, 0xF0)
		t.Fail()
	}
}

func TestIncBc(t *testing.T) {
	system := New()
	system.cpu.registers.writeBC(0x0101)
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x3)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x3]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x01 {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x01)
		t.Fail()
	}

	if system.cpu.registers.C != 0x02 {
		t.Logf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x02)
		t.Fail()
	}
}

func TestIncB(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x01
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x4)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x4]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x02 {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x02)
		t.Fail()
	}

	//TODO test affected flags
}

func TestDecB(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x01
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x5)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x5]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x00 {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x00)
		t.Fail()
	}

	//TODO test affected flags
}

func TestLdBn(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x00
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x6)
	system.cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x6]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x06 {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x06)
		t.Fail()
	}
}

func TestRlca(t *testing.T) {
	system := New()
	system.cpu.registers.F |= negativeFlag | zeroFlag | halfCarryFlag
	system.cpu.registers.A = 0x81
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x7)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x7]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.A != 0x02 {
		t.Logf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.A, 0x02)
		t.Fail()
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Logf("system.cpu.registers.F = %X, expected = %X", system.cpu.registers.F, carryFlag)
		t.Fail()
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Logf("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Logf("Zero flag must be reseted")
	}

	if (system.cpu.registers.F & halfCarryFlag) == halfCarryFlag {
		t.Logf("Half carry flag must be reseted")
	}
}

func TestLdNnpSp(t *testing.T) {
	system := New()
	system.cpu.registers.sp = 0xF0FF
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x8)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, workRAMBank0BaseAddress)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x8]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	value := system.cpu.mmu.readWord(workRAMBank0BaseAddress)
	if value != 0xF0FF {
		t.Logf("(%X) address value = %X, expected = %X", workRAMBank0BaseAddress, value, 0xF0FF)
		t.Fail()
	}
}

func TestAddHlBc(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x9)

	system.cpu.registers.F = negativeFlag | zeroFlag

	//   0100 0100 0000 0000
	// + 1100 1100 0000 0000
	system.cpu.registers.writeBC(0x4400)
	system.cpu.registers.writeHL(0xCC00)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x9]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	value := system.cpu.registers.readHL()
	if value != 4096 {
		t.Logf("HL register value = %d, expected = %d", value, 4096)
		t.Fail()
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Logf("halfCarryFlag value = %X, expected = %X", system.cpu.registers.F&halfCarryFlag, halfCarryFlag)
		t.Fail()
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Logf("carryFlag value = %X, expected = %X", system.cpu.registers.F&carryFlag, carryFlag)
		t.Fail()
	}

	// Must be reseted
	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Logf("negativeFlag value = %X, expected = %X", system.cpu.registers.F&negativeFlag, 0)
		t.Fail()
	}

	// Must stay untouched
	if (system.cpu.registers.F & zeroFlag) != zeroFlag {
		t.Logf("zeroFlag value = %X, expected = %X", system.cpu.registers.F&zeroFlag, 0)
		t.Fail()
	}
}

func TestLdABcp(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xA)
	system.cpu.registers.writeBC(workRAMBank0BaseAddress)
	system.cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)
	system.cpu.registers.A = 0x01
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0A]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.A != 0x0F {
		t.Logf("system.cpu.registers.A = %X, expected = %X", system.cpu.registers.A, 0x0F)
	}
}
