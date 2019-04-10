package graph

import (
	"fmt"
	"testing"
)

func TestMatrixGraph_Dijkstra(t *testing.T) {
	vertexes := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	matrix := [][]int{
		         /*A*//*B*//*C*//*D*//*E*//*F*//*G*/
		/*A*/ {   0,  12, INF, INF, INF,  16,  14},
		/*B*/ {  12,   0,  10, INF, INF,   7, INF},
		/*C*/ { INF,  10,   0,   3,   5,   6, INF},
		/*D*/ { INF, INF,   3,   0,   4, INF, INF},
		/*E*/ { INF, INF,   5,   4,   0,   2,   8},
		/*F*/ {  16,   7,   6, INF,   2,   0,   9},
		/*G*/ {  14, INF, INF, INF,   8,   9,   0},
	}
	g := &MatrixGraph{}
	g.InitGraph(vertexes, matrix)
	for i := 0; i < len(g.Vertex); i++ {
		fmt.Printf("%d, %d\n", i, g.Matrix[0][i])
	}
	dist := make([]int, len(vertexes))
	g.Dijkstra(0, dist)
}
