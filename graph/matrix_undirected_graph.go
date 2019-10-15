package graph

import (
	"fmt"
)

type MatrixUndirectedGraph struct {
	Vertex []byte		// 顶点集合
	Matrix [][]int		// 邻接矩阵
}

/**
	初始化图:
	1. v 顶点数组
	2. m 边数组 m[i][0] m[i][1] 表示 两定点有边
 */
func (g *MatrixUndirectedGraph) InitGraph(vs []byte, es [][]byte) {
	// 顶点数目
	vLen := len(vs)
	// 边数目
	eLen := len(es)
	// 初始化顶点
	g.Vertex = make([]byte, vLen)
	for i := 0; i < vLen; i++ {
		g.Vertex[i] = vs[i]
	}

	// 初始化边
	g.Matrix = make([][]int, 0, vLen)
	for i := 0; i < vLen; i++ {
		g.Matrix = append(g.Matrix, make([]int, vLen))
	}
	// 初始化邻接矩阵
	for j := 0; j < eLen; j++ {
		p  := g.getPosition(es[j][0])
		p1 := g.getPosition(es[j][1])
		g.Matrix[p][p1] = 1
		g.Matrix[p1][p] = 1
	}
}

// 返回 ch 位置
func (g *MatrixUndirectedGraph) getPosition(ch byte) int {
	for i := 0; i < len(g.Vertex); i++ {
		if g.Vertex[i] == ch {
			return i
		}
	}
	return -1
}

// 返回 顶点v 第一个临界点的位置索引,
func (g *MatrixUndirectedGraph) firstVertex(v int) int  {
	if v < 0 || v > len(g.Matrix) - 1 {
		return -1
	}
	for i := 0; i < len(g.Matrix); i++ {
		if g.Matrix[v][i] == 1 {
			return i
		}
 	}
	return -1
}

// 返回顶点 v 相对 w 的下一个临接顶点索引，失败则返回 -1
func (g *MatrixUndirectedGraph) nextVertex(v, w int) int  {
	if v < 0 || v > len(g.Matrix) - 1 || w < 0  || w > len(g.Matrix) - 1 {
		return -1
	}
	for i := w+1; i < len(g.Matrix); i++ {
		if g.Matrix[v][i] == 1 {
			return i
		}
	}
	return -1
}

func (g *MatrixUndirectedGraph) doDFS(i int, visited []bool)  {
	visited[i] = true
	fmt.Printf("%c ", g.Vertex[i])
	// 递归深度遍历，沿着某一未访问过的点向下遍历
	for w := g.firstVertex(i); w >= 0; w = g.nextVertex(i, w) {
		if ! visited[w] {
			g.doDFS(w, visited)
		}
	}
}
// 深度优先遍历
func (g *MatrixUndirectedGraph) DFS()  {
	visited := make([]bool, len(g.Vertex))
	fmt.Println("DFS: ")
	// 依次对每个节点进行深度访问
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			g.doDFS(i, visited)
		}
	}
	fmt.Println()
}

// 广度优先搜索, 类似于树的层次遍历，都需要用到队列
func (g *MatrixUndirectedGraph) BFS()  {
	head, rear := 0, 0
	queue := make([]int, len(g.Vertex))
	visited := make([]bool, len(g.Vertex))
	fmt.Println("BFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			visited[i] = true
			fmt.Printf("%c ", g.Vertex[i])
			queue[rear] = i			// 入队列
			rear+=1
		}
		for head != rear {
			j := queue[head]		// 出队列
			head+=1
			for k := g.firstVertex(j); k >= 0; k = g.nextVertex(j, k) {	// k 是未访问的临接顶点
				if !visited[k] {
					visited[k] = true
					fmt.Printf("%c ", g.Vertex[k])
					queue[rear] = k
					rear+=1
				}
			}
		}
	}
	fmt.Println()
}

func (g *MatrixUndirectedGraph) printGraph()  {
	fmt.Println("print graph")
	for i := 0; i < len(g.Vertex); i++ {
		for j := 0; j < len(g.Vertex); j++ {
			fmt.Printf("%d ", g.Matrix[i][j])
		}
		fmt.Println()
	}
}


// 双色问题, 用来判断图是否可以使用两种颜色对顶点进行染色。使用深度优先搜索 DFS
func (g *MatrixUndirectedGraph) Solve() bool {
	color := make([]int, len(g.Vertex))
	for i := 0; i < len(g.Vertex); i++ {
		if color[i] == 0 {
			return g.doSolve(i, 1, color)
		}
	}
	return false
}

func (g *MatrixUndirectedGraph) doSolve(v, c int, color []int) bool  {
	// 当前顶点颜色
	color[v] = c
	for w := g.firstVertex(v); w >= 0; w = g.nextVertex(v, w) {
		// 邻接定点颜色和定点颜色相同
		if color[w] == c {
			return false
		// 邻接顶点未被染色且对其染不同颜色失败
		}else if color[w] == 0 && !g.doSolve(w, -c, color){
			return false
		}
	}
	return true
}