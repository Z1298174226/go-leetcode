package binarytree

import (
	"fmt"
)

func Restore(in []int, pre []int) {
	var helpMap = make(map[int]int)
	for i := 0; i < len(in); i++ {
		helpMap[in[i]] = i
	}
	Dfs(pre, helpMap, 0, len(pre)-1, 0)
}

func Dfs(pre []int, helpMap map[int]int, left int, right int, index int) {
	if left > right {
		return
	}
	var pos = helpMap[pre[index]]
	leftSize := pos - left
	f := Dfs
	f(pre, helpMap, left, pos-1, index+1)
	f(pre, helpMap, pos+1, right, index+leftSize+1)
	fmt.Println(pre[index])
}
