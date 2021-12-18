package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day11\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := [][]int{}

	res := 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		nums := make([]int, len(str))
		for i, ch := range str {
			nums[i] = int(ch - '0')
		}
		g = append(g, nums)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		res += process(g)
	}

	fmt.Println(res)
}

var dx = []int{1, 1, 1, -1, -1, -1, 0, 0}
var dy = []int{1, -1, 0, 1, -1, 0, 1, -1}

func process(g [][]int) int {
	n, m := len(g), len(g[0])
	queue := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i][j]++
			if g[i][j] > 9 {
				queue = append(queue, i*m+j)
			}
		}
	}

	count := 0
	for len(queue) > 0 {
		count += len(queue)
		newQueue := []int{}
		for _, num := range queue {
			row, col := num/m, num%m
			g[row][col] = -1000000
			for i := 0; i < 8; i++ {
				nr, nc := row+dx[i], col+dy[i]
				if nr < 0 || nr >= n || nc < 0 || nc >= m || g[nr][nc] > 9 {
					continue
				}

				g[nr][nc]++
				if g[nr][nc] > 9 {
					newQueue = append(newQueue, nr*m+nc)
				}
			}
		}

		queue = newQueue
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] < 0 {
				g[i][j] = 0
			}
		}
	}
	return count
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day11\\data", os.O_RDONLY, 0751)
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
		nums := make([]int, len(str))
		for i, ch := range str {
			nums[i] = int(ch - '0')
		}
		g = append(g, nums)
	}

	step, total := 0, len(g)*len(g[0])
	for {
		step++
		if total == process(g) {
			break
		}
	}

	fmt.Println(step)
}
