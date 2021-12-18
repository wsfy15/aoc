package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day12\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := [][]int{}
	isSmall := []bool{}
	id := map[string]int{}
	count, from, to := 0, 0, 0
	var ok bool
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, "-")
		if from, ok = id[items[0]]; !ok {
			id[items[0]] = count
			from = count
			count++
			g = append(g, []int{})
			isSmall = append(isSmall, 'a' <= items[0][0] && items[0][0] <= 'z')
		}
		if to, ok = id[items[1]]; !ok {
			id[items[1]] = count
			to = count
			count++
			g = append(g, []int{})
			isSmall = append(isSmall, 'a' <= items[1][0] && items[1][0] <= 'z')
		}

		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	visited := make([]bool, len(g))
	res := dfs(id["start"], id["end"], g, isSmall, visited)

	fmt.Println(res)
}

func dfs(from int, to int, g [][]int, isSmall, visited []bool) int {
	if from == to {
		return 1
	}

	res := 0
	visited[from] = true
	for _, neighbor := range g[from] {
		if !visited[neighbor] || !isSmall[neighbor] {
			res += dfs(neighbor, to, g, isSmall, visited)
		}
	}
	visited[from] = false

	return res
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day12\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := [][]int{}
	isSmall := []bool{}
	id := map[string]int{}
	count, from, to := 0, 0, 0
	var ok bool
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, "-")
		if from, ok = id[items[0]]; !ok {
			id[items[0]] = count
			from = count
			count++
			g = append(g, []int{})
			isSmall = append(isSmall, 'a' <= items[0][0] && items[0][0] <= 'z')
		}
		if to, ok = id[items[1]]; !ok {
			id[items[1]] = count
			to = count
			count++
			g = append(g, []int{})
			isSmall = append(isSmall, 'a' <= items[1][0] && items[1][0] <= 'z')
		}

		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	visited := make([]bool, len(g))
	res := dfs2(id["start"], id["start"], id["end"], -1, false, g, isSmall, visited)

	fmt.Println(res)
}

func dfs2(start, from, to, twicePos int, useTwice bool, g [][]int, isSmall, visited []bool) int {
	if from == to {
		return 1
	}

	res := 0
	visited[from] = true
	for _, neighbor := range g[from] {
		if neighbor != start {
			if !visited[neighbor] || !isSmall[neighbor] {
				res += dfs2(start, neighbor, to, twicePos, useTwice, g, isSmall, visited)
			} else if !useTwice {
				res += dfs2(start, neighbor, to, neighbor, true, g, isSmall, visited)
			}
		}
	}

	// Important!!!!  没有该判断会导致提前将visited[from]置为false，导致爆栈
	if useTwice && from == twicePos {
		return res
	}

	visited[from] = false
	return res
}
