package main

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	kvs := make(map[rune]int)
	for _, c := range s {
		kvs[c]++
	}

	for _, c := range t {
		if count, _ := kvs[c]; count > 0 {
			kvs[c]--
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "rat"
	t := "car"
	b := isAnagram(s, t)
	fmt.Println(b)

}
