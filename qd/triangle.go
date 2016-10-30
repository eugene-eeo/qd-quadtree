package main

import "github.com/paulmach/go.geo"

type Triangle struct {
	A, B, C *geo.Point
	lines   []*geo.Line
}

func NewTriangleFromPoints(A, B, C *geo.Point) *Triangle {
	return &Triangle{
		A: A, B: B, C: C,
		lines: []*geo.Line{
			geo.NewLine(A, B),
			geo.NewLine(B, C),
			geo.NewLine(C, A),
		},
	}
}

func line_intersects_bound(l *geo.Line, b *geo.Bound) bool {
	nw := b.NorthWest()
	ne := b.NorthEast()
	sw := b.SouthWest()
	se := b.SouthEast()
	return l.Intersects(geo.NewLine(nw, ne)) ||
		l.Intersects(geo.NewLine(sw, se)) ||
		l.Intersects(geo.NewLine(sw, nw)) ||
		l.Intersects(geo.NewLine(se, sw))
}

func (t *Triangle) IsWithin(b *geo.Bound) bool {
	// best case: one of the vertexes is in the box.
	if b.Contains(t.A) || b.Contains(t.B) || b.Contains(t.C) {
		return true
	}
	for _, line := range t.lines {
		if line_intersects_bound(line, b) {
			return true
		}
	}
	return false
}

func (t *Triangle) ContainsPoint(p *geo.Point) bool {
	for _, line := range t.lines {
		if line.Side(p) == -1 { // lies on left
			return false
		}
	}
	return true
}

func (t *Triangle) ContainsRange(b *geo.Bound) bool {
	return t.ContainsPoint(b.NorthWest()) &&
		t.ContainsPoint(b.NorthEast()) &&
		t.ContainsPoint(b.SouthEast()) &&
		t.ContainsPoint(b.SouthWest())
}
