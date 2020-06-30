package sort

import (
	"sort"
)

// 插入排序
// 数组分为已排序段和未排序段，每次选择未排序段第一个数，倒着和已排序段比较，找好位置插入（由于这里没用额外的空间，感觉中间赋值的语句可能会更多）
// 时间复杂度O(n2)，空间复杂度O(1)，稳定排序
func InsertionSort(data sort.Interface) {
	length := data.Len()
	for i := 1; i <= length-1; i++ {
		for j := i - 1; j >= 0; j-- {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			} else {
				break
			}
		}
	}
}
