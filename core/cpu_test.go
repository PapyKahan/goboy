package core

import (
	"fmt"
	"testing"
)

func TestNop(t *testing.T) {
	system := New()
	system.cpu.ProgramCounter = 1
	instruction := (*system.cpu.instructionSet)[0x0]
	instruction.execute()
	if system.cpu.ProgramCounter != 1 {
		t.Logf("system.cpu.ProgramCounter = %d, expected = 1", system.cpu.ProgramCounter)
		t.Fail()
	}
}

func TestLdBcNn(t *testing.T) {
	system := New()
	system.cpu.registers.B = 0
	system.cpu.registers.C = 0
	system.cpu.ProgramCounter = 0
	system.cpu.mmu.writeWord(0x0, 0xFF0F)

	instruction := (*system.cpu.instructionSet)[0x1]
	instruction.execute()
	fmt.Println("initialize instruciton set")

	if system.cpu.ProgramCounter != uint16(instruction.length-1) {
		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, 2)
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
