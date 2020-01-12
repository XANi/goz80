package z80

import (
	"encoding/binary"
	"fmt"
)

var memSize =65536

type CPU struct {
	AF [2]byte
	BC [2]byte
	DE [2]byte
	HL [2]byte
	SP [2]byte
	PC [2]byte
	IX [2]byte
	IY [2]byte

	I uint8
	R uint8 // DRAM refresh counter

	// shadow registers
	sAF [2]byte
	sBC [2]byte
	sDE [2]byte
	sHL [2]byte
	// memory
	Data []byte
	// current work data
	pc uint16
	op func([]byte)
	opId uint8
	opLen uint8
	// instruction Decode map
	id [256]func(param []byte)()
	// instruction size map
	is [256]uint8
}


func InitCPU() (*CPU) {
	cpu := CPU{
		Data:  make([]byte, memSize),
	}
	cpu.InitializeInstructionDecoder()
	return &cpu
}

func (c *CPU) InitializeInstructionDecoder() {
	funcmap := []struct {
		I uint8  // ID/Op
		F  func(param []byte) // function
		S  uint8              // size
		C  uint               // cycles
	}{
		{I: 0x00, F: c.INOP, S: 1},
		{I: 0x01, F: c.ILD_BC, S: 3},
		{I: 0x02, F: c.ILD_PBC_A, S: 1},
		{I: 0x03, F: c.IINC_BC, S: 1},
		{I: 0x04, F: c.IINC_B, S: 1},
		{I: 0x05, F: c.IDEC_B, S: 1},
		{I: 0x06, F: c.ILD_B_const, S: 2},
		{I: 0x07, F: c.IRLCA, S: 1},
		{I: 0x08, F: c.IEX_AF_AF, S: 1},
		{I: 0x09, F: c.IADD_HL_BC, S: 1},
		{I: 0x0a, F: c.ILD_A_PBC, S: 1},
		{I: 0x0b, F: c.IDEC_BC, S: 1},
		{I: 0x0c, F: c.IINC_C, S: 1},
		{I: 0x0d, F: c.IDEC_C, S: 1},
		{I: 0x0e, F: c.ILD_C_const, S: 2},
		{I: 0x0f, F: c.IDEC_C, S: 1},
		// ...
		{I: 0x3c, F: c.IINC_A, S: 1},
		{I: 0x3e, F: c.ILD_A_const,S: 2},
		// ...
		{I: 0xc3,F: c.IJPXX,S:3},
	}
	for _, op := range funcmap {
		if c.id[op.I] != nil {
			panic(fmt.Sprintf("duplicate for function %d", op.I))
		}
		c.id[op.I] = op.F
		c.is[op.I] = op.S
	}
}

func (c *CPU) Step() {
	c.pc = binary.BigEndian.Uint16(c.PC[:])
	c.opId= c.Data[c.pc]
	c.opLen = c.is[ c.opId ]
	c.op = c.id[c.opId]
	binary.BigEndian.PutUint16(c.PC[:],c.pc + uint16(c.opLen))
	c.op(c.Data[c.pc + 1:c.pc + uint16(c.opLen)])
	fmt.Printf("PC: %d, opid %02x size: %d \n",c.pc,c.opId,c.opLen)

}



// Flags:
// | 7 |  S  |  Sign           | Set if the 2-complement value is negative (copy of MSB)
// | 6 |  Z  |  Zero           | Set if value is zero
// | 5 |  _  |
// | 4 |  H  | Half Carry      |
// | 3 |  _  |                 |
// | 2 | P/V | Parity/Overflow | Parity set if even number of bits set, Overflow set if the 2-complement result does not fit in the register
// | 1 |  N  | Subtract        | last op was subtraction
// | 0 |  C  | Carry           | result did not fit in register


func (c *CPU) GetF_S() bool {
	return (c.AF[1] & (1<<7)) != 0
}
func (c *CPU) SetF_S(f bool) {
	if f {
		c.AF[1] = c.AF[1] | (1<<7)
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1<<7)
	}
}

func (c *CPU) GetF_Z() bool {
	return (c.AF[1] & (1<<6)) != 0
}
func (c *CPU) SetF_Z(f bool) {
	if f {
		c.AF[1] = c.AF[1] | (1<<6)
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1<<6)
	}
}

func (c *CPU) GetF_H() bool {
	return (c.AF[1] & (1<<4)) != 0
}
func (c *CPU) SetF_H(f bool) {
	if f {
		c.AF[1] = c.AF[1] | (1<<4)
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1<<4)
	}
}

func (c *CPU) GetF_PV() bool {
	return (c.AF[1] & (1<<2)) != 0
}
func (c *CPU) SetF_PV(f bool) {
	if f {
		c.AF[1] = c.AF[1] | (1<<2)
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1<<2)
	}
}

func (c *CPU) GetF_N() bool {
	return (c.AF[1] & (1<<1)) != 0
}
func (c *CPU) SetF_N(f bool) {
	if f {
		c.AF[1] = c.AF[1] | (1<<1)
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1<<1)
	}
}

func (c *CPU) GetF_C() bool {
	return (c.AF[1] & 1) != 0
}
func (c *CPU) SetF_C(f bool) {
	if f {
		c.AF[1] = c.AF[1] | 1
	} else {
		c.AF[1] = c.AF[1] & (0xff -  1)
	}
}
