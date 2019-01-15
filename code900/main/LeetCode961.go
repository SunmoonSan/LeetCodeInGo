package main

import (
	"fmt"
)

func repeatedNTimes(A []int) int {
	N := len(A) / 2

	kvs := map[int]int{}
	for _, a := range A {
		_, ok := kvs[a]
		if ok {
			kvs[a] = kvs[a] + 1
			if kvs[a] == N {
				return a
			}
		} else {
			kvs[a] = 1
		}
	}

	return -1
}

func main() {
	A := []int{5, 1, 5, 2, 5, 3, 5, 4}
	n := repeatedNTimes(A)
	fmt.Println(n)
}
