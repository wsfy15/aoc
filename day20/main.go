package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day20\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	// 0对应的位置是#！！！ 最后一个位置是.
	patternStr := "#####..#.#######.##.#.#.#.#####..######...#.##.###...#.##..#####..###.#.####..#.##.#....#...##.###..#.#......####.#...#.####.#..#.###.#.#.#.###.###..########.##.#.#....#.#####.......###.#..#.#.###.###.###.#.......#.#....#.##.###.....##...#.#.#..#.#....##..#####...##...#.##..##.##.#...###...#.##...#...#.####.###...........#...#..##...##.###.##.###.....#.##...###.....#.#.###..##..#..##..#.....##..##....#....####...#.###.#######.#.##.####.....####.#.#.##.#####...#.##.##.#.##..####.#.#.######.#..###..#.##.##..."
	// 因此对于无穷的图片，初始时，小图片外部的全是0（暗），然后每个时刻后黑暗交替（0、1、0、1...）

	pattern := convert(patternStr)
	g := [][]int{}

	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		g = append(g, convert(str))
	}

	for i := 0; i < 50; i++ {
		//printGraph(g)
		g = enhance(g, pattern, i%2)
		if i == 1 {
			fmt.Println(count(g))
		}
	}

	//printGraph(g)
	fmt.Println(count(g))
}

func printGraph(g [][]int) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func enhance(g [][]int, pattern []int, defaultVal int) [][]int {
	// 四周加上一圈defaultVal
	g = append(g, make([]int, len(g[0])+2))
	g = append([][]int{make([]int, len(g[0])+2)}, g...)
	for i := 1; i < len(g)-1; i++ {
		g[i] = append(g[i], defaultVal)
		g[i] = append([]int{defaultVal}, g[i]...)
	}

	for i := 0; i < len(g[0]); i++ {
		g[0][i] = defaultVal
		g[len(g)-1][i] = defaultVal
	}

	n, m := len(g), len(g[0])
	newG := make([][]int, n)
	for i := 0; i < n; i++ {
		newG[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			num, index := 0, 8
			for r := i - 1; r < i+2; r++ {
				for c := j - 1; c < j+2; c++ {
					if r < 0 || r >= n || c < 0 || c >= m {
						num |= defaultVal << index
						index--
						continue
					}

					num |= g[r][c] << index
					index--
				}
			}

			newG[i][j] = pattern[num]
		}
	}

	return newG
}

func convert(str string) []int {
	line := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			line[i] = 1
		}
	}

	return line
}

func count(g [][]int) int {
	res := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			res += g[i][j]
		}
	}

	return res
}
