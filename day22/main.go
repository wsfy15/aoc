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
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day22\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := [102][102][102]int{}
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " ")
		value := 0
		if items[0] == "on" {
			value = 1
		}

		items = strings.Split(items[1], ",")
		xstart, xend := getStartAndEnd(items[0][2:])
		ystart, yend := getStartAndEnd(items[1][2:])
		zstart, zend := getStartAndEnd(items[2][2:])

		xstart, ystart, zstart = max(-50, xstart)+50, max(-50, ystart)+50, max(-50, zstart)+50
		xend, yend, zend = min(50, xend)+50, min(50, yend)+50, min(50, zend)+50
		for i := xstart; i <= xend; i++ {
			for j := ystart; j <= yend; j++ {
				for k := zstart; k <= zend; k++ {
					g[i][j][k] = value
				}
			}
		}
	}

	res := 0
	for i := 0; i < 102; i++ {
		for j := 0; j < 102; j++ {
			for k := 0; k < 102; k++ {
				res += g[i][j][k]
			}
		}
	}

	fmt.Println(res)
}

type cube struct {
	xstart, xend int
	ystart, yend int
	zstart, zend int
	value        int
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day22\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	cubes := []cube{}
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " ")
		cur := cube{}
		if items[0] == "on" {
			cur.value = 1
		}

		items = strings.Split(items[1], ",")
		xstart, xend := getStartAndEnd(items[0][2:])
		ystart, yend := getStartAndEnd(items[1][2:])
		zstart, zend := getStartAndEnd(items[2][2:])

		// 左闭右开，方便计算体积
		xend++
		yend++
		zend++

		cur.xstart = xstart
		cur.xend = xend
		cur.ystart = ystart
		cur.yend = yend
		cur.zstart = zstart
		cur.zend = zend

		newCubes := []cube{}
		for _, c := range cubes {
			// 与现有cube比较，是否存在交集
			// 存在的话，就将现有的cube 非交集部分切割出来
			if (xstart < c.xend && xend > c.xstart) && (ystart < c.yend && yend > c.ystart) &&
				(zstart < c.zend && zend > c.zstart) {
				if xstart > c.xstart {
					newCubes = append(newCubes, cube{
						xstart: c.xstart,
						xend:   xstart,
						ystart: c.ystart,
						yend:   c.yend,
						zstart: c.zstart,
						zend:   c.zend,
						value:  c.value,
					})
					c.xstart = xstart
				}

				if xend < c.xend {
					newCubes = append(newCubes, cube{
						xstart: xend,
						xend:   c.xend,
						ystart: c.ystart,
						yend:   c.yend,
						zstart: c.zstart,
						zend:   c.zend,
						value:  c.value,
					})
					c.xend = xend
				}

				if ystart > c.ystart {
					newCubes = append(newCubes, cube{
						xstart: c.xstart,
						xend:   c.xend,
						ystart: c.ystart,
						yend:   ystart,
						zstart: c.zstart,
						zend:   c.zend,
						value:  c.value,
					})
					c.ystart = ystart
				}

				if yend < c.yend {
					newCubes = append(newCubes, cube{
						xstart: c.xstart,
						xend:   c.xend,
						ystart: yend,
						yend:   c.yend,
						zstart: c.zstart,
						zend:   c.zend,
						value:  c.value,
					})
					c.yend = yend
				}

				if zstart > c.zstart {
					newCubes = append(newCubes, cube{
						xstart: c.xstart,
						xend:   c.xend,
						ystart: c.ystart,
						yend:   c.yend,
						zstart: c.zstart,
						zend:   zstart,
						value:  c.value,
					})
					c.zstart = zstart
				}

				if zend < c.zend {
					newCubes = append(newCubes, cube{
						xstart: c.xstart,
						xend:   c.xend,
						ystart: c.ystart,
						yend:   c.yend,
						zstart: zend,
						zend:   c.zend,
						value:  c.value,
					})
					c.zend = zend
				}
			} else {
				newCubes = append(newCubes, c)
			}
		}

		newCubes = append(newCubes, cur)
		cubes = newCubes
	}

	res := 0
	for _, c := range cubes {
		if c.value == 1 {
			res += (c.xend - c.xstart) * (c.yend - c.ystart) * (c.zend - c.zstart)
		}
	}

	fmt.Println(res)
}

func getStartAndEnd(str string) (int, int) {
	items := strings.Split(str, "..")
	start, _ := strconv.Atoi(items[0])
	end, _ := strconv.Atoi(items[1])
	return start, end
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
