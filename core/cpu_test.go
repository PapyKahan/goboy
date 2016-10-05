package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0)
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestLdBcNn(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x1)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, 0xFF0F)

	system.cpu.registers.writeBC(0x0)
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x1]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0xFF {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0xFF)
	}

	if system.cpu.registers.C != 0x0F {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x0F)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestLdBcpA(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x2)

	system.cpu.registers.A = 0xF0
	system.cpu.registers.writeBC(workRAMBank0BaseAddress)
	system.cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag

	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x2]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	value := system.cpu.mmu.readByte(workRAMBank0BaseAddress)
	if value != 0xF0 {
		t.Errorf("system.cpu.mmu.memory[0xC000] = %X, expected = %X", value, 0xF0)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func untouchedFlagsTest(f flags, t *testing.T) {
	if f&zeroFlag == 0x0 {
		t.Error("Zero flag must stay untouched")
	}
	if f&negativeFlag == 0x0 {
		t.Error("Negative flag must stay untouched")
	}
	if f&halfCarryFlag == 0x0 {
		t.Log("Half carry flag must stay untouched")
		t.Error()
	}
	if f&carryFlag == 0x0 {
		t.Error("Carry flag must stay untouched")
	}
}

func TestIncBc(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x3)
	system.cpu.registers.writeBC(0x0101)
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x3]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x01 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x01)
	}

	if system.cpu.registers.C != 0x02 {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x02)
	}
	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestIncB(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x4)
	system.cpu.registers.B = 0x0F
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x4]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x10 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x010)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Errorf("Negative flag must be stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Errorf("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Errorf("Zero flag must be reseted")
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Errorf("This operation must set half carry flag")
	}
}

func TestIncBOverflow(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x4)
	system.cpu.registers.B = 0xFF
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x4]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x00 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x0)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Errorf("Negative flag must be stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Errorf("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) != zeroFlag {
		t.Errorf("This operation must set zero flag")
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Errorf("This operation must set half carry flag")
	}
}

func TestDecB(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x5)
	system.cpu.registers.B = 0xF
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x5]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x0E {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x0E)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Errorf("Negative flag must be stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) != negativeFlag {
		t.Errorf("Negative flag must be set")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Errorf("Zero flag must be reseted.")
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Errorf("This operation must set half carry flag")
	}
}

func TestDecBZeroFlag(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x5)
	system.cpu.registers.B = 0x1
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x5]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x0 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x0)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Errorf("Negative flag must be stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) != negativeFlag {
		t.Errorf("Negative flag must be set")
	}

	if (system.cpu.registers.F & zeroFlag) != zeroFlag {
		t.Errorf("Zero flag must be set")
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Errorf("This operation must set half carry flag")
	}
}

func TestLdBn(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x6)
	system.cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	system.cpu.registers.B = 0x00
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x6]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.B != 0x06 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.B, 0x06)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
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
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.A != 0x02 {
		t.Errorf("system.cpu.registers.B = %X, expected = %X", system.cpu.registers.A, 0x02)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Errorf("Carry flag must stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Errorf("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Errorf("Zero flag must be reseted")
	}

	if (system.cpu.registers.F & halfCarryFlag) == halfCarryFlag {
		t.Errorf("Half carry flag must be reseted")
	}
}

func TestLdNnpSp(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x8)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, workRAMBank0BaseAddress)
	system.cpu.registers.sp = 0xF0FF
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag | halfCarryFlag
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x8]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	value := system.cpu.mmu.readWord(workRAMBank0BaseAddress)
	if value != 0xF0FF {
		t.Errorf("(%X) address value = %X, expected = %X", workRAMBank0BaseAddress, value, 0xF0FF)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
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
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	value := system.cpu.registers.readHL()
	if value != 4096 {
		t.Errorf("HL register value = %d, expected = %d", value, 4096)
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Error("Half carry flag must be set")
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Error("Carry flag must be set")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Logf("Negative flag must be reset")
		t.Fail()
	}

	// Must stay untouched
	if (system.cpu.registers.F & zeroFlag) != zeroFlag {
		t.Logf("Zero flag must stay untouched")
	}
}

func TestLdABcp(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xA)
	system.cpu.registers.writeBC(workRAMBank0BaseAddress)
	system.cpu.mmu.writeByte(workRAMBank0BaseAddress, 0x0F)
	system.cpu.registers.A = 0x01
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0A]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.A != 0x0F {
		t.Errorf("system.cpu.registers.A = %X, expected = %X", system.cpu.registers.A, 0x0F)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestDecBc(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xB)
	system.cpu.registers.writeBC(0x1)
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0B]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	value := system.cpu.registers.readBC()
	if value != 0x0 {
		t.Errorf("system.cpu.registers.BC = %X, expected = %X", value, 0x0)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestIncC(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xC)
	system.cpu.registers.C = 0x1
	system.cpu.registers.F = zeroFlag | negativeFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0C]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.C != 0x2 {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x2)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Error("Carry flag must stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Error("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Error("Zero flag must be unset")
	}

	if (system.cpu.registers.F & halfCarryFlag) == halfCarryFlag {
		t.Error("Half carry flag must be unset")
	}
}

func TestIncCOverflowAndHalfCarry(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xC)
	system.cpu.registers.C = 0xFF
	system.cpu.registers.F = negativeFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0C]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.C != 0x0 {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x0)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Error("Carry flag must stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Error("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) != zeroFlag {
		t.Error("Zero flag must be set")
	}

	if (system.cpu.registers.F & halfCarryFlag) != halfCarryFlag {
		t.Error("Half carry flag must be set")
	}
}

func TestDecC(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xD)
	system.cpu.registers.C = 0x1
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0D]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.C != 0x0 {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0x0)
	}
}

func TestLdCN(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0xE)
	system.cpu.mmu.writeByte(romBank00BaseAddress+1, 0xE)
	system.cpu.registers.C = 0x1
	system.cpu.registers.F = zeroFlag | negativeFlag | halfCarryFlag | carryFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0E]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.C != 0xE {
		t.Errorf("system.cpu.registers.C = %X, expected = %X", system.cpu.registers.C, 0xE)
	}

	untouchedFlagsTest(system.cpu.registers.F, t)
}

func TestRrca(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0F)
	system.cpu.registers.A = 0x05
	system.cpu.registers.F = carryFlag | zeroFlag | halfCarryFlag | negativeFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0F]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.A != 0x82 {
		t.Errorf("system.cpu.registers.A = %X, expected = %X", system.cpu.registers.A, 0x82)
	}

	if (system.cpu.registers.F & carryFlag) != carryFlag {
		t.Error("Carry flag must stay untouched")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Error("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Error("Zero flag must be reseted")
	}

	if (system.cpu.registers.F & halfCarryFlag) == halfCarryFlag {
		t.Error("Half carry flag must be reseted")
	}
}

func TestRrcaResetCarryFlag(t *testing.T) {
	system := New()
	system.cpu.registers.pc = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0F)
	system.cpu.registers.A = 0x02
	system.cpu.registers.F = carryFlag | zeroFlag | halfCarryFlag | negativeFlag
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0F]
	if system.cpu.registers.pc != uint16(instruction.length) {
		t.Errorf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.registers.pc, instruction.length)
	}

	if system.cpu.registers.A != 0x81 {
		t.Errorf("system.cpu.registers.A = %X, expected = %X", system.cpu.registers.A, 0x81)
	}

	if (system.cpu.registers.F & carryFlag) == carryFlag {
		t.Error("Carry flag must be reseted")
	}

	if (system.cpu.registers.F & negativeFlag) == negativeFlag {
		t.Error("Negative flag must be reseted")
	}

	if (system.cpu.registers.F & zeroFlag) == zeroFlag {
		t.Error("Zero flag must be reseted")
	}

	if (system.cpu.registers.F & halfCarryFlag) == halfCarryFlag {
		t.Error("Half carry flag must be reseted")
	}
}
