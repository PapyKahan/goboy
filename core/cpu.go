package core

import "fmt"

func init() {

}

// CPU GameBoy cpu
type CPU struct {
	system         *System
	registers      registers
	ProgramCounter uint16
	StackPointer   uint16
	ticks          uint64
	m              uint64
	instructionSet map[byte]instruction
}

type instruction struct {
	Execute func(cpu *CPU)
	Name    string
	Ticks   uint64
	Length  uint8
}

// New creates a new instance of Gameboy's CPU
func New(system *System) *CPU {
	cpu := &CPU{}
	cpu.system = system
	cpu.initializeInstructionset()
	return cpu
}

func nop(cpu *CPU) {
	cpu.StackPointer += 2
}

func (cpu *CPU) initializeInstructionset() error {
	cpu.instructionSet = make(map[byte]instruction)
	if cpu.instructionSet == nil {
		return fmt.Errorf("Could not allocate memory for cpu.instructionSet")
	}

	cpu.instructionSet[0x0] = instruction{Name: "NOP", Ticks: 4, Length: 1, Execute: nop}
	return nil
}

// Execute execute next instruction
func (cpu *CPU) Execute() error {
	cpu.ProgramCounter++
	opcode := cpu.system.mmu.readByte(cpu.ProgramCounter)

	inst := cpu.instructionSet[byte(opcode&0xFF>>8)]

	switch inst.Length {
	case 0:
		inst.Execute(cpu)
	case 1:
		inst.Execute(cpu)
	case 2:
	}

	return nil
}
