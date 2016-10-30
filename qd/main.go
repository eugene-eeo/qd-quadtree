package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/paulmach/go.geo"
)

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
