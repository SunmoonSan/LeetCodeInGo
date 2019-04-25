/*
@desc : Created by San on 2019/4/24 23:40
*/
package main

func containsDuplicate(nums []int) bool {
	kvs := make(map[int]bool)
	for _, n := range nums {
		if _, ok := kvs[n]; ok {
			return true
		}

		kvs[n] = false
	}

	return false
}
