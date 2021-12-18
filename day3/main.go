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
	fd, err := os.OpenFile("E:\\test\\aoc\\day3\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	var count []int // positive: 1 negative: 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		n := len(str) - 1
		if len(count) == 0 {
			count = make([]int, n)
		}

		for i := 0; i < n; i++ {
			if str[i] == '1' {
				count[i]++
			} else {
				count[i]--
			}
		}
	}

	gamma := toDecimal(count)
	epsilon := 1<<len(count) - 1 - gamma
	fmt.Println(gamma * epsilon)
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day3\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	var strs []string // positive: 1 negative: 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		strs = append(strs, str[:len(str)-1])
	}

	fmt.Println(getOxygen(strs) * getCO2(strs))
}

func toDecimal(count []int) int {
	n, res := len(count), 0
	for i := 0; i < n; i++ {
		res <<= 1
		if count[i] > 0 {
			res |= 1
		}
	}

	return res
}

// most common   1 priority
func getOxygen(strs []string) int {
	queue := strs
	n, res, index := len(strs[0]), 0, 0
	for len(queue) > 1 && index < n {
		count := 0
		for _, str := range queue {
			if str[index] == '1' {
				count++
			} else {
				count--
			}
		}

		newQueue := []string{}
		for _, str := range queue {
			if str[index] == '1' && count >= 0 {
				newQueue = append(newQueue, str)
			} else if str[index] == '0' && count < 0 {
				newQueue = append(newQueue, str)
			}
		}
		queue = newQueue
		index++
	}

	str := queue[0]
	for i := 0; i < n; i++ {
		res <<= 1
		if str[i] == '1' {
			res |= 1
		}
	}
	return res
}

// least common 0 priority
func getCO2(strs []string) int {
	queue := strs
	n, res, index := len(strs[0]), 0, 0
	for len(queue) > 1 && index < n {
		count := 0
		for _, str := range queue {
			if str[index] == '1' {
				count++
			} else {
				count--
			}
		}

		newQueue := []string{}
		for _, str := range queue {
			if str[index] == '0' && count >= 0 {
				newQueue = append(newQueue, str)
			} else if str[index] == '1' && count < 0 {
				newQueue = append(newQueue, str)
			}
		}
		queue = newQueue
		index++
	}

	str := queue[0]
	for i := 0; i < n; i++ {
		res <<= 1
		if str[i] == '1' {
			res |= 1
		}
	}
	return res
}
