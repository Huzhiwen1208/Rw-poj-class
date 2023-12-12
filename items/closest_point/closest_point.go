package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	X, Y int
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func closestDistance(points []Point) float64 {
	n := len(points)
	if n < 2 {
		return math.Inf(1) // 不存在距离小于10000的点
	}

	// 按照 x 坐标排序
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	// 分治法求解最近距离
	var divideAndConquer func([]Point) float64
	divideAndConquer = func(pts []Point) float64 {
		m := len(pts)
		if m <= 3 {
			minDist := math.Inf(1)
			for i := 0; i < m-1; i++ {
				for j := i + 1; j < m; j++ {
					dist := distance(pts[i], pts[j])
					if dist < minDist {
						minDist = dist
					}
				}
			}
			return minDist
		}

		mid := m / 2
		midPoint := pts[mid]
		leftDist := divideAndConquer(pts[:mid])
		rightDist := divideAndConquer(pts[mid:])
		minDist := math.Min(leftDist, rightDist)

		// 找出中间区域内横坐标距离中点不超过最小距离的点
		midStrip := []Point{}
		for i := 0; i < m; i++ {
			if math.Abs(float64(pts[i].X-midPoint.X)) < minDist {
				midStrip = append(midStrip, pts[i])
			}
		}

		// 在中间区域内寻找更小的距离
		stripSize := len(midStrip)
		for i := 0; i < stripSize-1; i++ {
			for j := i + 1; j < stripSize && float64(midStrip[j].Y-midStrip[i].Y) < minDist; j++ {
				dist := distance(midStrip[i], midStrip[j])
				if dist < minDist {
					minDist = dist
				}
			}
		}

		return minDist
	}

	return divideAndConquer(points)
}

func main() {
	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nStr[:len(nStr)-1])

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		coords := []int{}
		for _, coordStr := range line[:len(line)-1] {
			coord, _ := strconv.Atoi(string(coordStr))
			coords = append(coords, coord)
		}
		points[i] = Point{X: coords[0], Y: coords[1]}
	}

	result := closestDistance(points)

	if result < 10000 {
		fmt.Printf("%.4f\n", result)
	} else {
		fmt.Println("INFINITY")
	}
}
