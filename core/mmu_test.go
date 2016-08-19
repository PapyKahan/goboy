package core

import "testing"

func TestReadByte(t *testing.T) {
	system := New()
	system.cpu.mmu.memory[0] = 1
	value := system.cpu.mmu.readByte(0)
	if value != 0 {
		t.Errorf("Readed value = %d, expected = %d", value, 1)
		t.Fail()
	}
}

func TestReadWord(t *testing.T) {
	system := New()
	system.cpu.mmu.memory[0] = 0xFF
	system.cpu.mmu.memory[0] = 0x0F
	value := system.cpu.mmu.readWord(0)
	if value != 0 {
		t.Errorf("Readed value = %d, expected = %d", value, 0xFF0F)
		t.Fail()
	}
}
