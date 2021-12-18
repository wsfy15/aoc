package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//n, _ := construct("[1,3]", 0)
	//fmt.Println(n)
	//p1()
	p2()
}

type node struct {
	left   *node
	right  *node
	parent *node
	val    int
	isPair bool
}

func construct(str string, index int) (*node, int) {
	if index >= len(str) {
		return nil, index
	}

	cur := &node{}
	if str[index] == '[' {
		cur.isPair = true
		cur.left, index = construct(str, index+1) // index + 1 跳过 [
		// 返回的index已经跳过 ] 或 ,
		cur.right, index = construct(str, index)
		cur.left.parent = cur
		cur.right.parent = cur
	} else {
		cur.val, index = parseNum(str, index)
		// 返回的index指向,或]
	}

	return cur, index + 1 // index + 1 跳过 ] 或 ,
}

func parseNum(str string, index int) (int, int) {
	res := 0
	for index < len(str) && '0' <= str[index] && str[index] <= '9' {
		res = res*10 + int(str[index]-'0')
		index++
	}

	return res, index
}

func add(prev, cur *node) *node {
	newHead := &node{
		left:   prev,
		right:  cur,
		val:    0,
		isPair: true,
	}

	prev.parent = newHead
	cur.parent = newHead

	for {
		needContinue := true
		for needContinue {
			needContinue = explode(newHead, 0)
		}

		// 不能进行explode操作，才会考虑split

		// split一次就重新整个循环
		needContinue = split(newHead)
		if !needContinue {
			break
		}
	}

	return newHead
}

func explode(n *node, depth int) bool {
	if depth == 4 {
		if n.isPair {
			prev, next := findPrev(n), findNext(n)
			if prev != nil {
				prev.val += n.left.val
			}
			if next != nil {
				next.val += n.right.val
			}

			n.left, n.right = nil, nil
			n.isPair, n.val = false, 0
			return true
		}

		return false
	}

	if n.isPair {
		l, r := explode(n.left, depth+1), explode(n.right, depth+1)
		return l || r
	}
	return false
}

func findNext(n *node) *node {
	parent, cur := n.parent, n
	for parent != nil && parent.right == cur {
		parent, cur = parent.parent, parent
	}

	if parent == nil {
		return nil
	}

	cur = parent.right
	for cur.isPair && cur.left != nil {
		cur = cur.left
	}

	return cur
}

func findPrev(n *node) *node {
	parent, cur := n.parent, n
	for parent != nil && parent.left == cur {
		parent, cur = parent.parent, parent
	}

	if parent == nil {
		return nil
	}

	cur = parent.left
	for cur.isPair && cur.right != nil {
		cur = cur.right
	}

	return cur
}

func split(n *node) bool {
	if n.isPair {
		// 一旦完成了一次切割，就得重新扫描
		//l, r := split(n.left), split(n.right)
		//return l || r
		return split(n.left) || split(n.right)
	}

	if n.val > 9 {
		n.isPair = true
		n.left = &node{
			parent: n,
			val:    n.val / 2,
			isPair: false,
		}
		n.right = &node{
			parent: n,
			val:    (n.val + 1) / 2,
			isPair: false,
		}
		n.val = 0
		return true
	}

	return false
}

func getMagnitude(n *node) int {
	if !n.isPair {
		return n.val
	}

	return 3*getMagnitude(n.left) + 2*getMagnitude(n.right)
}

func p1() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day18\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	var prev *node
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		cur, _ := construct(str, 0)
		if prev == nil {
			prev = cur
		} else {
			prev = add(prev, cur)
		}
		printTree(prev)
		fmt.Println()
	}

	printTree(prev)
	fmt.Println(getMagnitude(prev))
}

func p2() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day18\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	var prev *node
	strs := []string{}
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF || str == "\n" {
			break
		}

		str = str[:len(str)-1]
		strs = append(strs, str)
		cur, _ := construct(str, 0)
		if prev == nil {
			prev = cur
		} else {
			prev = add(prev, cur)
		}
		printTree(prev)
		fmt.Println()
	}

	res := 0
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			cur, _ := construct(strs[i], 0)
			another, _ := construct(strs[j], 0)
			res = max(res, getMagnitude(add(cur, another)))

			cur, _ = construct(strs[i], 0)
			another, _ = construct(strs[j], 0)
			res = max(res, getMagnitude(add(another, cur)))
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTree(cur *node) {
	if cur.isPair {
		fmt.Print("[")
		printTree(cur.left)
		fmt.Print(",")
		printTree(cur.right)
		fmt.Print("]")
	} else {
		fmt.Print(cur.val)
	}
}

