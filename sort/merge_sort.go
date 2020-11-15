package sort

// 归并排序
// 把数据二分拆开，然后两两归并
// 空间复杂度O(n)，这里空间可以再优化一下，时间复杂度O(nlogn),稳定排序
func MergeSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	index := length / 2
	left := MergeSort(data[0:index])
	right := MergeSort(data[index:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	lengthLeft := len(left)
	lengthRight := len(right)
	result := make([]int, lengthLeft+lengthRight)
	i, j := 0, 0
	for i < lengthLeft && j < lengthRight {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}
	for i < lengthLeft {
		result[i+lengthRight] = left[i]
		i++
	}
	for j < lengthRight {
		result[j+lengthLeft] = right[j]
		j++
	}

	return result
}
