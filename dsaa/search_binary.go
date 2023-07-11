package main

import "fmt"

/*
二分查找

时间复杂度：O(logn)
局限性：
1、二分查找依赖的是顺序表结构，简单地说就是数组
2、针对的是有序的数据
3、数据量太小也不适合，因为这和顺序遍历的查找速度相差不大
4、数据量太大也不适合，因为内存可能无法分配出过大且连续的空间
*/

func bSearch1(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		// 当 low 和 high 数值较大的话，两者之和可能会出现溢出，需要优化一下
		//mid := (low + high) / 2
		mid := low + (high-low)/2
		fmt.Printf("low: %v, high: %v, mid: %v\n", low, high, mid)

		if arr[mid] == val {
			return mid
		}

		if arr[mid] > val {
			high = mid - 1
			continue
		}

		low = mid + 1
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 9, 10, 12, 16, 18, 20}

	search := 13
	pos := bSearch1(arr, search)
	fmt.Printf("bSearch1(%v): %v\n", search, pos)
}
