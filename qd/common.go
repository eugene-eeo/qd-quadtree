package main

func midpoint(a, b float64) float64 {
	return (a + b) / 2.0
}

func within(x, start, end float64) bool {
	return x >= start && x <= end
}

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

type Range struct {
	X0, X1, Y0, Y1 float64
}

func NewRange(x0, x1, y0, y1 float64) *Range {
	return &Range{x0, x1, y0, y1}
}

func (r *Range) Contains(p *Point) bool {
	return within(p.X, r.X0, r.X1) && within(p.Y, r.Y0, r.Y1)
}
