package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	//puzzle1()
	puzzle2()
}

func puzzle2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day1\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	nums := []int{}
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		num, err := strconv.Atoi(str[:len(str)-1])
		if err != nil {
			log.Println(err)
			return
		}
		nums = append(nums, num)
	}

	res := 0
	sum := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		if nums[i] > nums[i-3] {
			res++
		}
		sum += nums[i] - nums[i-3]
	}
	fmt.Println(res)
}

func puzzle1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day1\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	count, prev := 0, -1
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		num, err := strconv.Atoi(str[:len(str)-1])
		if err != nil {
			log.Println(err)
			return
		}

		if prev != -1 && num > prev {
			count++
		}
		prev = num
	}

	fmt.Println(count)
}
