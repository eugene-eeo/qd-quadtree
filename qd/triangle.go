package main

import "github.com/paulmach/go.geo"

// Triangle represents a Triangle with vertices (in no particular
// order) A, B, C.
type Triangle struct {
	A, B, C *geo.Point
	lines   []*geo.Line
}

// NewTriangleFromPoints creates a new triangle from the given points.
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

func lineIntersectsBound(l *geo.Line, b *geo.Bound) bool {
	nw := b.NorthWest()
	ne := b.NorthEast()
	sw := b.SouthWest()
	se := b.SouthEast()
	return l.Intersects(geo.NewLine(nw, ne)) ||
		l.Intersects(geo.NewLine(sw, se)) ||
		l.Intersects(geo.NewLine(sw, nw)) ||
		l.Intersects(geo.NewLine(se, sw))
}

// IsWithin checks if a triangle is contained within a given rectangle,
// i.e. if any of its vertices are inside, or if any of it's edges
// intersects the rectangle's edges.
func (t *Triangle) IsWithin(b *geo.Bound) bool {
	// one of the vertices is in the box.
	if b.Contains(t.A) || b.Contains(t.B) || b.Contains(t.C) {
		return true
	}
	for _, line := range t.lines {
		if lineIntersectsBound(line, b) {
			return true
		}
	}
	return false
}

// ContainsPoint checks if a given point is within the triangle.
func (t *Triangle) ContainsPoint(p *geo.Point) bool {
	b0 := t.lines[0].Side(p) > 0
	b1 := t.lines[1].Side(p) > 0
	b2 := t.lines[2].Side(p) > 0
	return (b0 == b1) && (b1 == b2)
}

// ContainsBound checks if the given geo.Bound is within the triangle,
// i.e. if ContainsPoint returns true for all of its vertices.
func (t *Triangle) ContainsBound(b *geo.Bound) bool {
	return t.ContainsPoint(b.NorthWest()) &&
		t.ContainsPoint(b.NorthEast()) &&
		t.ContainsPoint(b.SouthEast()) &&
		t.ContainsPoint(b.SouthWest())
}
