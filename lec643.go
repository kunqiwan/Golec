package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1,12,-5,-6,50,3}
	k:=4
	fmt.Println(findMaxAverage(nums,k))
}

func findMaxAverage(nums []int, k int) float64 {
left:=0
right:=0
countsum :=0
var max float64 = 0.0
for right<len(nums){
	countsum = countsum+nums[right]
	right++
	if right - left == k{
			var avg float64 = (float64(countsum)/float64((k)))
			max = math.Max(avg,max)
		    countsum = countsum -nums[left]
			left++

}
}
return max
}
