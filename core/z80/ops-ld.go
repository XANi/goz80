package z80

import (
	"encoding/binary"
)


func (c *CPU) ILD_BC (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	c.BC[1] = c.Data[addr]
	c.BC[0] = c.Data[addr+1]
}
func (c *CPU) ILD_DE (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	c.DE[1] = c.Data[addr]
	c.DE[0] = c.Data[addr+1]
}
func (c *CPU) ILD_HL (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	c.HL[1] = c.Data[addr]
	c.HL[0] = c.Data[addr+1]
}
func (c *CPU) ILD_SP (p []byte) {
	addr := binary.BigEndian.Uint16(p[0:2])
	c.SP[1] = c.Data[addr]
	c.SP[0] = c.Data[addr+1]
}

// register to pointer register addr
func (c *CPU) ILD_PBC_A (p []byte) {
	c.Data[binary.BigEndian.Uint16(c.BC[:])]= c.AF[0]
}
func (c *CPU) ILD_PDE_A (p []byte) {
	c.Data[binary.BigEndian.Uint16(c.DE[:])]= c.AF[0]
}


func (c *CPU) ILD_A_PBC(p []byte) {
	c.AF[0] = c.Data[binary.BigEndian.Uint16(c.BC[:])]
}

func (c *CPU) ILD_A_PDE(p []byte) {
	c.AF[0] = c.Data[binary.BigEndian.Uint16(c.BC[:])]
}

func (c *CPU) ILD_A_PHL(p []byte) {
	c.AF[0] = c.Data[binary.BigEndian.Uint16(c.BC[:])]
}

func (c *CPU) ILD_A_PSP(p []byte) {
	c.AF[0] = c.Data[binary.BigEndian.Uint16(c.BC[:])]
}

func (c *CPU) ILD_A_const(p []byte) {
	c.AF[0]= p[0]
}
func (c *CPU) ILD_B_const(p []byte) {
	c.BC[0]= p[0]
}
func (c *CPU) ILD_C_const(p []byte) {
	c.BC[1]= p[0]
}
func (c *CPU) ILD_D_const(p []byte) {
	c.DE[0]= p[0]
}
func (c *CPU) ILD_E_const(p []byte) {
	c.DE[1]= p[0]
}
func (c *CPU) ILD_H_const(p []byte) {
	c.HL[0]= p[0]
}
func (c *CPU) ILD_L_const(p []byte) {
	c.HL[1]= p[0]
}
