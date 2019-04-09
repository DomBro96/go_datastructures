package graph

import (
	"testing"
)

func TestMatrixUndirectedGraphSearch(t *testing.T) {
	vertexes := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edges := [][]byte{
		{'A', 'C'},
		{'A', 'D'},
		{'A', 'F'},
		{'B', 'C'},
		{'C', 'D'},
		{'E', 'G'},
		{'F', 'G'},
	}
	mug := &MatrixUndirectedGraph{}
	mug.InitGraph(vertexes, edges)
	mug.printGraph()
	mug.DFS()
	mug.BFS()
}

