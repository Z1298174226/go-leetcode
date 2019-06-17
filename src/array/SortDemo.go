package array

type pointInt *[]int

/**
快排
*/
func QuickSort(left int, right int, arr pointInt) {
	if left >= right {
		return
	} else {
		//j := partition(left, right, arr)
		j := PartitionUpdate(left, right, arr)
		QuickSort(left, j-1, arr)
		QuickSort(j+1, right, arr)
	}
}

func Partition(left int, right int, arr pointInt) int {
	i := left
	j := right + 1
	for {
		for {
			i = i + 1
			if LessInt(left, i, arr) {
				break
			} else {
				if i == right {
					break
				}
			}
		}
		for {
			j = j - 1
			if LessInt(j, left, arr) {
				break
			} else {
				if j == left {
					break
				}
			}
		}
		if i >= j {
			break
		}
		SwapInt(i, j, arr)
	}
	SwapInt(left, j, arr)
	return j
}

func PartitionUpdate(left int, right int, arr pointInt) int {
	i := left - 1
	for j := left; j < right; j++ {
		if LessInt(j, right, arr) {
			i++
			SwapInt(i, j, arr)
		}
	}
	i++
	SwapInt(i, right, arr)
	return i
}

func SwapInt(i int, j int, arr pointInt) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func LessInt(i int, j int, arr pointInt) bool {
	return (*arr)[i] < (*arr)[j]
}

/**
堆排序
*/

func SwapHeap(i int, j int, arr pointInt) {
	(*arr)[i-1], (*arr)[j-1] = (*arr)[j-1], (*arr)[i-1]
}

func LessHeap(i int, j int, arr pointInt) bool {
	return (*arr)[i-1] > (*arr)[j-1]
}

func sinkOperation(left int, right int, arr pointInt) {
	harf := right >> 1
	for {
		if left > harf {
			return
		}
		child := left << 1
		if child < right && LessHeap(child, child+1, arr) {
			child++
		}
		if LessHeap(left, child, arr) {
			SwapHeap(left, child, arr)
		}
		left = child
	}
}

func HeapSort(arr pointInt) {
	len := len(*arr)
	for i := len / 2; i >= 1; i-- {
		sinkOperation(i, len, arr)
	}
	for {
		if len <= 1 {
			return
		} else {
			SwapHeap(1, len, arr)
			len--
			sinkOperation(1, len, arr)
		}
	}
}

/**
shell sort
*/

func shellSort(arr pointInt) {
	len := len(*arr)
	h := 1
	for {
		if h >= len/3 {
			break
		} else {
			h = h*3 + 1
		}
	}
	for {
		if h < 1 {
			break
		}
		for i := h; i < len; i++ {
			for j := i; j >= h && lessShell(j, j-h, arr); j = j - h {
				swapShell(j, j-h, arr)
			}
		}
		h = h / 3
	}
}

func swapShell(i int, j int, arr pointInt) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func lessShell(i int, j int, arr pointInt) bool {
	return (*arr)[i] < (*arr)[j]
}
