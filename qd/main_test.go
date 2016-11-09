package main

import "github.com/paulmach/go.geo"
import "testing"
import "math/rand"

func TestMain(t *testing.T) {
	points := []*geo.Point{
		&geo.Point{0, 1},
		&geo.Point{3, 0},
		&geo.Point{3, 5},
		&geo.Point{6, 5},
		&geo.Point{8, 3},
	}
	triangles := []*Triangle{
		NewTriangleFromPoints(points[0], points[1], points[2]),
		NewTriangleFromPoints(points[1], points[2], points[3]),
		NewTriangleFromPoints(points[1], points[3], points[4]),
	}
	b := BoundFromPoints(points)
	q := NewNode(b, 1)
	q.AddTriangles(triangles)
	q.Partition(1, 3)

	for i := 0; i < 2000; i++ {
		x := 3 * rand.Float64()
		y := 5 * rand.Float64()
		p := &geo.Point{x, y}
		var correct *Triangle = nil
		for _, tr := range triangles {
			if tr.ContainsPoint(p) {
				correct = tr
				break
			}
		}
		triangle, scanned, found := q.FindTriangle(p)
		if correct != nil && !found {
			t.Error("expected triangle to be found")
		}
		if found && triangle != correct {
			t.Error("wrong triangle found")
		}
		if found && !(scanned >= 1 && scanned <= 3) {
			t.Error("expected scanned triangles to be >= 1 & <= 3, got", scanned)
		}
	}
}
