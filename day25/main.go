package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type node struct {
	x int
	y int
}

// 往右的先走，再走往南的
func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day25\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	east, south := map[node]bool{}, map[node]bool{}
	reader := bufio.NewReader(fd)
	row, col := 0, 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		col = len(str)
		for i, ch := range str {
			if ch == '>' {
				east[node{
					x: i,
					y: row,
				}] = true
			} else if ch == 'v' {
				south[node{
					x: i,
					y: row,
				}] = true
			}
		}
		row++

	}

	move, step := true, 0
	for move {
		move = false
		newEast, newSouth := map[node]bool{}, map[node]bool{}

		for item := range east {
			next := node{
				x: (item.x + 1) % col,
				y: item.y,
			}

			if east[next] || south[next] {
				newEast[item] = true
			} else {
				move = true
				newEast[next] = true
			}
		}

		east = newEast
		for item := range south {
			next := node{
				x: item.x,
				y: (item.y + 1) % row,
			}
			if east[next] || south[next] {
				newSouth[item] = true
			} else {
				move = true
				newSouth[next] = true
			}
		}

		south = newSouth
		step++
		fmt.Println(step)
	}

	fmt.Println(step)
}
