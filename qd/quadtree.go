package main

import "github.com/paulmach/go.geo"

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
	next_depth := n.Depth + 1
	n.Children = []*Node{
		NewNode(geo.NewBound(X0, XM, Y0, YM), next_depth),
		NewNode(geo.NewBound(XM, X1, Y0, YM), next_depth),
		NewNode(geo.NewBound(X0, XM, YM, Y1), next_depth),
		NewNode(geo.NewBound(XM, X1, YM, Y1), next_depth),
	}
	for _, node := range n.Children {
		node.AddTriangles(n.Triangles)
	}
}

func (node *Node) FindNode(p *geo.Point) (*Node, bool) {
	if !node.Contains(p) {
		return nil, false
	}
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

func (node *Node) FindTriangle(p *geo.Point) (*Triangle, int, bool) {
	node, ok := node.FindNode(p)
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

func (node *Node) Partition(q int, d int) {
	if len(node.Triangles) > q && node.Depth < d {
		node.Split()
		for _, child := range node.Children {
			child.Partition(q, d)
		}
	}
}
