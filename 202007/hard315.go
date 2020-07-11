package main

import "sort"

/*
给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例:

输入: [5,2,6,1]
输出: [2,1,1,0]
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.


*/
var a, c []int

func countSmaller(nums []int) []int {
	resultList := []int{}
	discretization(nums)
	c = make([]int, len(nums)+5)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		resultList = append(resultList, query(id-1))
		update(id)
	}
	for i := 0; i < len(resultList)/2; i++ {
		resultList[i], resultList[len(resultList)-1-i] = resultList[len(resultList)-1-i], resultList[i]
	}
	return resultList
}

func lowBit(x int) int {
	return x & (-x)
}

func update(pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowBit(pos)
	}
}

func query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a = make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)
}

func getId(x int) int {
	return sort.SearchInts(a, x) + 1
}
