package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEX (t *testing.T) {
	cpu := InitCPU()
	cpu.AF = [2]byte{0x11,0x22}
	cpu.IEX_AF_AF([]byte{})
	cpu.AF = [2]byte{0x33,0x44}
	cpu.IEX_AF_AF([]byte{})
	assert.Equal(t,[2]byte{0x11,0x22},cpu.AF)
	cpu.IEX_AF_AF([]byte{})
	assert.Equal(t,[2]byte{0x33,0x44},cpu.AF)
	cpu.IEX_AF_AF([]byte{})
	assert.Equal(t,[2]byte{0x11,0x22},cpu.AF)
}