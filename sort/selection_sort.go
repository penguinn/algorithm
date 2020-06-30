package sort

import (
	"sort"
)

// 选择排序
// 每次选出最小的数放到表头
// 时间复杂度O(n2)，空间复杂度O(1)，非稳定排序
func SelectionSort(data sort.Interface) {
	length := data.Len()
	var min int
	for i := 0; i < length-1; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(min, i)
		}
	}
}
