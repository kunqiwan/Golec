package main

import "fmt"

func main() {
	nums := []int{1, 2, 3,1}
	k := 4
	fmt.Println(containsNearbyDuplicate(nums,k))
}
func containsNearbyDuplicate(nums []int, k int) bool {
	//left :=0
	right :=0
	records:= make(map[int]int, len(nums))
	for right<len(nums){
		curr :=nums[right]
		_,ok := records[curr]
		if ok{
			//存在这个值，和最近值对比
			if     right -records[curr] <=k{
				return true
			}
		}
		records[curr] = right
		right++
	}
	return false
}
