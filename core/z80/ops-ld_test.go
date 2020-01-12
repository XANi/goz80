package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestLD (t *testing.T) {
	cpu := InitCPU()
	cpu.BC[1]=3
	cpu.DE[1]=4
	cpu.AF=[2]byte{0x10,0x12}
	t.Run("LD (BC),a", func(t *testing.T) {
		cpu.ILD_PBC_A([]byte{})
		assert.Equal(t, uint8(0x10), cpu.Data[3])
	})
	t.Run("LD (DE),a", func(t *testing.T) {
		cpu.ILD_PDE_A([]byte{})
		assert.Equal(t, uint8(0x10), cpu.Data[4])
	})
}


func BenchmarkLD_BC(b *testing.B) {
	c := InitCPU()
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.ILD_BC([]byte{0,uint8(n%256)})
	}
}