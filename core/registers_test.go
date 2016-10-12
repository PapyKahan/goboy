package core

import "testing"

func TestReadBC(t *testing.T) {
	r := registers{}
	r.B = 0xFF
	r.C = 0x0F

	value := r.readBC()
	if value != 0xFF0F {
		t.Logf("ReadBC value = %X, expected = %X", value, 0xFF0F)
		t.Fail()
	}
}

func TestWriteBC(t *testing.T) {
	r := registers{}
	r.writeBC(0xFF0F)

	if r.B != 0xFF || r.C != 0x0F {
		t.Logf("registers.B = %X, expected = %X", r.B, 0xFF)
		t.Logf("registers.C = %X, expected = %X", r.C, 0x0F)
		t.Fail()
	}
}

func TestReadHL(t *testing.T) {
	r := registers{}
	r.H = 0xFF
	r.L = 0x0F

	value := r.readHL()
	if value != 0xFF0F {
		t.Logf("ReadHL value = %X, expected = %X", value, 0xFF0F)
		t.Fail()
	}
}

func TestWriteHL(t *testing.T) {
	r := registers{}
	r.writeHL(0xFF0F)

	if r.H != 0xFF || r.L != 0x0F {
		t.Logf("registers.H = %X, expected = %X", r.H, 0xFF)
		t.Logf("registers.L = %X, expected = %X", r.L, 0x0F)
		t.Fail()
	}
}

func TestReadDE(t *testing.T) {
	r := registers{}
	r.D = 0xF0
	r.E = 0xFF
	value := r.readDE()
	if value != 0xF0FF {
		t.Errorf("ReadDE value = 0x%X, expected = 0x%X", value, 0xF0FF)
	}
}

func TestWriteDE(t *testing.T) {
	r := registers{}
	r.writeDE(0xFF0F)

	if r.D != 0xFF {
		t.Errorf("registers.D = 0x%X, expected = 0x%X", r.D, 0xFF)

	}

	if r.E != 0x0F {
		t.Errorf("registers.E = %X, expected = %X", r.E, 0x0F)
	}
}
