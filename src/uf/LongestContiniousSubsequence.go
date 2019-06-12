package uf

type point *[]int
type set map[int]bool

func ContinousSebsequence(arr point) int {
	parents := make(map[int]int)
	results := make(map[int]set)
	if len(*arr) == 0 {
		return 0
	}
	count := 1
	for i := range *arr {
		parents[(*arr)[i]] = (*arr)[i]
	}
	for i := range *arr {
		if _, ok := parents[(*arr)[i]-1]; ok {
			Union(&parents, (*arr)[i], (*arr)[i]-1)
		}
	}
	for i := range *arr {
		parent := Find(&parents, (*arr)[i])
		if _, ok := results[parent]; !ok {
			results[parent] = make(set)
		}
		if _, ok := results[parent][(*arr)[i]]; !ok {
			results[parent][(*arr)[i]] = true
		}
	}
	for _, result := range results {
		if count < len(result) {
			count = len(result)
		}
	}
	return count
}

func Find(parents *map[int]int, element int) int {
	if (*parents)[element] == element {
		return element
	} else {
		return Find(parents, (*parents)[element])
	}
}

func IsConnected(parents *map[int]int, ele1, ele2 int) bool {
	return Find(parents, ele1) == Find(parents, ele2)
}

func Union(parents *map[int]int, ele1, ele2 int) {
	pele1 := Find(parents, ele1)
	pele2 := Find(parents, ele2)
	(*parents)[pele2] = pele1
}
