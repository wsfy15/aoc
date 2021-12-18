package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

var m = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var m2 = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day10\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	res := 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		score, _ := corruptedScore(str)
		res += score
	}

	fmt.Println(res)
}

func corruptedScore(str string) (int, []byte) {
	stack := []byte{}
	for _, ch := range []byte(str) {
		if ch == '(' || ch == '{' || ch == '[' || ch == '<' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return m[ch], nil
			}

			top := stack[len(stack)-1]
			switch ch {
			case ')':
				if top != '(' {
					return m[')'], nil
				}
			case '}':
				if top != '{' {
					return m['}'], nil
				}
			case ']':
				if top != '[' {
					return m[']'], nil
				}
			case '>':
				if top != '<' {
					return m['>'], nil
				}
			}
			stack = stack[:len(stack)-1]
		}
	}

	return 0, stack
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day10\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	nums := []int{}
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		score, stack := corruptedScore(str)
		if score == 0 && len(stack) > 0 {
			nums = append(nums, completionScore(stack))
		}
	}

	sort.Ints(nums)
	fmt.Println(nums[len(nums)/2])
}

func completionScore(stack []byte) int {
	score := 0
	for len(stack) > 0 {
		score *= 5
		score += m2[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
	}

	return score
}
