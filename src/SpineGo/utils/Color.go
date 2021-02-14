package utils

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func (c *Color) RGBA8888ToColor(value uint) {
	c.R = float32(((value & 0xff000000) >> 24) / 255)
	c.G = float32(((value & 0x00ff0000) >> 16) / 255)
	c.B = float32(((value & 0x0000ff00) >> 8) / 255)
	c.A = float32((value & 0x000000ff) / 255)
}
