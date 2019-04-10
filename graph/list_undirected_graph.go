package graph

import "fmt"

type (
	// 临界表中的弧
	ENode struct {
		IVertex int  	// 该边指向的顶点位置
		NextEdge *ENode	// 指向下一条弧的指针(临接顶点)
	}
	// 邻接表中表的顶点
	VNode struct {
		Data byte	// 顶点信息
		FirstEdge *ENode	// 第一条依附于该顶点的弧
	}
	// 邻接表实现的图
	ListUndirectedGraph struct {
		Vertex []*VNode
	}
)

func (g *ListUndirectedGraph) InitGraph(vs []byte, es [][]byte)  {
	vLen := len(vs)
	eLen := len(es)
	g.Vertex = make([]*VNode, vLen)
	// 初始化顶点
	for i := 0; i < vLen; i++ {
		g.Vertex[i] = &VNode{}
		g.Vertex[i].Data = vs[i]
		g.Vertex[i].FirstEdge = nil
	}
	// 初始化边
	for i := 0; i < eLen; i++ {
		// 读取边的起始顶点和结束顶点
		v  := es[i][0]
		v1 := es[i][1]
		p := g.getPosition(v)
		p1 := g.getPosition(v1)
		eNode := &ENode{
			IVertex: p1,
		}
		// 查看是否是该顶点第一条边
		if g.Vertex[p].FirstEdge == nil {
			g.Vertex[p].FirstEdge = eNode 
		//  否则加入该顶点的末尾
		}else {
			g.linkLast(g.Vertex[p].FirstEdge, eNode)
		}
		eNode1 := &ENode{
			IVertex: p,
		}
		if g.Vertex[p1].FirstEdge == nil {
			g.Vertex[p1].FirstEdge = eNode1
		}else {
			g.linkLast(g.Vertex[p1].FirstEdge, eNode1)
		}
	}
}

func (g *ListUndirectedGraph) getPosition(v byte) int {
	for i := 0; i < len(g.Vertex); i++ {
		if g.Vertex[i].Data == v {
			return i
		}
	}
	return -1
}

func (g *ListUndirectedGraph) linkLast(list *ENode, node *ENode) {
	pNode := list
	if pNode.NextEdge != nil {
		pNode = pNode.NextEdge
	}
	pNode.NextEdge = node
}

// 深度优先遍历
func (g *ListUndirectedGraph) DFS()  {
	visited := make([]bool, len(g.Vertex))
	fmt.Println("DFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			g.doDFS(i, visited)
		}
	}
	fmt.Println()
}

func (g *ListUndirectedGraph) doDFS(i int, visited []bool)  {
	node := &ENode{}
	visited[i] = true
	fmt.Printf("%c ", g.Vertex[i].Data)
	node = g.Vertex[i].FirstEdge
	for node != nil {
		if !visited[node.IVertex] {
			g.doDFS(node.IVertex, visited)
		}
		node = node.NextEdge
	}
}

// 广度优先搜索
func (g *ListUndirectedGraph) BFS()  {
	head, rear := 0, 0
	queue := make([]int, len(g.Vertex))
	visited := make([]bool, len(g.Vertex))
	fmt.Println("BFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			visited[i] = true
			fmt.Printf("%c ", g.Vertex[i].Data)
			queue[rear] = i
			rear += 1
		}
		for head != rear {
			j := queue[head]
			head += 1
			node := g.Vertex[j].FirstEdge
			for node != nil {
				k := node.IVertex
				if !visited[k] {
					visited[k] = true
					fmt.Printf("%c ", g.Vertex[k].Data)
					queue[rear] = k
					rear += 1
				}
				node = node.NextEdge
			}
		}
	}
	fmt.Println()
}
