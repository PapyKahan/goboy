package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0)
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}
}

func TestLdBcNn(t *testing.T) {
	system := New()
	system.cpu.registers.writeBC(0x0)
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x1)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, 0xFF0F)

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x1]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0xFF {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.B, 0xFF)
		t.Fail()
	}

	if system.cpu.registers.C != 0x0F {
		t.Logf("system.cpu.registers.C = %d, expected = %d", system.cpu.registers.C, 0x0F)
		t.Fail()
	}
}

func TestLdBcpA(t *testing.T) {
	system := New()
	system.cpu.registers.A = 0xF0
	system.cpu.registers.writeBC(workRAMBank0BaseAddress)
	system.cpu.mmu.writeByte(workRAMBank0BaseAddress, 0xF0)
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x2)
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x2]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	value := system.cpu.mmu.readByte(workRAMBank0BaseAddress)
	if value != 0xF0 {
		t.Logf("system.cpu.mmu.memory[0xC000] = %d, expected = %d", value, 0xF0)
		t.Fail()
	}
}

func TestIncBc(t *testing.T) {
	system := New()
	system.cpu.registers.writeBC(0x0101)
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x3)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x3]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x01 {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.B, 0x01)
		t.Fail()
	}

	if system.cpu.registers.C != 0x02 {
		t.Logf("system.cpu.registers.C = %d, expected = %d", system.cpu.registers.C, 0x02)
		t.Fail()
	}
}

func TestIncB(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x01
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x4)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x4]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x02 {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.B, 0x02)
		t.Fail()
	}
}

func TestDecB(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x01
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x5)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x5]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x00 {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.B, 0x00)
		t.Fail()
	}
}

func TestLdBn(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0x00
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x6)
	system.cpu.mmu.writeByte(romBank00BaseAddress+1, 0x6)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x6]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	if system.cpu.registers.B != 0x06 {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.B, 0x06)
		t.Fail()
	}
}

func TestRlca(t *testing.T) {
	system := New()
	system.cpu.registers.A = 0x00
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x7)
	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x7]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, instruction.length)
		t.Fail()
	}

	// TODO write more accurate tests.
	if system.cpu.registers.A != 0x00 {
		t.Logf("system.cpu.registers.B = %d, expected = %d", system.cpu.registers.A, 0x00)
		t.Fail()
	}
}
