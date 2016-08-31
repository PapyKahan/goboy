package core

import "testing"

func TestReadByte(t *testing.T) {
	system := New()
	system.cpu.mmu.memory[0] = 1
	value := system.cpu.mmu.readByte(0)
	if value != 1 {
		t.Errorf("Readed value = %d, expected = %d", value, 1)
		t.Fail()
	}
}

func TestReadWord(t *testing.T) {
	system := New()
	system.cpu.mmu.memory[0] = 0xFF
	system.cpu.mmu.memory[1] = 0x0F
	value := system.cpu.mmu.readWord(0)
	if value != 0xFF0F {
		t.Errorf("Readed value = %d, expected = %d", value, 0xFF0F)
		t.Fail()
	}
}

func TestWriteWord(t *testing.T) {
	system := New()
	system.cpu.mmu.writeWord(0x0, 0xFF0F)
	if system.cpu.mmu.memory[0x0] != 0xFF || system.cpu.mmu.memory[0x1] != 0x0F {
		t.Fail()
	}
}

func TestWriteByte(t *testing.T) {
	system := New()
	system.cpu.mmu.writeByte(0x0, 0xF0)
	if system.cpu.mmu.memory[0x0] != 0xF0 {
		t.Fail()
	}
}
