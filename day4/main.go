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
	fd, err := os.Open("E:\\test\\aoc\\day4\\data")
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	items := strings.Split(str, ",")
	nums := make([]int, len(items))
	for i, item := range items {
		num, _ := strconv.Atoi(item)
		nums[i] = num
	}

	boards := [][][]int{}
	index, row := -1, 5
	for {
		sb, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil || len(sb) != 14 {
			continue
		}

		if row == 5 {
			row = 0
			index++
			// 6x6 矩阵 最后一行和最后一列表示该行/列 marked元素数量
			board := make([][]int, 6)
			for i := 0; i < 6; i++ {
				board[i] = make([]int, 6)
			}
			boards = append(boards, board)
		}

		for i := 0; i < 15; i += 3 {
			num := int(sb[i+1] - '0')
			if sb[i] != ' ' {
				num += 10 * int(sb[i]-'0')
			}
			boards[index][row][i/3] = num
		}

		row++
	}

	fmt.Println("finish constructing boards")

	for _, num := range nums {
		for i := range boards {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if boards[i][j][k] == num {
						boards[i][j][k] = -1
						boards[i][j][5]++
						boards[i][5][k]++

						if boards[i][j][5] == 5 || boards[i][5][k] == 5 {
							fmt.Println(calculate(boards[i]) * num)
							return
						}
					}
				}
			}
		}
	}
}

func p2() {
	fd, err := os.Open("E:\\test\\aoc\\day4\\data")
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	items := strings.Split(str, ",")
	nums := make([]int, len(items))
	for i, item := range items {
		num, _ := strconv.Atoi(item)
		nums[i] = num
	}

	boards := [][][]int{}
	index, row := -1, 5
	for {
		sb, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil || len(sb) != 14 {
			continue
		}

		if row == 5 {
			row = 0
			index++
			// 6x6 矩阵 最后一行和最后一列表示该行/列 marked元素数量
			board := make([][]int, 6)
			for i := 0; i < 6; i++ {
				board[i] = make([]int, 6)
			}
			boards = append(boards, board)
		}

		for i := 0; i < 15; i += 3 {
			num := int(sb[i+1] - '0')
			if sb[i] != ' ' {
				num += 10 * int(sb[i]-'0')
			}
			boards[index][row][i/3] = num
		}

		row++
	}

	fmt.Println("finish constructing boards")

	n := len(boards)
	finish := make([]bool, n)
	for _, num := range nums {
		for i := range boards {
			if finish[i] {
				continue
			}
			for j := 0; j < 5; j++ {
				if finish[i] {
					break
				}
				for k := 0; k < 5; k++ {
					if boards[i][j][k] == num {
						boards[i][j][k] = -1
						boards[i][j][5]++
						boards[i][5][k]++

						if boards[i][j][5] == 5 || boards[i][5][k] == 5 {
							finish[i] = true
							n--
							if n == 0 {
								fmt.Println(num)
								fmt.Println(boards[i])
								fmt.Println(num * calculate(boards[i]))
								return
							}
							break
						}
					}
				}
			}
		}
	}
}

func calculate(board [][]int) int {
	res := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				res += board[i][j]
			}
		}
	}
	return res
}
