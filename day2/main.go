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
	fd, err := os.OpenFile("E:\\test\\aoc\\day2\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	horizon, vertical := 0, 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " ")
		num, _ := strconv.Atoi(items[1])
		switch items[0] {
		case "forward":
			horizon += num
		case "down":
			vertical += num
		case "up":
			vertical -= num
		}
	}

	fmt.Println(horizon * vertical)
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day2\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	aim, horizon, vertical := 0, 0, 0
	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		items := strings.Split(str, " ")
		num, _ := strconv.Atoi(items[1])
		switch items[0] {
		case "forward":
			horizon += num
			vertical += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	fmt.Println(horizon * vertical)
}
