package main

import (
	"fmt"
)

func main() {
	//var in = make([]int, 5)
	//var pre = make([]int, 5)
	var in = [...]int{1, 2, 3, 4, 5}
	var pre = [...]int{3, 2, 1, 4, 5}
	restore(in[:5], pre[:5])
}

func restore(in []int, pre []int) {
	var helpMap = make(map[int]int)
	for i := 0; i < len(in); i++ {
		helpMap[in[i]] = i
	}
	dfs(pre, helpMap, 0, len(pre)-1, 0)
}

func dfs(pre []int, helpMap map[int]int, left int, right int, index int) {
	if left > right {
		return
	}
	var pos = helpMap[pre[index]]
	leftSize := pos - left
	f := dfs
	f(pre, helpMap, left, pos-1, index+1)
	f(pre, helpMap, pos+1, right, index+leftSize+1)
	fmt.Println(pre[index])
}
