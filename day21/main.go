package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day21\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
	}
}
