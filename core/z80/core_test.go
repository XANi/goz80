package z80

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	cpu := InitCPU()
	t.Run("Sign", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_S(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1 << 7
		assert.Equal(t, true, cpu.GetF_S(), fmt.Sprintf("%016b", cpu.AF))
		cpu.AF[1] = 0xff ^ (1 << 7)
		assert.Equal(t, false, cpu.GetF_S(), fmt.Sprintf("%016b", cpu.AF))
		cpu.AF[1] = 0
		cpu.SetF_S(true)
		assert.Equal(t, uint8(1<<7), cpu.AF[1])
		cpu.SetF_S(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})
	t.Run("Zero", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_Z(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1 << 6
		assert.Equal(t, true, cpu.GetF_Z(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0xff ^ (1 << 6)
		assert.Equal(t, false, cpu.GetF_Z(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0
		cpu.SetF_Z(true)
		assert.Equal(t, uint8(1<<6), cpu.AF[1])
		cpu.SetF_Z(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})

	t.Run("Half Carry", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_H(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1 << 4
		assert.Equal(t, true, cpu.GetF_H(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0xff ^ (1 << 4)
		assert.Equal(t, false, cpu.GetF_H(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0
		cpu.SetF_H(true)
		assert.Equal(t, uint8(1<<4), cpu.AF[1])
		cpu.SetF_H(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})
	t.Run("Parity/Overflow", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_PV(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1 << 2
		assert.Equal(t, true, cpu.GetF_PV(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0xff ^ (1 << 2)
		assert.Equal(t, false, cpu.GetF_PV(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0
		cpu.SetF_PV(true)
		assert.Equal(t, uint8(1<<2), cpu.AF[1])
		cpu.SetF_PV(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})
	t.Run("Subtract", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_N(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1 << 1
		assert.Equal(t, true, cpu.GetF_N(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0xff ^ (1 << 1)
		assert.Equal(t, false, cpu.GetF_N(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0
		cpu.SetF_N(true)
		assert.Equal(t, uint8(1<<1), cpu.AF[1])
		cpu.SetF_N(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})
	t.Run("Carry", func(t *testing.T) {
		assert.Equal(t, false, cpu.GetF_C(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 1
		assert.Equal(t, true, cpu.GetF_C(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0xff ^ (1)
		assert.Equal(t, false, cpu.GetF_C(), fmt.Sprintf("%016b", cpu.AF[1]))
		cpu.AF[1] = 0
		cpu.SetF_C(true)
		assert.Equal(t, uint8(1), cpu.AF[1])
		cpu.SetF_C(false)
		assert.Equal(t, uint8(0), cpu.AF[1])

	})
}

func TestProg(t *testing.T) {
	cpu := InitCPU()
	prog := []byte{
		0x0e, 0x0b,
		0x3c,
		0x3c,
		0x02,
		0xC3, 0x00, 0x00}

	copy(cpu.Data[:], prog)
	assert.Contains(t, cpu.State(), "PC: 0")
	for i := 0; i <= 9; i++ {
		cpu.Step()
	}
	assert.Equal(t, 4, int(cpu.AF[0]))
	assert.Equal(t, 4, int(cpu.Data[0x0B]))
}
func BenchmarkProg(b *testing.B) {
	cpu := InitCPU()
	prog := []byte{
		0x0e, 0x0b,
		0x3c,
		0x3c,
		0x02,
		0xC3, 0x00, 0x00}

	copy(cpu.Data[:], prog)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cpu.Step()
	}

}
