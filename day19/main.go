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

type point struct {
	axis [3]int
}

type scanner struct {
	points []point
	id     int
	x      int
	y      int
	z      int
}

type axisInfo struct {
	axis int
	sign int
	diff int
}

func main() {
	fd, err := os.OpenFile("E:\\test\\aoc\\day19\\data", os.O_RDONLY, 0751)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()

	var scanners []*scanner
	var s *scanner

	reader := bufio.NewReader(fd)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if str == "\n" {
			continue
		}

		str = str[:len(str)-1]
		if strings.HasPrefix(str, "---") {
			s = &scanner{id: len(scanners)}
			scanners = append(scanners, s)
		} else {
			items := strings.Split(str, ",")
			x, _ := strconv.Atoi(items[0])
			y, _ := strconv.Atoi(items[1])
			z, _ := strconv.Atoi(items[2])
			s.points = append(s.points, point{[3]int{x, y, z}})
		}
	}

	pointSet := map[point]struct{}{}
	visited := make([]bool, len(scanners))
	visited[0] = true
	queue := []*scanner{scanners[0]}
	for _, p := range scanners[0].points {
		pointSet[p] = struct{}{}
	}

	for len(queue) > 0 {
		cur := queue[0]
		//fmt.Println(cur.id)
		queue = queue[1:]

		// 通过edges知道与哪些scanner是相邻的
		edges := getXEdges(cur.id, scanners, visited)
		// 求相邻的scanner的y、z轴对应关系
		yEdges, zEdges := inference(edges, cur.id, scanners)
		for k, v := range edges { // k: scanner ID
			//fmt.Println(cur.id, k)
			visited[k] = true
			queue = append(queue, scanners[k])
			scanners[k].x, scanners[k].y, scanners[k].z = v.diff, yEdges[k].diff, zEdges[k].diff

			fmt.Println(v.diff, v.sign, v.axis)
			fmt.Println(yEdges[k].diff, yEdges[k].sign, yEdges[k].axis)
			fmt.Println(zEdges[k].diff, zEdges[k].sign, zEdges[k].axis)

			for i := range scanners[k].points {
				//fmt.Println("before", scanners[k].points[i])
				scanners[k].points[i].axis[0], scanners[k].points[i].axis[1], scanners[k].points[i].axis[2] =
					v.diff+v.sign*scanners[k].points[i].axis[v.axis],
					yEdges[k].diff+yEdges[k].sign*scanners[k].points[i].axis[yEdges[k].axis],
					zEdges[k].diff+zEdges[k].sign*scanners[k].points[i].axis[zEdges[k].axis]
				pointSet[scanners[k].points[i]] = struct{}{}
			}
			//fmt.Println(len(pointSet))
		}

	}

	fmt.Println(len(pointSet))

	maxDist := 0
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			maxDist = max(maxDist, abs(scanners[i].x-scanners[j].x)+
				abs(scanners[i].y-scanners[j].y)+abs(scanners[i].z-scanners[j].z))
		}
	}

	fmt.Println(maxDist)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inference(edges map[int]axisInfo, id int, scanners []*scanner) (map[int]axisInfo, map[int]axisInfo) {
	y, z := map[int]axisInfo{}, map[int]axisInfo{}
	for otherId := range edges {
		for axis := 0; axis < 3; axis++ {
			for sign := -1; sign <= 1; sign += 2 {
				dy, dz := map[int]int{}, map[int]int{}
				for _, p := range scanners[id].points {
					for _, otherPoint := range scanners[otherId].points {
						dy[p.axis[1]-otherPoint.axis[axis]*sign]++
						dz[p.axis[2]-otherPoint.axis[axis]*sign]++
					}
				}

				yDiff, zDiff := getMaxValue(dy), getMaxValue(dz)
				if dy[yDiff] >= 12 {
					y[otherId] = axisInfo{
						axis: axis,
						sign: sign,
						diff: yDiff,
					}
				}
				if dz[zDiff] >= 12 {
					z[otherId] = axisInfo{
						axis: axis,
						sign: sign,
						diff: zDiff,
					}
				}
			}
		}
	}
	return y, z
}

func getXEdges(id int, scanners []*scanner, visited []bool) map[int]axisInfo {
	m := map[int]axisInfo{}

	// 先判断与其他scanner是否有至少12个重叠的beacon
	for i, s := range scanners {
		if visited[i] {
			continue
		}

		for axis := 0; axis < 3; axis++ {
			for sign := -1; sign <= 1; sign += 2 {
				dx := map[int]int{}
				for _, p := range scanners[id].points {
					for _, otherPoint := range s.points {
						dx[p.axis[0]-otherPoint.axis[axis]*sign]++
					}
				}

				maxK := getMaxValue(dx)
				// 如果两点重合，那么它们之间的diff就是固定的
				if dx[maxK] >= 12 {
					fmt.Println("maxK", maxK, dx[maxK])
					m[i] = axisInfo{
						axis: axis,
						sign: sign,
						diff: maxK,
					}
				}
			}
		}
	}

	return m
}

func getMaxValue(m map[int]int) int {
	maxK := 0
	for k, v := range m {
		if v > m[maxK] {
			maxK = k
		}
	}

	return maxK
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
