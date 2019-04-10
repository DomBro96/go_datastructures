package graph

import "testing"

func TestListDirectedGraph_BFS(t *testing.T) {
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
	g := &ListDirectedGraph{}
	g.InitGraph(vertex, edges)
	g.BFS()
}

func TestListDirectedGraph_DFS(t *testing.T) {
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
	g := &ListDirectedGraph{}
	g.InitGraph(vertex, edges)
	g.DFS()
}

func TestListUndirectedGraph_BFS(t *testing.T) {
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
	g := &ListUndirectedGraph{}
	g.InitGraph(vertexes, edges)
	g.BFS()
}

func TestListUndirectedGraph_DFS(t *testing.T) {
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
	g := &ListUndirectedGraph{}
	g.InitGraph(vertexes, edges)
	g.DFS()
}

