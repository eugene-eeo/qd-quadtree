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

func min(a float64, xs... float64) float64 {
	min := a
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

func max(a float64, xs... float64) float64 {
	max := a
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}

func NewRangeFromMesh(mesh []*Triangle) *Range {
	minX := mesh[0].A.X
	maxX := mesh[0].A.X
	minY := mesh[0].A.Y
	maxY := mesh[0].A.Y
	for _, t := range mesh {
		X := []float64{t.A.X, t.B.X, t.C.X}
		Y := []float64{t.A.Y, t.B.Y, t.C.Y}
		minX = min(minX, X...)
		maxX = max(maxX, X...)
		minY = min(minY, Y...)
		maxY = max(maxY, Y...)
	}
	return NewRange(minX, maxX, minY, maxY)
}

func (r *Range) Contains(p *Point) bool {
	return within(p.X, r.X0, r.X1) && within(p.Y, r.Y0, r.Y1)
}
