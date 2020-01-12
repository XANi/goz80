package z80

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
	sAF uint16
	sBC uint16
	sDE uint16
	sHL uint16
	// memory
	Data []byte
	// instruction Decode map
	id [256]func(param []byte)()
}


func InitCPU() (*CPU) {
	cpu := CPU{
		Data:  make([]byte, memSize),
	}
	return &cpu
}

func (c *CPU) InstructionDecoder()  [256]func(param []byte)() {
	var d [256]func(param []byte)
	d[0x00] = c.INOP
	d[0x01] = c.ILD_BC
	d[0x02] = c.INOP

	return d
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
