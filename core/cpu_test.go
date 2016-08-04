package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.ProgramCounter = 0
	instruction := (*system.cpu.instructionSet)[0x0]
	instruction.execute()
	if system.cpu.ProgramCounter != 1 {
		t.Fail()
	}
}
