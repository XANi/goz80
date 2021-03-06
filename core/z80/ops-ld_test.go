package z80

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestLD (t *testing.T) {
	cpu := InitCPU()
	var tests = []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
	}{
		{Name: "LD BC, **", Func: cpu.ILD_BC, Register: &cpu.BC },
		{Name: "LD DE, **", Func: cpu.ILD_DE, Register: &cpu.DE },
		{Name: "LD HL, **", Func: cpu.ILD_HL, Register: &cpu.HL },
		{Name: "LD SP, **", Func: cpu.ILD_SP, Register: &cpu.SP },
	}
	for id, td := range  tests {
		t.Run(td.Name, func(t *testing.T) {
			binary.LittleEndian.PutUint16(cpu.Data[123:125],uint16(id+1))
			addr := make([]byte, 2)
			binary.BigEndian.PutUint16(addr, 123)
			td.Func(addr)
			assert.Equal(t, id+1, int(binary.BigEndian.Uint16((*td.Register)[:])))
		})
	}
}

func TestLD_P_A (t *testing.T) {
	cpu := InitCPU()
	var tests = []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
		RegIdx int
	}{
		{Name: "LD (BC), A", Func: cpu.ILD_PBC_A, Register: &cpu.BC,RegIdx: 0},
		{Name: "LD (DE), A", Func: cpu.ILD_PDE_A, Register: &cpu.DE,RegIdx: 0},
	}
	for id, td := range  tests {
		t.Run(td.Name, func(t *testing.T) {
			cpu.AF[0]= uint8(id+1)
			binary.BigEndian.PutUint16((*td.Register)[:],555)
			td.Func([]byte{})
			assert.Equal(t,uint8(id + 1),cpu.Data[555])
		})
	}
}


func TestLD_A_PTR(t *testing.T) {
	cpu := InitCPU()
	var tests = []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
	}{
		{ Name: "LD A, (BC)", Func: cpu.ILD_A_PBC, Register: &cpu.BC },
		{ Name: "LD A, (DE)", Func: cpu.ILD_A_PDE, Register: &cpu.DE },
		{ Name: "LD A, (HL)", Func: cpu.ILD_A_PHL, Register: &cpu.HL },
		{ Name: "LD A, (SP)", Func: cpu.ILD_A_PSP, Register: &cpu.SP },
	}
	for _, td := range  tests {
		t.Run(td.Name, func(t *testing.T) {
			cpu.Data[123] = 0xbd
			(*td.Register) = [2]byte{0,123}
			td.Func([]byte{})
			assert.Equal(t, uint8(0xbd), cpu.AF[0], "load to A")
		})
	}
}

func TestLD_X_const(t *testing.T) {
	cpu := InitCPU()
	var tests = []struct {
		Name string
		Func     func([]byte)
		Register *[2]byte
		RegIdx int
	}{
		{ Name: "LD A, (int)", Func: cpu.ILD_A_const, Register:&cpu.AF,RegIdx:0 },
		{ Name: "LD B, (int)", Func: cpu.ILD_B_const, Register:&cpu.BC,RegIdx:0 },
		{ Name: "LD C, (int)", Func: cpu.ILD_C_const, Register:&cpu.BC,RegIdx:1 },
		{ Name: "LD D, (int)", Func: cpu.ILD_D_const, Register:&cpu.DE,RegIdx:0 },
		{ Name: "LD E, (int)", Func: cpu.ILD_E_const, Register:&cpu.DE,RegIdx:1 },
		{ Name: "LD H, (int)", Func: cpu.ILD_H_const, Register:&cpu.HL,RegIdx:0 },
		{ Name: "LD L, (int)", Func: cpu.ILD_L_const, Register:&cpu.HL,RegIdx:1 },
	}
	for idx, td := range  tests {
		t.Run(td.Name, func(t *testing.T) {
			v := uint8(idx+1)
			td.Func([]byte{v})
			assert.Equal(t,v,(*td.Register)[td.RegIdx])
		})
	}
}





func BenchmarkLD_BC(b *testing.B) {
	c := InitCPU()
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.ILD_BC([]byte{0,uint8(n%256)})
	}
}

