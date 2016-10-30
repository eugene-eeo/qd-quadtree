package main

import (
	"encoding/json"
	"fmt"
	"github.com/paulmach/go.geo"
	"os"
	"strings"
)

func printNode(node *Node) {
	fmt.Println(strings.Repeat(" ", node.Depth-1), node)
	for _, child := range node.Children {
		printNode(child)
	}
}

func countNodes(node *Node, count int) int {
	count++
	for _, child := range node.Children {
		count = countNodes(child, count)
		count++
	}
	return count
}

type JSONData struct {
	Points    []*geo.Point
	Simplices [][3]int
}

func main() {
	ctx := new(JSONData)
	mesh := []*Triangle{}
	source := json.NewDecoder(os.Stdin)
	if err := source.Decode(&ctx); err != nil {
		panic(err)
	}

	for _, idxs := range ctx.Simplices {
		mesh = append(mesh, NewTriangleFromPoints(
			ctx.Points[idxs[0]],
			ctx.Points[idxs[1]],
			ctx.Points[idxs[2]],
		))
	}

	b := BoundFromPoints(ctx.Points)
	for _, d := range []int{5, 10, 15} {
		for _, q := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
			quadtree := NewNode(b, 1)
			quadtree.Triangles = mesh
			quadtree.Partition(q, d)
			fmt.Printf("q=%d d=%d nodes=%d\n", q, d, countNodes(quadtree, 0))
		}
	}
}

/*
func main() {
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
	t, ok := q.FindTriangle(&geo.Point{2,2})
	fmt.Println(t == triangles[0])
	fmt.Println(ok)
}
*/
