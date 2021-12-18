package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day7\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	bytes, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}

	items := strings.Split(string(bytes), ",")
	nums := []int{}
	n, sum := len(items), 0
	for _, item := range items {
		num, _ := strconv.Atoi(item)
		nums = append(nums, num)
		sum += num
	}

	sort.Ints(nums)
	l, r := nums[0], nums[n-1]
	for l < r {
		mid := l + (r-l)>>1
		//d, d2 := delta(mid, nums), delta(mid+1, nums)
		d, d2 := delta2(mid, nums), delta2(mid+1, nums)
		if d < d2 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
	fmt.Println(delta2(l, nums))
}

func delta(avg int, nums []int) int {
	res := 0
	for _, num := range nums {
		res += abs(num - avg)
	}
	return res
}

func delta2(avg int, nums []int) int {
	res := 0
	for _, num := range nums {
		res += sumOfSequence(abs(num - avg))
	}
	return res
}

func sumOfSequence(num int) int {
	return num * (num + 1) / 2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}
