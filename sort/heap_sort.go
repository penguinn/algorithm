package sort

// 堆排序
// 先建堆再排序，排序的时候同样也会堆话
// 时间复杂度O(nlogn)，空间复杂度O(1)，
func HeapSort(data []int) []int {
	length := len(data)

	// 建堆
	for i := length/2 - 1; i >= 0; i-- {
		heap(data, i, length-1)
	}

	// 排序
	for i := length - 1; i > 0; i-- {
		data[i], data[0] = data[0], data[i]
		heap(data, 0, i-1)
	}

	return data
}

// 堆化函数
func heap(data []int, start, end int) {
	left := start*2 + 1
	right := start*2 + 2
	if left > end {
		return
	}
	tmp := left
	if right <= end && data[right] > data[left] {
		tmp = right
	}

	if data[tmp] > data[start] {
		data[tmp], data[start] = data[start], data[tmp]
		heap(data, tmp, end)
	} else {
		return
	}
}
