package search

// 查找一个有序的且没有重复的数组
func BinarySearch(data []int, value int) int {
	length := len(data)
	if length == 0 {
		return -1
	}

	start := 0
	end := length - 1

	for start <= end {
		mid := start + ((end - start) >> 1)
		if data[mid] == value {
			return mid
		} else if data[mid] < value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchGE(data []int, value int) int {
	length := len(data)
	if length == 0 {
		return -1
	}

	start := 0
	end := length - 1

	for start <= end {
		mid := start + ((end - start) >> 1)
		if data[mid] < value {
			start = mid + 1
		} else if mid == 0 || data[mid-1] < value {
			return mid
		} else {
			end = mid - 1
		}
	}
	return -1
}
