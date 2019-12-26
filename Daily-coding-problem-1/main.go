package main

import "fmt"

func FindSum(arr []int, num int, nums map[int]int, i int) bool {
	key := arr[i]

	_, pres := nums[num-key]
	if pres == true {
		return true
	}

	nums[key] = key
	if i < len(arr)-1 {
		return FindSum(arr, num, nums, i+1)
	}

	return false
}

func main() {
	nums := make(map[int]int)

	fmt.Println(FindSum([]int{10, 15, 3, 7}, 17, nums, 0))
}
