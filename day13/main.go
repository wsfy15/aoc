package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1()
	//p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day13\\point.data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := make([][]bool, 1500)
	for i := 0; i < 1500; i++ {
		g[i] = make([]bool, 1500)
	}

	reader := bufio.NewReader(fd)
	maxX, maxY := 0, 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, ",")
		x, _ := strconv.Atoi(items[0])
		y, _ := strconv.Atoi(items[1])
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		g[y][x] = true
	}

	process(g, maxX, maxY)
}

func process(g [][]bool, maxX, maxY int) {
	fd, err := os.OpenFile("E:\\test\\aoc\\day13\\instruction.data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		items := strings.Split(str[:len(str)-1], "=")
		num, _ := strconv.Atoi(items[1])
		if items[0] == "x" { // vertical
			for i := 0; i <= maxY; i++ {
				for j := num + 1; j <= maxX; j++ {
					if g[i][j] && num-(j-num) >= 0 {
						g[i][num-(j-num)] = true
					}
				}
			}
			maxX = num - 1
		} else { // horizontal
			for i := num + 1; i <= maxY; i++ {
				for j := 0; j <= maxX; j++ {
					if g[i][j] && num-(i-num) >= 0 {
						g[num-(i-num)][j] = true
					}
				}
			}
			maxY = num - 1
		}

		fmt.Println(count(g, maxX, maxY))
	}

	printGraph(g, maxX, maxY)
}

func count(g [][]bool, maxX, maxY int) int {
	res := 0
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if g[i][j] {
				res++
			}
		}
	}

	return res
}

func printGraph(g [][]bool, x int, y int) {
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if g[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
