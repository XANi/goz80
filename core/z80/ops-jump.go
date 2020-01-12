package z80

func (c *CPU) IJPXX(p []byte) {
	copy(c.PC[:],p[0:2])
}