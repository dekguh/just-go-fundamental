package helpers

import "math"

type Box struct {
	h, w, l float64
}

type Circle struct {
	r float64
}

func (b Box) AreaFormula() float64 {
	return 2 * (b.h*b.w + b.h*b.l + b.w*b.l)
}

func (c Circle) AreaFormula() float64 {
	return math.Pi * c.r * c.r
}
