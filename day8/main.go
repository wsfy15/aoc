package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day8\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	res := 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " | ")
		outputs := strings.Split(items[1], " ")
		for _, output := range outputs {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				res++
			}
		}
	}

	fmt.Println(res)
}

/*
在二极管中，0~9的表示中，每个字母出现的次数
a:8 b:6 c:8 d:7 e:4 f:9 g:7
1、4、7、8分别有2、4、3、7个竖线
通过1、7 可以 确定a，进而确定c   d存在于4中，也可以确定，g也就可以确定
*/

// 从高位到低位表示a-g
var m = map[int]int{
	0b1110111: 0,
	0b0010010: 1,
	0b1011101: 2,
	0b1011011: 3,
	0b0111010: 4,
	0b1101011: 5,
	0b1101111: 6,
	0b1010010: 7,
	0b1111111: 8,
	0b1111011: 9,
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day8\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	res := 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " | ")
		outputs := strings.Split(items[1], " ")
		patterns := strings.Split(items[0], " ")

		charMap := solution(patterns)
		res += getNum(outputs, charMap)
	}

	fmt.Println(res)
}

func solution(patterns []string) map[byte]byte {
	charMap := map[byte]byte{}
	count := map[byte]int{}
	var one, four, seven string
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			one = pattern
		case 3:
			seven = pattern
		case 4:
			four = pattern
			//case 7:
			//	eight = pattern
		}

		for _, ch := range []byte(pattern) {
			count[ch]++
		}
	}

	for _, ch := range []byte(seven) {
		exist := false
		for _, ach := range []byte(one) {
			if ch == ach {
				exist = true
			}
		}

		if !exist {
			charMap[ch] = 'a'
			break
		}
	}

	for key, value := range count {
		switch value {
		case 4:
			charMap[key] = 'e'
		case 6:
			charMap[key] = 'b'
		case 8:
			if _, ok := charMap[key]; !ok {
				charMap[key] = 'c'
			}
		case 9:
			charMap[key] = 'f'
		}
	}

	for _, ch := range []byte(four) {
		if _, ok := charMap[ch]; !ok {
			charMap[ch] = 'd'
		}
	}

	for key, _ := range count {
		if _, ok := charMap[key]; !ok {
			charMap[key] = 'g'
		}
	}

	return charMap
}

func getNum(outputs []string, charMap map[byte]byte) int {
	num := 0
	for _, output := range outputs {
		cur := 0
		for _, ch := range []byte(output) {
			cur |= 1 << (6 - (charMap[ch] - 'a'))
		}
		num = num*10 + m[cur]
	}
	return num
}
