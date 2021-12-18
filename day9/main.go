package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day9\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	g := make([][]int, 1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		n := len(str)
		nums := make([]int, n+2)
		for i, ch := range str {
			nums[i+1] = int(ch - '0')
		}
		nums[0], nums[n+1] = 10, 10
		g = append(g, nums)
	}

	g[0] = make([]int, len(g[1]))
	g = append(g, make([]int, len(g[1])))
	n := len(g) - 1
	for i := 0; i < len(g[0]); i++ {
		g[0][i] = 10
		g[n][i] = 10
	}

	res := 0
	for i := 1; i < n; i++ {
		for j := 1; j < len(g[i])-1; j++ {
			if g[i][j] < g[i-1][j] && g[i][j] < g[i+1][j] && g[i][j] < g[i][j-1] && g[i][j] < g[i][j+1] {
				res += g[i][j] + 1
			}
		}
	}

	fmt.Println(res)
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day9\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	g := make([][]int, 1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		n := len(str)
		nums := make([]int, n+2)
		for i, ch := range str {
			nums[i+1] = int(ch - '0')
		}
		nums[0], nums[n+1] = 10, 10
		g = append(g, nums)
	}

	g[0] = make([]int, len(g[1]))
	g = append(g, make([]int, len(g[1])))
	n := len(g) - 1
	for i := 0; i < len(g[0]); i++ {
		g[0][i] = 10
		g[n][i] = 10
	}

	visited := make([][]bool, len(g))
	for i := 0; i < len(g); i++ {
		visited[i] = make([]bool, len(g[0]))
	}
	for i := 0; i < len(visited); i++ {
		visited[i][0] = true
		visited[i][len(visited[0])-1] = true
	}
	for i := 0; i < len(visited[0]); i++ {
		visited[0][i] = true
		visited[len(visited)-1][i] = true
	}

	nums := []int{}
	for i := 1; i < n; i++ {
		for j := 1; j < len(g[i])-1; j++ {
			if g[i][j] != 9 && !visited[i][j] {
				nums = append(nums, dfs(i, j, g, visited))
			}
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums[0] * nums[1] * nums[2])
}

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func dfs(row, col int, g [][]int, visited [][]bool) int {
	res := 1
	visited[row][col] = true
	for i := 0; i < 4; i++ {
		nr, nc := row+dx[i], col+dy[i]
		if g[nr][nc] != 9 && !visited[nr][nc] {
			res += dfs(nr, nc, g, visited)
		}
	}

	return res
}
