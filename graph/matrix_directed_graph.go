package graph

import "fmt"

type MatrixDirectedGraph struct {
	Vertex []byte		// 顶点集合
	Matrix [][]int		// 邻接矩阵
}

func (g *MatrixDirectedGraph) InitGraph(vs []byte, es [][]byte)  {
	 vLen := len(vs)
	 eLen := len(es)
	 g.Vertex = make([]byte, vLen)
	 for i := 0; i < vLen; i++ {
	 	g.Vertex[i] = vs[i]
	 }
	 g.Matrix = make([][]int, 0, vLen)
	 for i := 0; i < vLen; i++ {
	 	g.Matrix = append(g.Matrix, make([]int, vLen))
	 }
	 for i := 0; i < eLen; i++ {
	 	p1 := g.getPosition(es[i][0])
	 	p2 := g.getPosition(es[i][1])
	 	g.Matrix[p1][p2] = 1
	 }
}

func (g *MatrixDirectedGraph) getPosition(v byte) int  {
	for i := 0; i < len(g.Vertex); i++ {
		if g.Vertex[i] == v {
			return i
		}
	}
	return -1
}

func (g *MatrixDirectedGraph) firstVertex(v int) int  {
	if v < 0 || v >= len(g.Vertex) {
		return -1
	}
	for i := 0; i < len(g.Vertex); i++ {
		if g.Matrix[v][i] == 1 {
			return i
		}
	}
	return -1
}

func (g *MatrixDirectedGraph) nextVertex(v, w int) int  {
	if v < 0 || v >= len(g.Vertex) || w < 0 || w >= len(g.Vertex) {
		return -1
	}
	for i := w+1; i < len(g.Vertex); i++ {
		if g.Matrix[v][i] == 1 {
			return i
		}
	}
	return -1
}

// 深度优先
func (g *MatrixDirectedGraph) DFS()  {
	visited := make([]bool, len(g.Vertex))
	fmt.Println("DFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			g.doDFS(visited, i)
		}
	}
	fmt.Println()
}

func (g *MatrixDirectedGraph) doDFS(visited []bool, v int)  {
	visited[v] = true
	fmt.Printf("%c ", g.Vertex[v])
	for w := g.firstVertex(v); w >= 0; w = g.nextVertex(v, w) {
		if !visited[w] {
			g.doDFS(visited, w)
		}
	}
}

// 广度优先
func (g *MatrixDirectedGraph) BFS()  {
	head, rear := 0, 0
	queue := make([]int, len(g.Vertex))
	visited := make([]bool, len(g.Vertex))
	fmt.Println("BFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			visited[i] = true
			fmt.Printf("%c ", g.Vertex[i])
			queue[rear] = i
			rear += 1
		}
		// 层次遍历
		for head != rear {
			j := queue[head]
			head += 1
			for k := g.firstVertex(j); k >= 0; k = g.nextVertex(j, k) {
				if !visited[k] {
					visited[k] = true
					fmt.Printf("%c ", g.Vertex[k])
					queue[rear] = k
					rear += 1
				}
			}
		}
	}

}

