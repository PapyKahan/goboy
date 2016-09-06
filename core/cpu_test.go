package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x0)
	system.Execute()

	instruction := (*system.cpu.instructionSet)[0x0]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = 1", system.cpu.ProgramCounter)
		t.Fail()
	}
}

func TestLdBcNn(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0
	system.cpu.registers.C = 0
	system.cpu.ProgramCounter = romBank00BaseAddress
	system.cpu.mmu.writeByte(romBank00BaseAddress, 0x1)
	system.cpu.mmu.writeWord(romBank00BaseAddress+1, 0xFF0F)

	system.Execute()
	instruction := (*system.cpu.instructionSet)[0x1]
	if system.cpu.ProgramCounter != uint16(instruction.length) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, 3)
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
