package utils

type Color struct {
	r float32
	g float32
	b float32
	a float32
}

func NewColor() *Color {
	c := new(Color)
	c.r = 0
	c.g = 0
	c.b = 0
	c.a = 0
	return c
}
