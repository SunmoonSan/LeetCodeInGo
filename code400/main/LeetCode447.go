/*
@desc : Created by San on 2019/4/23 23:29
*/
package main

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	count := 0
	for i := 0; i < len(points); i++ {
		kvs := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i != j {
				dis := getDistance(points[i], points[j])
				kvs[dis]++
			}
		}

		for k := range kvs {
			t, _ := kvs[k]
			count += t * (t - 1)
		}
	}

	return count
}

func getDistance(point1 []int, point2 []int) int {
	return (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1])
}

func main() {
	points := [][]int{{0, 0}, {1, 0}, {2, 0}}
	count := numberOfBoomerangs(points)
	fmt.Println(count)
}
