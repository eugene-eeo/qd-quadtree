package main

import "github.com/paulmach/go.geo"
import "math"

func midpoint(a, b float64) float64 {
	return (a + b) / 2.0
}

func BoundFromPoints(xs []*geo.Point) *geo.Bound {
	min_x := xs[0].X()
	min_y := xs[0].Y()
	max_x := xs[0].X()
	max_y := xs[0].Y()
	for _, p := range xs[1:] {
		min_x = math.Min(min_x, p.X())
		min_y = math.Min(min_y, p.Y())
		max_x = math.Max(max_x, p.X())
		max_y = math.Max(max_y, p.Y())
	}
	return geo.NewBound(min_x, max_x, min_y, max_y)
}
