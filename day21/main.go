package main

import "fmt"

func main() {
	//p1()
	p2()
}

func p1() {
	x, y := 1, 10
	//x, y := 4, 8
	xscore, yscore, count, index := 0, 0, 0, 0
	nums := [100]int{}
	for i := 0; i < 98; i++ {
		nums[i] = 3 * (i + 2)
	}
	nums[98], nums[99] = 200, 103

	for xscore < 1000 && yscore < 1000 {
		x = (x+nums[index]-1)%10 + 1
		index = (index + 3) % 100
		count++
		xscore += x
		if xscore >= 1000 {
			break
		}

		y = (y+nums[index]-1)%10 + 1
		index = (index + 3) % 100
		count++
		yscore += y
	}

	count *= 3
	fmt.Println(xscore, yscore, count)
	if xscore < 1000 {
		fmt.Println(xscore * count)
	} else {
		fmt.Println(yscore * count)
	}
}

func p2() {
	x, y := 1-1, 10-1
	//x, y := 4-1, 8-1
	dp := make([][][][][]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([][][][]int, 10)
		for j := 0; j < 10; j++ {
			dp[i][j] = make([][][]int, 21)
			for k := 0; k < 21; k++ {
				dp[i][j][k] = make([][]int, 21)
			}
		}
	}

	xwin, ywin := solution(x, y, 0, 0, dp)
	fmt.Println(xwin, ywin)
}

// 传入x、y两人当前的位置和得分
// 访问它们从这个位置和得分 在所有可能的宇宙中最终能赢的次数
func solution(x, y, xscore, yscore int, dp [][][][][]int) (int, int) {
	if xscore >= 21 {
		return 1, 0
	}
	if yscore >= 21 {
		return 0, 1
	}

	if len(dp[x][y][xscore][yscore]) == 2 {
		return dp[x][y][xscore][yscore][0], dp[x][y][xscore][yscore][1]
	}

	dp[x][y][xscore][yscore] = make([]int, 2)
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				newX := (x + d1 + d2 + d3) % 10
				newXscore := xscore + newX + 1

				a, b := solution(y, newX, yscore, newXscore, dp)
				dp[x][y][xscore][yscore][0] += b
				dp[x][y][xscore][yscore][1] += a
			}
		}
	}

	return dp[x][y][xscore][yscore][0], dp[x][y][xscore][yscore][1]
}
