package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day15\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := [][]int{}
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		nums := []int{}
		for _, ch := range []byte(str) {
			nums = append(nums, int(ch-'0'))
		}
		g = append(g, nums)
	}

	g = expand(5, g) // for part 2

	n := len(g)
	graph := NewGraph(n * n)
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sid := i*n + j
			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				graph.AddEdge(sid, ni*n+nj, g[ni][nj])
			}
		}
	}

	graph.Dijkstra(0, n*n-1)
}

func expand(k int, g [][]int) [][]int {
	n := len(g)
	res := make([][]int, n*k)
	for i := 0; i < n*k; i++ {
		res[i] = make([]int, n*k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for x := 0; x < k; x++ {
				for y := 0; y < k; y++ {
					res[i+x*n][j+y*n] = (g[i][j]+x+y-1)%9 + 1
				}
			}
		}
	}

	return res
}

type Edge struct {
	Sid    int // 起始顶点编号
	Eid    int // 终止顶点编号
	Weight int // 权重
}

func NewEdge(sid, eid, weight int) Edge {
	return Edge{
		Sid:    sid,
		Eid:    eid,
		Weight: weight,
	}
}

// 用于Dijkstra算法中
type Vertex struct {
	Id   int // 顶点编号
	Dist int // 从起点到该点距离
}

func NewVertex(id, dist int) Vertex {
	return Vertex{
		Id:   id,
		Dist: dist,
	}
}

type Graph struct {
	Count int      // 顶点个数
	Adj   [][]Edge // 邻接表
}

func NewGraph(count int) *Graph {
	return &Graph{
		Count: count,
		Adj:   make([][]Edge, count),
	}
}

func (this *Graph) AddEdge(a, b, weight int) { // 边a<->b
	if a >= this.Count || b >= this.Count {
		return
	}
	this.Adj[a] = append(this.Adj[a], NewEdge(a, b, weight))
}

// 从顶点a到b的最短路径
func (this Graph) Dijkstra(a, b int) {
	predecessor := make([]int, this.Count)
	vertexs := make([]Vertex, this.Count) // 记录从起始顶点到每个顶点的距离
	for i := range vertexs {
		vertexs[i] = NewVertex(i, math.MaxInt32) // 初始化距离无限大
	}

	queue := make([]Vertex, 1) // 优先级队列，最小堆
	vertexs[a].Dist = 0
	queue = append(queue, vertexs[a])
	inqueue := make([]bool, this.Count) // 记录是否已经在最小堆中
	inqueue[a] = true

	for len(queue) > 1 { // 下标0处不存储
		minVertex := queue[1] // 取 距离已选顶点 最近的顶点
		if minVertex.Id == b {
			fmt.Println("dist: ", minVertex.Dist)
			break
		}
		// 删除元素后需要堆化
		queue[1] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		heapify2(queue)

		// 取出每一条与当前顶点相连的边，计算使用当前顶点作为中继的距离
		for i := 0; i < len(this.Adj[minVertex.Id]); i++ {
			edge := this.Adj[minVertex.Id][i]
			oldDist := vertexs[edge.Eid].Dist
			if minVertex.Dist+edge.Weight < oldDist {
				vertexs[edge.Eid].Dist = minVertex.Dist + edge.Weight
				predecessor[edge.Eid] = minVertex.Id
				if inqueue[edge.Eid] == false {
					//从下往上堆化
					queue = append(queue, vertexs[edge.Eid])
					heapify(queue)
					inqueue[edge.Eid] = true
				} else {
					update(queue, edge.Eid, vertexs[edge.Eid].Dist)
				}
			}
		}
	}

	//print(a, b, predecessor)
}

// 从下往上堆化
func heapify(queue []Vertex) {
	i := len(queue) - 1
	for i > 1 && queue[i].Dist < queue[i/2].Dist {
		queue[i], queue[i/2] = queue[i/2], queue[i]
		i /= 2
	}
}

// 从上往下堆化
func heapify2(queue []Vertex) {
	i := 1
	size := len(queue)
	for {
		min := i
		if 2*i < size && queue[i].Dist > queue[i*2].Dist {
			min = i * 2
		}
		if 2*i+1 < size && queue[min].Dist > queue[i*2+1].Dist {
			min = i*2 + 1
		}

		if min == i {
			return
		}
		queue[i], queue[min] = queue[min], queue[i]
		i = min
	}
}

// 更新Id为id的顶点，因为每次更新后，该顶点的dist肯定是减少的，所以从下往上堆化即可
func update(queue []Vertex, id, newDist int) {
	var pos int
	for i := 1; i < len(queue); i++ {
		if queue[i].Id == id {
			pos = i
			break
		}
	}

	queue[pos].Dist = newDist
	for pos > 1 && queue[pos].Dist < queue[pos/2].Dist {
		queue[pos], queue[pos/2] = queue[pos/2], queue[pos]
		pos /= 2
	}
}

func print(a, b int, predecessor []int) {
	if b == a {
		fmt.Print(a)
	} else {
		print(a, predecessor[b], predecessor)
		fmt.Print("->", b)
	}
}
