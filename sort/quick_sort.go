package sort

// 快速排序
// 先找到分界点，小于分界点的放左边，大于分界点的放右边（找分界点的方法也很简单，弄个双指针，快指针用于遍历，慢指针用于存放左边小于临界点的数）
// 空间复杂度O(1)，时间复杂度O(nlogn), 非稳定排序（todo 感觉可以通过修改partition来成为有序排序，不要使用交换的方式）
func QuickSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	quickSort(data, 0, length-1)
	return data
}

func quickSort(data []int, start, end int) {
	if start >= end {
		return
	} else {
		index := partition(data, start, end)
		quickSort(data, start, index-1)
		quickSort(data, index+1, end)
	}
}

func partition(data []int, start, end int) int {
	guard := data[end]
	i := start
	for j := start; j <= end-1; j++ {
		if data[j] < guard {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[end] = data[end], data[i]
	return i
}
