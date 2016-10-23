package main

type Line interface {
	WithinRange(*Range) bool
}

type SlopedLine struct {
	m float64
	c float64
}

type VerticalLine struct {
	x float64
}

func NewLineFromPoints(A, B *Point) Line {
	dx := (A.X - B.X)
	if dx == 0 {
		return &VerticalLine{x: A.X}
	}
	m := (A.Y - B.Y) / dx
	c := A.Y - m*A.X
	return &SlopedLine{m: m, c: c}
}

func (v *VerticalLine) WithinRange(r *Range) bool {
	return within(v.x, r.X0, r.X1)
}

func (l *SlopedLine) YToX(y float64) float64 {
	return (y - l.c) / l.m
}

func (l *SlopedLine) XToY(x float64) float64 {
	return l.m*x + l.c
}

func (l *SlopedLine) WithinRange(r *Range) bool {
	return within(l.XToY(r.X0), r.Y0, r.Y1) ||
		within(l.XToY(r.X1), r.Y0, r.Y1) ||
		within(l.YToX(r.Y0), r.X0, r.X1) ||
		within(l.YToX(r.Y1), r.X0, r.X1)
}

type Triangle struct {
	A, B, C *Point
	lines   []Line
}

func NewTriangleFromPoints(A, B, C *Point) *Triangle {
	return &Triangle{
		A: A, B: B, C: C,
		lines: []Line{
			NewLineFromPoints(A, B),
			NewLineFromPoints(B, C),
			NewLineFromPoints(C, A),
		},
	}
}

func (t *Triangle) IsWithin(r *Range) bool {
	// best case: one of the vertexes is in the box.
	if r.Contains(t.A) || r.Contains(t.B) || r.Contains(t.C) {
		return true
	}
	for _, line := range t.lines {
		if line.WithinRange(r) {
			return true
		}
	}
	return false
}
