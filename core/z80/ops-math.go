package z80

import (
	"encoding/binary"
)

func (c *CPU) IINC_BC (p []byte) {
	c.BC[1]++
	if c.BC[1] == 0x00 {
		c.BC[0]++
	}
}
func (c *CPU) IINC_DE (p []byte) {
	c.DE[1]++
	if c.DE[1] == 0x00 {
		c.DE[0]++
	}
}
func (c *CPU) IINC_HL (p []byte) {
	c.HL[1]++
	if c.HL[1] == 0x00 {
		c.HL[0]++
	}
}
func (c *CPU) IINC_SP (p []byte) {
	c.SP[1]++
	if c.SP[1] == 0x00 {
		c.SP[0]++
	}
}
func (c *CPU) IINC_A (p []byte) {
	c.SetF_PV( c.AF[0] == 0x7f )
	c.SetF_H( (((c.AF[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.AF[0]++
	c.SetF_S( (c.AF[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.AF[0]==0)

}
func (c *CPU) IINC_B (p []byte) {
	c.SetF_PV( c.BC[0] == 0x7f )
	c.SetF_H( (((c.BC[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.BC[0]++
	c.SetF_S( (c.BC[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.BC[0]==0)
}
func (c *CPU) IINC_C (p []byte) {
	c.SetF_PV( c.BC[1] == 0x7f )
	c.SetF_H( (((c.BC[1] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.BC[1]++
	c.SetF_S( (c.BC[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.BC[1]==0)
}
func (c *CPU) IINC_D (p []byte) {
	c.SetF_PV( c.DE[0] == 0x7f )
	c.SetF_H( (((c.DE[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.DE[0]++
	c.SetF_S( (c.DE[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.DE[0]==0)
}
func (c *CPU) IINC_E (p []byte) {
	c.SetF_PV( c.DE[1] == 0x7f )
	c.SetF_H( (((c.DE[1] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.DE[1]++
	c.SetF_S( (c.DE[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.DE[1]==0)
}
func (c *CPU) IINC_H (p []byte) {
	c.SetF_PV( c.HL[0] == 0x7f )
	c.SetF_H( (((c.HL[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.HL[0]++
	c.SetF_S( (c.HL[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.HL[0]==0)
}
func (c *CPU) IINC_L (p []byte) {
	c.SetF_PV( c.HL[1] == 0x7f )
	c.SetF_H( (((c.HL[1] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
	c.HL[1]++
	c.SetF_S( (c.HL[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.HL[1]==0)
}
//func (c *CPU) IINC_PHL (p []byte) {
//	panic("op not implemented")
//}

func (c *CPU) IDEC_BC (p []byte) {
	c.BC[1]--
	if c.BC[1] == 0xff {
		c.BC[0]--
	}
}
func (c *CPU) IDEC_DE (p []byte) {
	c.DE[1]--
	if c.DE[1] == 0xff {
		c.DE[0]--
	}

}
func (c *CPU) IDEC_HL (p []byte) {
	c.HL[1]--
	if c.HL[1] == 0xff {
		c.HL[0]--
	}

}
func (c *CPU) IDEC_SP (p []byte) {
	c.SP[1]--
	if c.SP[1] == 0xff {
		c.SP[0]--
	}

}
func (c *CPU) IDEC_A (p []byte) {
	c.SetF_PV( c.AF[0] == 0x80 )
	c.AF[0]--
	c.SetF_S( (c.AF[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.AF[0]==0)

}
func (c *CPU) IDEC_B (p []byte) {
	c.SetF_PV( c.BC[0] == 0x80 )
	c.BC[0]--
	c.SetF_S( (c.BC[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.BC[0]==0)
}
func (c *CPU) IDEC_C (p []byte) {
	c.SetF_PV( c.BC[1] == 0x80 )
	c.BC[1]--
	c.SetF_S( (c.BC[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.BC[1]==0)
}
func (c *CPU) IDEC_D (p []byte) {
	c.SetF_PV( c.DE[0] == 0x80 )
	c.DE[0]--
	c.SetF_S( (c.DE[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.DE[0]==0)

}
func (c *CPU) IDEC_E (p []byte) {
	c.SetF_PV( c.DE[1] == 0x80 )
	c.DE[1]--
	c.SetF_S( (c.DE[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.DE[1]==0)

}
func (c *CPU) IDEC_H (p []byte) {
	c.SetF_PV( c.HL[0] == 0x80 )
	c.HL[0]--
	c.SetF_S( (c.HL[0] & (1 << 7)) != 0 )
	c.SetF_Z(c.HL[0]==0)

}
func (c *CPU) IDEC_L (p []byte) {
	c.SetF_PV( c.HL[1] == 0x80 )
	c.HL[1]--
	c.SetF_S( (c.HL[1] & (1 << 7)) != 0 )
	c.SetF_Z(c.HL[1]==0)

}
func (c *CPU) IDEC_PHL (p []byte) {
	panic("op not implemented")
}
func (c *CPU) IADD_HL_BC(p []byte) {
	x := binary.BigEndian.Uint16(c.HL[:])
	y := binary.BigEndian.Uint16(c.BC[:])
	sum32 := uint32(x) + uint32(y) //+ uint32(carry)
   	sum := uint16(sum32)
   	carryOut := uint16(sum32 >> 16)
   	binary.BigEndian.PutUint16(c.HL[:],sum)
	c.SetF_C(carryOut > 0)
   	c.SetF_PV(carryOut > 1)
   	c.SetF_H( (((c.HL[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )
}
func (c *CPU) IADD_HL_DE(p []byte) {
	x := binary.BigEndian.Uint16(c.HL[:])
	y := binary.BigEndian.Uint16(c.DE[:])
	sum32 := uint32(x) + uint32(y) //+ uint32(carry)
   	sum := uint16(sum32)
   	carryOut := uint16(sum32 >> 16)
   	binary.BigEndian.PutUint16(c.HL[:],sum)
	c.SetF_C(carryOut > 0)
   	c.SetF_PV(carryOut > 1)
	c.SetF_H( (((c.HL[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )

}
func (c *CPU) IADD_HL_HL(p []byte) {
	x := binary.BigEndian.Uint16(c.HL[:])
	y := binary.BigEndian.Uint16(c.HL[:])
	sum32 := uint32(x) + uint32(y) //+ uint32(carry)
   	sum := uint16(sum32)
   	carryOut := uint16(sum32 >> 16)
   	binary.BigEndian.PutUint16(c.HL[:],sum)
	c.SetF_C(carryOut > 0)
   	c.SetF_PV(carryOut > 1)
	c.SetF_H( (((c.HL[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )

}
func (c *CPU) IADD_HL_SP(p []byte) {
	x := binary.BigEndian.Uint16(c.HL[:])
	y := binary.BigEndian.Uint16(c.SP[:])
	sum32 := uint32(x) + uint32(y) //+ uint32(carry)
   	sum := uint16(sum32)
   	carryOut := uint16(sum32 >> 16)
   	binary.BigEndian.PutUint16(c.HL[:],sum)
	c.SetF_C(carryOut > 0)
   	c.SetF_PV(carryOut > 1)
	c.SetF_H( (((c.HL[0] & 0xf) + (1& 0xf)) & 0x10) == 0x10 )


}
