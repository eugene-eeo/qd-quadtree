package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"os"
)

func countNodes(node *Node, count int) int {
	for _, child := range node.Children {
		count = countNodes(child, count)
		count++
	}
	return count
}

func main() {
	mesh := []*Triangle{}
	source := bufio.NewScanner(os.Stdin)
	points := [3][2]float64{}
	for source.Scan() {
		err := json.Unmarshal(source.Bytes(), &points)
		if err != nil {
			fmt.Println(err)
			return
		}
		mesh = append(mesh, NewTriangleFromPoints(
			NewPoint(points[0][0], points[0][1]),
			NewPoint(points[1][0], points[1][1]),
			NewPoint(points[2][0], points[2][1]),
		))
	}
	r := NewRangeFromMesh(mesh)
	for _, q := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
		quadtree := NewNode(r, 1)
		quadtree.Triangles = mesh
		quadtree.Partition(q, 10)
		fmt.Println(countNodes(quadtree, 0))
	}
}
