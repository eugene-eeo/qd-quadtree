package main

import (
	"encoding/json"
	"fmt"
	"github.com/paulmach/go.geo"
	"math/rand"
	"os"
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
	Points    []*geo.Point `json:"points"`
	Simplices [][3]int     `json:"simplices"`
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
	x0 := b.West()
	dx := b.East() - x0
	y0 := b.South()
	dy := b.North() - y0

	for _, q := range []int{256, 128, 64, 32, 16, 8, 4, 2, 1} {
		quadtree := NewNode(b, 1)
		quadtree.Triangles = mesh
		for _, d := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} {
			quadtree.Partition(q, d)
			scanned := 0
			for i := 0; i < 2000; i++ {
				p := &geo.Point{
					dx*rand.Float64() + x0,
					dy*rand.Float64() + y0,
				}
				_, s, _ := quadtree.FindTriangle(p)
				scanned += s
			}
			fmt.Printf("q=%d d=%d nodes=%d scanned=%d\n", q, d, countNodes(quadtree, 0), scanned)
		}
	}
}
