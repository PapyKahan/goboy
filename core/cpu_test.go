package core

import "testing"

func TestNop(t *testing.T) {
	system := New()
	system.cpu.ProgramCounter = 1
	instruction := (*system.cpu.instructionSet)[0x0]
	instruction.execute()
	if system.cpu.ProgramCounter != 1 {
		t.Logf("system.cpu.ProgramCounter = %d, expected = 2", system.cpu.ProgramCounter)
		t.Fail()
	}
}

// func TestLdBcNn(t *testing.T) {
// 	system := New()
// 	system.cpu.registers.B = 0
// 	system.cpu.registers.C = 0
// 	system.cpu.ProgramCounter = 1

// 	instruction := (*system.cpu.instructionSet)[0x1]
// 	instruction.execute()
// 	fmt.Println("initialize instruciton set")

// 	if system.cpu.ProgramCounter != uint16(instruction.length) {
// 		t.Logf("system.cpu.ProgramCounter = %d, expected = %d", system.cpu.ProgramCounter, 3)
// 		t.Fail()
// 	}
// }
