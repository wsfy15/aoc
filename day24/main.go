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

type instruction struct {
	isInput bool
	op      func(int, int) (int, bool)
	l       string
	r       string
}

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day24\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	program := []instruction{}
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " ")
		switch items[0] {
		case "inp":
			program = append(program, instruction{
				isInput: true,
				l:       items[1],
			})
		case "add":
			program = append(program, instruction{
				isInput: false,
				op: func(l int, r int) (int, bool) {
					return l + r, true
				},
				l: items[1],
				r: items[2],
			})
		case "mul":
			program = append(program, instruction{
				isInput: false,
				op: func(l int, r int) (int, bool) {
					return l * r, true
				},
				l: items[1],
				r: items[2],
			})
		case "div":
			program = append(program, instruction{
				isInput: false,
				op: func(l int, r int) (int, bool) {
					if r == 0 {
						return 0, false
					}
					return l / r, true
				},
				l: items[1],
				r: items[2],
			})
		case "mod":
			program = append(program, instruction{
				isInput: false,
				op: func(l int, r int) (int, bool) {
					if l < 0 || r <= 0 {
						return 0, false
					}
					return l % r, true
				},
				l: items[1],
				r: items[2],
			})
		case "eql":
			program = append(program, instruction{
				isInput: false,
				op: func(l int, r int) (int, bool) {
					if l == r {
						return 1, true
					}
					return 0, true
				},
				l: items[1],
				r: items[2],
			})

		}
	}

	fmt.Println(validate(92928914999991, program))
	fmt.Println(validate(91811211611981, program))
	return
	var num int = 1e14 - 1
	for num > 1e13 {
		fmt.Println(num)
		if validate(num, program) {
			fmt.Println(num)
			return
		}

		num--
	}

}

func validate(num int, program []instruction) bool {
	str := make([]int, 14)
	for i := 13; i >= 0; i-- {
		str[i] = num % 10
		if str[i] == 0 {
			return false
		}

		num /= 10
	}

	index := 0
	variable := map[string]int{
		"x": 0,
		"y": 0,
		"z": 0,
		"w": 0,
	}
	var success bool
	for _, ins := range program {
		if ins.isInput {
			variable[ins.l] = str[index]
			index++
			continue
		}

		if v, ok := variable[ins.r]; ok {
			variable[ins.l], success = ins.op(variable[ins.l], v)
		} else {
			num, _ := strconv.Atoi(ins.r)
			variable[ins.l], success = ins.op(variable[ins.l], num)
		}

		if !success {
			return false
		}
	}

	return variable["z"] == 0
}
