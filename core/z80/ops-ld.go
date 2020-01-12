package z80

import "encoding/binary"


func (c *CPU) ILD_BC (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	copy(c.BC[:],c.Data[addr:addr+2])
}
func (c *CPU) ILD_DE (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	copy(c.BC[:],c.Data[addr:addr+2])
}
func (c *CPU) ILD_HL (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	copy(c.BC[:],c.Data[addr:addr+2])
}
func (c *CPU) ILD_SP (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	copy(c.BC[:],c.Data[addr:addr+2])
}

// register to pointer register addr
func (c *CPU) ILD_PBC_A (p []byte) {
	c.Data[binary.BigEndian.Uint16(c.BC[:])]= c.AF[0]
}
func (c *CPU) ILD_PDE_A (p []byte) {
	c.Data[binary.BigEndian.Uint16(c.DE[:])]= c.AF[0]
}



