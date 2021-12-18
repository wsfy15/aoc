package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var m map[int]int

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day6\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	bytes, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}

	nums := []int{}
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ',' {
			continue
		}
		nums = append(nums, int(bytes[i]-'0'))
	}

	res := 0
	m = make(map[int]int)
	for _, num := range nums {
		res += dfs(256 + 6 - num)
	}
	res += len(nums)
	fmt.Println(res)
}

func dfs(days int) int {
	if days <= 0 {
		return 0
	}

	if v, ok := m[days]; ok {
		return v
	}
	
	tmp := days
	res := days / 7
	for days >= 7 {
		days -= 7
		res += dfs(days - 2)
	}

	m[tmp] = res
	return res
}
