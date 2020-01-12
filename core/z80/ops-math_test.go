package z80

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	cpu := InitCPU()
	var inctest16 = []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
	}{
		{ Name: "INC BC", Func: cpu.IINC_BC, Register: &cpu.BC },
		{ Name: "INC DE", Func: cpu.IINC_DE, Register: &cpu.DE },
		{ Name: "INC HL", Func: cpu.IINC_HL, Register: &cpu.HL },
		{ Name: "INC SP", Func: cpu.IINC_SP, Register: &cpu.SP },
	}
	for _, td := range  inctest16 {
		t.Run(td.Name, func(t *testing.T) {
			(*td.Register) = [2]byte{0, 0}
			td.Func([]byte{})
			assert.Equal(t, uint16(1), binary.BigEndian.Uint16((*td.Register)[:]))
			(*td.Register) = [2]byte{0, 0xff}
			td.Func([]byte{})
			assert.Equal(t, uint16(256), binary.BigEndian.Uint16((*td.Register)[:]))
			(*td.Register) = [2]byte{0xff, 0xff}
			td.Func([]byte{})
			assert.Equal(t, uint16(0), binary.BigEndian.Uint16((*td.Register)[:]))
		})
	}
	var inctest8 =  []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
		Byte int
	}{
		{ Name: "INC A", Func: cpu.IINC_A, Register: &cpu.AF, Byte: 0 },
		{ Name: "INC B", Func: cpu.IINC_B, Register: &cpu.BC, Byte: 0 },
		{ Name: "INC C", Func: cpu.IINC_C, Register: &cpu.BC, Byte: 1 },
		{ Name: "INC D", Func: cpu.IINC_D, Register: &cpu.DE, Byte: 0 },
		{ Name: "INC E", Func: cpu.IINC_E, Register: &cpu.DE, Byte: 1 },
		{ Name: "INC H", Func: cpu.IINC_H, Register: &cpu.HL, Byte: 0 },
		{ Name: "INC L", Func: cpu.IINC_L, Register: &cpu.HL, Byte: 1 },
	}
	for _, td := range  inctest8 {
		t.Run(td.Name, func(t *testing.T) {
			*td.Register = [2]byte{0, 0}
			td.Func([]byte{})
			assert.Equal(t, uint8(1), (*td.Register)[td.Byte])
			(*td.Register)[td.Byte] = 0x7f
			td.Func([]byte{})
			assert.Equal(t, uint8(128), (*td.Register)[td.Byte])
			assert.Equal(t, true, cpu.GetF_PV(), "PV bit")
			assert.Equal(t, true, cpu.GetF_S(), "S bit")
			// TODO check half-carry
		})
	}
}

func TestDec(t *testing.T) {
	cpu := InitCPU()
var dectest8 =  []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
		Byte int
	}{
		{ Name: "DEC A", Func: cpu.IDEC_A, Register: &cpu.AF, Byte: 0 },
		{ Name: "DEC B", Func: cpu.IDEC_B, Register: &cpu.BC, Byte: 0 },
		{ Name: "DEC C", Func: cpu.IDEC_C, Register: &cpu.BC, Byte: 1 },
		{ Name: "DEC D", Func: cpu.IDEC_D, Register: &cpu.DE, Byte: 0 },
		{ Name: "DEC E", Func: cpu.IDEC_E, Register: &cpu.DE, Byte: 1 },
		{ Name: "DEC H", Func: cpu.IDEC_H, Register: &cpu.HL, Byte: 0 },
		{ Name: "DEC L", Func: cpu.IDEC_L, Register: &cpu.HL, Byte: 1 },
	}
	for _, td := range  dectest8 {
		t.Run(td.Name, func(t *testing.T) {
			*td.Register = [2]byte{0, 0}
			(*td.Register)[td.Byte] = 1
			cpu.SetF_Z(false)
			td.Func([]byte{})
			assert.Equal(t, uint8(0), (*td.Register)[td.Byte])
			assert.Equal(t,true,cpu.GetF_Z(), "Z flag true")
			td.Func([]byte{})
			assert.Equal(t, uint8(255), (*td.Register)[td.Byte])
			assert.Equal(t,false,cpu.GetF_Z(), "Z flag false")
			assert.Equal(t,true,cpu.GetF_S(), "S flag true")


		})
	}
}
