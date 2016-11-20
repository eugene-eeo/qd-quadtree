package main

import (
	"github.com/paulmach/go.geo"
	"math"
)

func BoundFromPoints(xs []*geo.Point) *geo.Bound {
	minX := xs[0].X()
	minY := xs[0].Y()
	maxX := xs[0].X()
	maxY := xs[0].Y()
	for _, p := range xs[1:] {
		minX = math.Min(minX, p.X())
		minY = math.Min(minY, p.Y())
		maxX = math.Max(maxX, p.X())
		maxY = math.Max(maxY, p.Y())
	}
	return geo.NewBound(minX, maxX, minY, maxY)
}

func midpoint(a, b float64) float64 {
	return (a + b) / 2.0
}

type Node struct {
	Triangles []*Triangle
	Children  []*Node
	Bound     *geo.Bound
	Depth     int
}

func NewNode(b *geo.Bound, depth int) *Node {
	return &Node{Bound: b, Depth: depth}
}

func (n *Node) AddTriangles(t []*Triangle) {
	for _, triangle := range t {
		if triangle.IsWithin(n.Bound) || triangle.ContainsRange(n.Bound) {
			n.Triangles = append(n.Triangles, triangle)
		}
	}
}

func (n *Node) Contains(p *geo.Point) bool {
	return n.Bound.Contains(p)
}

func (n *Node) Split() {
	if len(n.Children) > 0 {
		return
	}
	X0 := n.Bound.West()
	X1 := n.Bound.East()
	Y0 := n.Bound.South()
	Y1 := n.Bound.North()
	XM := midpoint(X0, X1)
	YM := midpoint(Y0, Y1)
	nextDepth := n.Depth + 1
	n.Children = []*Node{
		NewNode(geo.NewBound(X0, XM, Y0, YM), nextDepth),
		NewNode(geo.NewBound(XM, X1, Y0, YM), nextDepth),
		NewNode(geo.NewBound(X0, XM, YM, Y1), nextDepth),
		NewNode(geo.NewBound(XM, X1, YM, Y1), nextDepth),
	}
	for _, node := range n.Children {
		node.AddTriangles(n.Triangles)
	}
}

func (n *Node) FindNode(p *geo.Point) (*Node, bool) {
	if !n.Contains(p) {
		return nil, false
	}
	node := n
	for {
		hasChild := false
		for _, child := range node.Children {
			if child.Contains(p) {
				hasChild = true
				node = child
				break
			}
		}
		if !hasChild {
			break
		}
	}
	return node, true
}

func (n *Node) FindTriangle(p *geo.Point) (*Triangle, int, bool) {
	node, ok := n.FindNode(p)
	if !ok {
		return nil, 0, false
	}
	scanned := 0
	for _, triangle := range node.Triangles {
		scanned++
		if triangle.ContainsPoint(p) {
			return triangle, scanned, true
		}
	}
	return nil, scanned, false
}

func (n *Node) Partition(q int, d int) {
	if len(n.Triangles) > q && n.Depth < d {
		n.Split()
		for _, child := range n.Children {
			child.Partition(q, d)
		}
	}
}
