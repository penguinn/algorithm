package sort

import (
	"sort"
)

// 冒泡排序
// 每次从第一个元素开始，一直冒泡到他不能冒泡的数
// 时间复杂度O(n2)，空间复杂度O(n)，稳定排序
func BubbleSort(data sort.Interface) {
	length := data.Len()
	var isChanged bool
	for i := 0; i < length-1; i++ {
		isChanged = false
		for j := 0; j < length-i-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
