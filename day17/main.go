package main

import (
	"fmt"
)

func p1() {
	lx, rx := 217, 240
	ly, ry := -69, -126
	fmt.Println(lx, rx, ly, ry)

	y := -ly - 1 // 初始y速度等于进入水面时的y速度，只是方向相反。所以入水后的下一秒速度为ry就行了
	fmt.Println((y + 1) * y / 2)
}

func p2() {
	lx, rx := 217, 240
	ly, ry := -69, -126

	res := 0
	for x := 1; x < 241; x++ {
		for y := -126; y < 126; y++ {
			if canReach(x, y, lx, rx, ly, ry) {
				res++
			}
		}
	}

	fmt.Println(res)
}

func canReach(x, y, lx, rx, ly, ry int) bool {
	posX, posY := 0, 0
	for {
		posX += x
		posY += y
		if x > 0 {
			x--
		}
		y--

		if lx <= posX && posX <= rx && posY <= ly && ry <= posY {
			return true
		}
		if posX > rx || posY < ry {
			return false
		}
	}

	return false
}

//  水平方向以最小速度出发
func main() {
	//p1()
	p2()
}

// x*(x+1) >= n 的最小x
func lowerBound(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)>>1
		if mid*(mid+1) >= n {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// x*(x+1) <= n 的最大x
func upperBound(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l+1)>>1
		if mid*(mid+1) <= n {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}
