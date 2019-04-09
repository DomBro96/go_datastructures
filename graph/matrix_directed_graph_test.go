package graph

import "testing"

func TestMatrixDirectedGraph_DFS(t *testing.T) {
	vertex := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edges := [][]byte{
		{'A', 'B'},
		{'B', 'C'},
		{'B', 'E'},
		{'B', 'F'},
		{'C', 'E'},
		{'D', 'C'},
		{'E', 'B'},
		{'E', 'D'},
		{'F', 'G'},
	}
	graph := &MatrixDirectedGraph{
		Vertex: nil,
		Matrix: nil,
	}
	graph.InitGraph(vertex, edges)
	graph.DFS()

}

func TestMatrixDirectedGraph_BFS(t *testing.T) {
	vertex := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edges := [][]byte{
		{'A', 'B'},
		{'B', 'C'},
		{'B', 'E'},
		{'B', 'F'},
		{'C', 'E'},
		{'D', 'C'},
		{'E', 'B'},
		{'E', 'D'},
		{'F', 'G'},
	}
	graph := &MatrixDirectedGraph{
		Vertex: nil,
		Matrix: nil,
	}
	graph.InitGraph(vertex, edges)
	graph.BFS()
}