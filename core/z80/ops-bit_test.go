package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLCA (t *testing.T) {
	cpu := InitCPU()
	cpu.AF[0] =0x80
	cpu.IRLCA([]byte{})
	assert.Equal(t,uint8(0x01),cpu.AF[0],"%#08b -> %#08b",0x80,0x01)
	assert.Equal(t,true,cpu.GetF_C(),"carry true")

	cpu.AF[0] =0x40
	cpu.IRLCA([]byte{})
	assert.Equal(t,uint8(0x80),cpu.AF[0],"%#08b -> %#08b",0x40,0x80)
	assert.Equal(t,false,cpu.GetF_C(),"carry true")
}

func TestRRCA (t *testing.T) {
	cpu := InitCPU()
	cpu.AF[0] =0x01
	cpu.IRRCA([]byte{})
	assert.Equal(t,uint8(0x80),cpu.AF[0],"%#08b -> %#08b",0x01,0x80)
	assert.Equal(t,true,cpu.GetF_C(),"carry true")

	cpu.AF[0] =0x80
	cpu.IRRCA([]byte{})
	assert.Equal(t,uint8(0x40),cpu.AF[0],"%#08b -> %#08b",0x80,0x40)
	assert.Equal(t,false,cpu.GetF_C(),"carry true")
}