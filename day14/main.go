package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day14\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	g := make([][]int, 26)
	for i := 0; i < 26; i++ {
		g[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			g[i][j] = -1
		}
	}

	state := "CFFPOHBCVVNPHCNBKVNV"
	//state := "NNCB"
	exist := [26]bool{}
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " -> ")
		l1, l2, r := items[0][0]-'A', items[0][1]-'A', items[1][0]-'A'
		g[l1][l2] = int(r)
		exist[l1], exist[l2], exist[r] = true, true, true
	}

	count := make([]int, 26)
	process(40, state, count, g)
	fmt.Println(count)
	fmt.Println(max(exist, count...))
	fmt.Println(min(exist, count...))
	fmt.Println(max(exist, count...) - min(exist, count...))
}

func process(round int, state string, count []int, g [][]int) {
	pairs := map[string]int{}
	for i := 0; i < len(state)-1; i++ {
		pairs[state[i:i+2]]++
	}

	for round > 0 {
		newPairs := map[string]int{}
		for key, value := range pairs {
			cur, next := key[0]-'A', key[1]-'A'
			if g[cur][next] != -1 {
				str := string([]byte{key[0], byte(g[cur][next] + 'A')})
				newPairs[str] += value
				str = string([]byte{byte(g[cur][next] + 'A'), key[1]})
				newPairs[str] += value
			} else {
				newPairs[key] = value
			}
		}

		pairs = newPairs
		round--
	}

	for key, value := range pairs {
		cur, next := key[0]-'A', key[1]-'A'
		count[cur] += value
		count[next] += value
	}

	count[state[0]-'A']++
	count[state[len(state)-1]-'A']++
	for i := 0; i < len(count); i++ { // every char exist in two pair
		count[i] >>= 1
	}
	return
}

func max(exist [26]bool, a ...int) int {
	res := a[0]
	for i, num := range a {
		if !exist[i] {
			continue
		}
		if num > res {
			res = num
		}
	}
	return res
}

func min(exist [26]bool, a ...int) int {
	res := math.MaxInt64
	for _, num := range a {
		if num == 0 {
			continue
		}
		if num < res {
			res = num
		}
	}
	return res
}
