package main

import (
	"fmt"
	"strings"
)

func printTree(qd *Node) {
	fmt.Println(strings.Repeat(" ", qd.Depth-1), qd.Depth, qd.Bound, len(qd.Triangles))
	for _, n := range qd.Children {
		printTree(n)
	}
}

func main() {
	mesh := []*Triangle{
		NewTriangleFromPoints(
			NewPoint(0, 1),
			NewPoint(3, 0),
			NewPoint(3, 5),
		),
		NewTriangleFromPoints(
			NewPoint(6, 5),
			NewPoint(3, 0),
			NewPoint(3, 5),
		),
		NewTriangleFromPoints(
			NewPoint(6, 5),
			NewPoint(3, 0),
			NewPoint(8, 3),
		),
	}
	bigBound := NewRange(0, 8, 0, 5)
	quadtree := NewNode(bigBound, 1)
	quadtree.AddTriangles(mesh)
	quadtree.Partition(1, 10)
	printTree(quadtree)
	t, ok := quadtree.FindTriangle(NewPoint(4, 4))
	fmt.Println(ok, t == mesh[1])
}
