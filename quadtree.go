package main

type Node struct {
	Triangles []*Triangle
	Children  []*Node
	Bound     *Range
	Depth     int
}

func NewNode(r *Range, depth int) *Node {
	return &Node{Bound: r, Depth: depth}
}

func (n *Node) AddTriangles(t []*Triangle) {
	for _, triangle := range t {
		if triangle.IsWithin(n.Bound) {
			n.Triangles = append(n.Triangles, triangle)
		}
	}
}

func (n *Node) Contains(p *Point) bool {
	return n.Bound.Contains(p)
}

func (n *Node) Split() {
	X0 := n.Bound.X0
	X1 := n.Bound.X1
	Y0 := n.Bound.Y0
	Y1 := n.Bound.Y1
	XM := midpoint(X0, X1)
	YM := midpoint(Y0, Y1)
	next_depth := n.Depth + 1
	n.Children = make([]*Node, 4)
	n.Children[0] = NewNode(NewRange(X0, XM, Y0, YM), next_depth)
	n.Children[1] = NewNode(NewRange(XM, X1, Y0, YM), next_depth)
	n.Children[3] = NewNode(NewRange(X0, XM, YM, Y1), next_depth)
	n.Children[2] = NewNode(NewRange(XM, X1, YM, Y1), next_depth)
	for _, node := range n.Children {
		node.AddTriangles(n.Triangles)
	}
}

func (node *Node) FindNode(p *Point) (*Node, bool) {
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

func (node *Node) Partition(q int, d int) {
	if len(node.Triangles) > q && node.Depth < d {
		node.Split()
		for _, child := range node.Children {
			child.Partition(q, d)
		}
	}
}
