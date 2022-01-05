package z80

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	cpu := InitCPU()

	var inctest8 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
		Byte     int
	}{
		{Name: "INC A", Func: cpu.IINC_A, Register: &cpu.AF, Byte: 0},
		{Name: "INC B", Func: cpu.IINC_B, Register: &cpu.BC, Byte: 0},
		{Name: "INC C", Func: cpu.IINC_C, Register: &cpu.BC, Byte: 1},
		{Name: "INC D", Func: cpu.IINC_D, Register: &cpu.DE, Byte: 0},
		{Name: "INC E", Func: cpu.IINC_E, Register: &cpu.DE, Byte: 1},
		{Name: "INC H", Func: cpu.IINC_H, Register: &cpu.HL, Byte: 0},
		{Name: "INC L", Func: cpu.IINC_L, Register: &cpu.HL, Byte: 1},
	}
	for _, td := range inctest8 {
		t.Run(td.Name, func(t *testing.T) {
			*td.Register = [2]byte{0, 0}
			td.Func([]byte{})
			assert.Equal(t, uint8(1), (*td.Register)[td.Byte])
			assertMathFlags8(t, cpu, 0, (*td.Register)[td.Byte])

			(*td.Register)[td.Byte] = 0x7E
			td.Func([]byte{})
			assert.Equal(t, uint8(127), (*td.Register)[td.Byte])
			assertMathFlags8(t, cpu, 0x7E, (*td.Register)[td.Byte])

			(*td.Register)[td.Byte] = 0x7F
			td.Func([]byte{})
			assert.Equal(t, uint8(128), (*td.Register)[td.Byte])
			assertMathFlags8(t, cpu, 0x7F, (*td.Register)[td.Byte])

			(*td.Register)[td.Byte] = 0xFF
			td.Func([]byte{})
			assert.Equal(t, uint8(0), (*td.Register)[td.Byte])
			assertMathFlags8(t, cpu, 0xFF, (*td.Register)[td.Byte])
			// TODO check half-carry
		})
	}
	var inctest16 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
	}{
		{Name: "INC BC", Func: cpu.IINC_BC, Register: &cpu.BC},
		{Name: "INC DE", Func: cpu.IINC_DE, Register: &cpu.DE},
		{Name: "INC HL", Func: cpu.IINC_HL, Register: &cpu.HL},
		{Name: "INC SP", Func: cpu.IINC_SP, Register: &cpu.SP},
	}
	for _, td := range inctest16 {
		t.Run(td.Name, func(t *testing.T) {
			td.Func([]byte{})
			assert.Equal(t, uint16(1), binary.BigEndian.Uint16((*td.Register)[:]))
			binary.BigEndian.PutUint16((*td.Register)[:], 255)
			td.Func([]byte{})
			assert.Equal(t, uint16(256), binary.BigEndian.Uint16((*td.Register)[:]))
			(*td.Register) = [2]byte{0xff, 0xff}
			td.Func([]byte{})
			assert.Equal(t, uint16(0), binary.BigEndian.Uint16((*td.Register)[:]))
		})
	}
}

func assertMathFlags8(t *testing.T, cpu *CPU, in byte, out byte) {
	assert.Equal(t, int8(out) < 0, cpu.GetF_S(), "(S)ign flag")
	assert.Equal(t, out == 0, cpu.GetF_Z(), "(Z)ero flag")
	if in == 0x7f && out > 0x7f ||
		in == 0x80 && out < 0x80 {
		assert.Equal(t, true, cpu.GetF_PV(), "PV flag true: %#02x(%d) -> %#02x(%d)", in, int8(in), out, int8(out))
	} else {
		assert.Equal(t, false, cpu.GetF_PV(), "PV flag false: %#02x -> %#02x", in, out)
	}
}

func TestDec(t *testing.T) {
	cpu := InitCPU()
	var dectest8 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
		Byte     int
	}{
		{Name: "DEC A", Func: cpu.IDEC_A, Register: &cpu.AF, Byte: 0},
		{Name: "DEC B", Func: cpu.IDEC_B, Register: &cpu.BC, Byte: 0},
		{Name: "DEC C", Func: cpu.IDEC_C, Register: &cpu.BC, Byte: 1},
		{Name: "DEC D", Func: cpu.IDEC_D, Register: &cpu.DE, Byte: 0},
		{Name: "DEC E", Func: cpu.IDEC_E, Register: &cpu.DE, Byte: 1},
		{Name: "DEC H", Func: cpu.IDEC_H, Register: &cpu.HL, Byte: 0},
		{Name: "DEC L", Func: cpu.IDEC_L, Register: &cpu.HL, Byte: 1},
	}
	for _, td := range dectest8 {
		t.Run(td.Name, func(t *testing.T) {
			*td.Register = [2]byte{0, 0}
			(*td.Register)[td.Byte] = 1
			cpu.SetF_Z(false)
			td.Func([]byte{})
			assert.Equal(t, uint8(0), (*td.Register)[td.Byte], "DEC 1 -> 0")
			assertMathFlags8(t, cpu, 0x01, (*td.Register)[td.Byte])

			(*td.Register)[td.Byte] = 0
			td.Func([]byte{})
			assert.Equal(t, uint8(255), (*td.Register)[td.Byte], "DEC 0 -> -1")
			assertMathFlags8(t, cpu, 0x00, (*td.Register)[td.Byte])

			(*td.Register)[td.Byte] = 0x80
			td.Func([]byte{})
			assert.Equal(t, uint8(127), (*td.Register)[td.Byte], "DEC -128 -> 127")
			assertMathFlags8(t, cpu, 0x80, (*td.Register)[td.Byte])
		})
	}
	var inctest16 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
	}{
		{Name: "DEC BC", Func: cpu.IDEC_BC, Register: &cpu.BC},
		{Name: "DEC DE", Func: cpu.IDEC_DE, Register: &cpu.DE},
		{Name: "DEC HL", Func: cpu.IDEC_HL, Register: &cpu.HL},
		{Name: "DEC SP", Func: cpu.IDEC_SP, Register: &cpu.SP},
	}
	for _, td := range inctest16 {
		t.Run(td.Name, func(t *testing.T) {
			(*td.Register) = [2]byte{0, 0x01}
			td.Func([]byte{})
			assert.Equal(t, uint16(0), binary.BigEndian.Uint16((*td.Register)[:]))
			(*td.Register) = [2]byte{0, 0x00}
			td.Func([]byte{})
			assert.Equal(t, uint16(65535), binary.BigEndian.Uint16((*td.Register)[:]))
		})
	}
}

func TestAdd(t *testing.T) {
	cpu := InitCPU()
	var addtest8 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
		Byte     int
	}{}
	_ = addtest8

	var test16 = []struct {
		Name     string
		Func     func([]byte)
		Register *[2]byte
		Byte     int
	}{
		{Name: "DEC BC", Func: cpu.IADD_HL_BC, Register: &cpu.BC},
		{Name: "DEC DE", Func: cpu.IADD_HL_DE, Register: &cpu.DE},
		{Name: "DEC HL", Func: cpu.IADD_HL_HL, Register: &cpu.HL},
		{Name: "DEC SP", Func: cpu.IADD_HL_SP, Register: &cpu.SP},
	}
	// TODO test half carry

	for _, td := range test16 {
		t.Run(td.Name, func(t *testing.T) {
			binary.BigEndian.PutUint16((cpu.HL)[:], 255)
			binary.BigEndian.PutUint16((*td.Register)[:], 255)
			cpu.SetF_C(true)
			cpu.SetF_PV(true)
			td.Func([]byte{})
			assert.Equal(t, 255*2, int(binary.BigEndian.Uint16(cpu.HL[:])))
			assert.Equal(t, false, cpu.GetF_C(), "C(arry) flag true")
			assert.Equal(t, false, cpu.GetF_PV(), "PV flag true")

			binary.BigEndian.PutUint16((cpu.HL)[:], (1 << 15))
			binary.BigEndian.PutUint16((*td.Register)[:], (1 << 15))
			td.Func([]byte{})
			assert.Equal(t, 0, int(binary.BigEndian.Uint16(cpu.HL[:])))
			assert.Equal(t, true, cpu.GetF_C(), "C(arry) flag true")
			assert.Equal(t, false, cpu.GetF_PV(), "PV flag true")

			cpu.HL = [2]byte{0xFF, 0xFF}
			(*td.Register) = [2]byte{0xFF, 0xFF}
			td.Func([]byte{})
			assert.Equal(t, 65534, int(binary.BigEndian.Uint16(cpu.HL[:])))
			assert.Equal(t, true, cpu.GetF_C(), "C(arry) flag true")
			assert.Equal(t, false, cpu.GetF_PV(), "PV flag true")
		})
	}
}

func BenchmarkADD_HL_BC(b *testing.B) {
	c := InitCPU()
	// run the Fib function b.N times
	c.HL = [2]byte{0x01, 0x91}
	c.BC = [2]byte{0x01, 0x91}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IADD_HL_BC([]byte{})
	}
}

func BenchmarkINC_A(b *testing.B) {
	c := InitCPU()
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IINC_A([]byte{})
	}
}

func BenchmarkINC_BC(b *testing.B) {
	c := InitCPU()
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IINC_A([]byte{})
	}
}
