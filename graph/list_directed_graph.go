package graph

import "fmt"

type (
	DENode struct {
		VIndex int
		NextEdge *DENode
	}
	DVNode struct {
		Data byte
		FirstEdge *DENode
	}
	ListDirectedGraph struct {
		Vertex []*DVNode
	}
)

func (g *ListDirectedGraph) InitGraph(vs []byte, es [][]byte)  {
	vLen := len(vs)
	eLen := len(es)
	vertex := make([]*DVNode, vLen)
	// 初始化点
	for i := 0; i < vLen; i++ {
		vertex[i] = &DVNode{}
		vertex[i].Data = vs[i]
		vertex[i].FirstEdge = nil
	}
	g.Vertex = vertex
	// 初始化边
	for i := 0; i < eLen; i++ {
		v := es[i][0]
		v1 := es[i][1]
		p := g.getPosition(v)
		p1 := g.getPosition(v1)
		node := &DENode{
			VIndex: p1,
		}
		if g.Vertex[p].FirstEdge == nil {
			g.Vertex[p].FirstEdge = node
		}else {
			g.linkLast(g.Vertex[p].FirstEdge, node)
		}
	}
}

func (g *ListDirectedGraph) getPosition(v byte) int {
	for i := 0; i < len(g.Vertex); i++ {
		if g.Vertex[i].Data == v {
			return i
		}
	}
	return -1
}

func (g *ListDirectedGraph) linkLast(list *DENode, node *DENode) {
	pNode := list
	if pNode.NextEdge != nil {
		pNode = pNode.NextEdge
	}
	pNode.NextEdge = node
}

func (g *ListDirectedGraph) DFS()  {
	visited := make([]bool, len(g.Vertex))
	fmt.Println("DFS: ")
	for i := 0; i < len(g.Vertex); i++ {
		if !visited[i] {
			g.doDFS(i, visited)
		}
	}
	fmt.Println()
}

func (g *ListDirectedGraph) doDFS(i int, visited []bool)  {
	node := &DENode{}
	visited[i] = true
	fmt.Printf("%c ", g.Vertex[i].Data)
	node = g.Vertex[i].FirstEdge
	for node != nil {
		if !visited[node.VIndex] {
			g.doDFS(node.VIndex, visited)
		}
		node = node.NextEdge
	}
}

func (g *ListDirectedGraph) BFS()  {
	head, rear := 0, 0
	queue := make([]int, len(g.Vertex))
	visited := make([]bool, len(g.Vertex))
	fmt.Println("BSF: ")
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
				k := node.VIndex
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

}