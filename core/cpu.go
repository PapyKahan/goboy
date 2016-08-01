package core

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

type registers struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte
	F byte
}

type instruction struct {
	exec  func()
	name  string
	ticks uint64
	m     uint64
}

// New creates a new instance of Gameboy's CPU
func New(system *System) *CPU {
	cpu := &CPU{}

	cpu.system = system
	cpu.instructionSet = initializeInstructionset()

	return cpu
}

// TODO create instruction set map.
func initializeInstructionset() map[byte]instruction {
	return nil
}
