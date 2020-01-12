package z80

import "math/bits"

func (c *CPU) IRLCA ([]byte) {
	c.AF[0] = bits.RotateLeft8(c.AF[0],1)
	c.SetF_C((c.AF[0] & 1)==1)
}
func (c *CPU) IRRCA ([]byte) {
	c.SetF_C((c.AF[0] & 1)==1)
	c.AF[0] = bits.RotateLeft8(c.AF[0],-1)
}