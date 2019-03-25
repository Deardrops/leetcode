package main

import "fmt"

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	for _, c := range s {
		alphabet[(c-'a')]++
	}
	for _, c := range t {
		alphabet[(c-'a')]--
	}
	for _, c := range alphabet {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "rat"
	t := "tar"
	res := isAnagram(s, t)
	fmt.Println(res)
}
