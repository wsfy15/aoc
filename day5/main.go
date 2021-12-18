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
	fd, err := os.OpenFile("E:\\test\\aoc\\day5\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		g[i] = make([]int, 1000)
	}

	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " -> ")
		left := strings.Split(items[0], ",")
		right := strings.Split(items[1], ",")
		x1, _ := strconv.Atoi(left[0])
		y1, _ := strconv.Atoi(left[1])
		x2, _ := strconv.Atoi(right[0])
		y2, _ := strconv.Atoi(right[1])
		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				g[x1][i]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				g[i][y1]++
			}
		} else {
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			delta := (x2 - x1) / (y2 - y1)
			x, y := x1, y1
			for i := 0; i <= x2-x1; i++ {
				g[x][y]++
				x++
				y += delta
			}
		}
	}

	res := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if g[i][j] > 1 {
				res++
			}
		}
	}
	fmt.Println(res)
}
