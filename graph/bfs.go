package graph

/**
	bfs 搜索时将初始状态放入队列里，伺候队列最前端不断取出状态，把从该状态可以转移得到的状态中尚未访问过的
部分加入队列，如此反复，直到问题得到解或队列被取空。
 */

 // 迷宫最短路径问题

const (
	inf = 10000
)

// 表示点的类型
type Point struct {
	X int
	Y int
}

type PointQueue struct {
	 data []*Point
}

func (p *PointQueue) Pop() *Point  {
	if p.IsEmpty() {
		return nil
	}else {
		head := p.data[0]
		p.data = p.data[1:]
		return head
	}
}

func (p *PointQueue) Push(point *Point) {
	p.data = append(p.data, point)
}

func (p *PointQueue) IsEmpty() bool {
	return len(p.data) == 0
}

func (p *PointQueue) Front() *Point  {
	if p.IsEmpty() {
		return nil
	}else {
		head := p.data[0]
		return head
	}
}

func NewPointQueue() *PointQueue  {
	return &PointQueue{
		data: make([]*Point, 0),
	}
}

func Maze(maze [][]byte, sx, sy, ex, ey int) int {
	n := len(maze[0])
	m := len(maze)
	disMatrix := make([][]int, 0, m)
	for i := 0; i < len(disMatrix); i++ {
		disMatrix[i] = make([]int, 0, n)
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}


	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			disMatrix[i][j] = inf
		}
	}
	disMatrix[sx][sy] = 0
	queue := NewPointQueue()
	sp := &Point{sx, sy}
	queue.Push(sp)

	for !queue.IsEmpty() {
		p := queue.Pop()
		if p.X == ex && p.Y == ey {
			break
		}

		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]

			if 0 <= nx && nx < n && 0 <= ny && ny < m && maze[nx][ny] != '#' && disMatrix[nx][ny] == inf {
				np := &Point{nx, ny}
				queue.Push(np)
				disMatrix[nx][ny] = disMatrix[p.X][p.Y] + 1
			}
		}
	}
	return disMatrix[ex][ey]
}
