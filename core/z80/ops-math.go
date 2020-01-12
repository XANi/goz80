package z80

func (c *CPU) IINC_BC (p []byte) {
	if c.BC[1]== 0xff {
		c.BC[0]++
	}
	c.BC[1]++
}
func (c *CPU) IINC_DE (p []byte) {
	if c.DE[1]== 0xff {
		c.DE[0]++
	}
	c.DE[1]++
}
func (c *CPU) IINC_HL (p []byte) {
	if c.HL[1]== 0xff {
		c.HL[0]++
	}
	c.HL[1]++
}
func (c *CPU) IINC_SP (p []byte) {
	if c.SP[1]== 0xff {
		c.SP[0]++
	}
	c.SP[1]++
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
func (c *CPU) IINC_PHL (p []byte) {

}

func (c *CPU) IDEC_BC (p []byte) {

}
func (c *CPU) IDEC_DE (p []byte) {

}
func (c *CPU) IDEC_HL (p []byte) {

}
func (c *CPU) IDEC_SP (p []byte) {

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

}