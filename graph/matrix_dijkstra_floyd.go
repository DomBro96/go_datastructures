package graph

import (
	"fmt"
	"math"
)

// 利用临界矩阵实现的 Floyd 算法

const (
	INF = math.MaxInt64
)


type MatrixGraph struct {
	Vertex []byte		// 顶点集合
	Matrix [][]int		// 邻接矩阵
	EdgeNum int			// 边的数量
}

func (g *MatrixGraph) InitGraph(vs []byte, m [][]int)  {
	vLen := len(vs)
	g.Vertex = make([]byte, vLen)
	// 初始化顶点
	for i := 0; i < vLen; i++ {
		g.Vertex[i] = vs[i]
	}
	g.Matrix = make([][]int, 0, vLen)
	for i := 0; i < vLen; i++ {
		g.Matrix = append(g.Matrix, make([]int, vLen))
	}
	// 初始化边
	for i := 0; i < vLen; i++ {
		for j := 0; j < vLen; j++ {
			g.Matrix[i][j] = m[i][j]
		}
	}
	for i := 0; i < vLen; i++ {
		for j := i+1; j < vLen; j++ {
			if g.Matrix[i][j] != INF {
				// 边的数量
				g.EdgeNum++
			}
		}
	}
}

func (g *MatrixGraph) FirstVertex(v int) int {
	if v < 0 || v > len(g.Vertex) - 1 {
		return -1
	}
	for i := 0; i < len(g.Vertex); i++ {
		if g.Matrix[v][i] != 0 && g.Matrix[v][i] != INF {
			return i
		}
	}
	return -1
}


func (g *MatrixGraph) NextVertex(v, w int) int {
	if v < 0 || v > len(g.Vertex) - 1 ||  w < 0 || w > len(g.Vertex) - 1{
		return -1
	}
	for i := w+1; i < len(g.Vertex); i++ {
		if g.Matrix[v][i] != 0 && g.Matrix[v][i] != INF {
			return i
		}
	}
	return -1
}

// dist[i] 源点 v 到 i 的最段距离
func (g *MatrixGraph) Dijkstra(v int, dist []int) {
	flag := make([]bool, len(g.Vertex))
	// 到源点最近的距离, 初始化为源点
	nearest := v
	// flag 表示该点有没有被当过中间点
	flag[v] = true
	// 距离初始化
	for i := 0; i < len(g.Vertex); i++ {
		dist[i] = g.Matrix[v][i]
	}
	// 对除源点外的顶点求最最短距离
	for i := 1; i < len(g.Vertex); i++ {
		minDis := INF
		// 寻找距离源点最近的且没有访问过的中间点
		for j := 0; j < len(g.Vertex); j++ {
			if !flag[j] && minDis > dist[j] {
				minDis = dist[j]
				nearest = j				// 找到离源点最近的点
			}
		}
		flag[nearest] = true
		for j := 0; j < len(g.Vertex); j++ {
			temp := 0
			if g.Matrix[nearest][j] == INF {
				temp = INF
			}else {
				temp = minDis + g.Matrix[nearest][j]
			}
			if !flag[j] && temp < dist[j] {
				dist[j] = temp
			}
 		}
	}

	fmt.Printf("dijkstra %c : \n", g.Vertex[v])
	for i := 0; i < len(g.Vertex); i++ {
		fmt.Printf("shortest distance(%c, %c)=%d ", g.Vertex[v], g.Vertex[i], dist[i])
	}
}

func min(a, b int) int  {
	if a <= b {
		return a
	}else {
		return b
	}
}



