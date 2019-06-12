package array

func MedianOfTwo(arr1, arr2 []int) float64 {
	len1 := len(arr1)
	len2 := len(arr2)
	len := len1 + len2
	if (len & 1) == 1 {
		result := Help(arr1, arr2, len1, len2, 0, 0, len/2+1)
		return float64(result)
	} else {
		result := Help(arr1, arr2, len1, len2, 0, 0, len/2+1) +
			Help(arr1, arr2, len1, len2, 0, 0, len/2)
		return float64(result) / 2
	}

}

func Help(arr1, arr2 []int, len1, len2 int, index1, index2 int, k int) int {
	if len1 > len2 {
		return Help(arr2, arr1, len2, len1, index2, index1, k)
	} else {
		if k == 1 {
			if arr1[index1] > arr2[index2] {
				return arr2[index2]
			} else {
				return arr1[index1]
			}
		}
		if index1 == len1 {
			return arr2[index2+k]
		}
		p := k >> 1
		q := k - p
		if arr1[index1+p] < arr2[index2+q] {
			return Help(arr1, arr2, len1, len2, index1+p, index2, k-p)
		} else {
			return Help(arr1, arr2, len1, len2, index1, index2+q, k-q)
		}
	}
}
