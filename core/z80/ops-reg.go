package z80

func (c *CPU) IEX_AF_AF([]byte) {
	c.AF,c.sAF = c.sAF, c.AF

}