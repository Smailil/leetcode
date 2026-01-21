package main

func isReflected(points [][]int) bool {
	minX := points[0][0]
	maxX := points[0][0]
	mapPoints := make(map[[2]int]bool)
	for _, p := range points {
		minX = min(p[0], minX)
		maxX = max(p[0], maxX)
		mapPoints[[2]int{p[0], p[1]}] = true
	}
	s := minX + maxX
	// med = (x'+x)/2, med=s/2 => x' = 2*med - x, s=2*med
	for point := range mapPoints {
		if !mapPoints[[2]int{s - point[0], point[1]}] {
			return false
		}
	}
	return true
}
